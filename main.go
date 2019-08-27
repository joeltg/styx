package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	uuid "github.com/google/uuid"
	cid "github.com/ipfs/go-cid"
	ipfs "github.com/ipfs/go-ipfs-api"
	ld "github.com/piprate/json-gold/ld"

	styx "github.com/underlay/styx/db"
	loader "github.com/underlay/styx/loader"
)

// Replace at your leisure
const defaultPath = "/tmp/styx"
const defaultPort = "8000"
const defaultHost = "localhost:5001"

var path = os.Getenv("STYX_PATH")
var port = os.Getenv("STYX_PORT")
var host = os.Getenv("IPFS_HOST")

var shError = "IPFS Daemon not running"

func walkValues(values []interface{}, files map[string]string) {
	for _, value := range values {
		if object, is := value.(map[string]interface{}); is {
			for key, val := range object {
				if id, is := val.(string); is && key == "@id" {
					if uri, has := files[id]; has {
						object["@id"] = uri
					}
				} else if array, is := val.([]interface{}); is && (key == "@list" || key == "@set") {
					walkValues(array, files)
				}
			}
		}
	}
}

func walk(graph []interface{}, files map[string]string) {
	for _, element := range graph {
		if node, is := element.(map[string]interface{}); is {
			for key, val := range node {
				if id, is := val.(string); is && key == "@id" {
					if uri, has := files[id]; has {
						node["@id"] = uri
					}
				} else if values, is := val.([]interface{}); is && key == "@graph" {
					walk(values, files)
				} else if is {
					walkValues(values, files)
				}
			}
		}
	}
}

func main() {
	if path == "" {
		path = defaultPath
	}

	if port == "" {
		port = defaultPort
	}

	if host == "" {
		host = defaultHost
	}

	// Replace at your leisure
	var sh = ipfs.NewShell(host)

	if !sh.IsUp() {
		log.Fatal(shError)
	}

	peerID, err := sh.ID()
	if err != nil {
		log.Fatal(err)
	}

	db, err := styx.OpenDB(path)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	proc := ld.NewJsonLdProcessor()
	dl := loader.NewShellDocumentLoader(sh)
	options := styx.GetStringOptions(dl)

	store := styx.MakeShellDocumentStore(sh)

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	dir := http.Dir(wd + "/www")
	fs := http.FileServer(dir)
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" && req.URL.Path == "/" {
			ct := req.Header.Get("Content-Type")
			m, params, err := mime.ParseMediaType(ct)
			if err != nil {
				res.WriteHeader(400)
				return
			}

			var cid cid.Cid
			var reader io.Reader
			if m == "application/ld+json" {
				decoder := json.NewDecoder(req.Body)
				var doc interface{}
				if err := decoder.Decode(&doc); err != nil {
					res.WriteHeader(400)
					return
				}

				// Convert to RDF
				rdf, err := proc.Normalize(doc, options)
				if s, is := rdf.(string); !is || err != nil {
					res.WriteHeader(400)
					return
				} else if c, err := store(strings.NewReader(s)); err != nil {
					res.WriteHeader(500)
					return
				} else {
					cid = c
					reader = strings.NewReader(s)
				}
			} else if m == "application/n-quads" {
				if b, err := ioutil.ReadAll(req.Body); err != nil {
					res.WriteHeader(500)
					return
				} else if c, err := store(bytes.NewReader(b)); err != nil {
					res.WriteHeader(400)
					return
				} else {
					cid = c
					reader = bytes.NewReader(b)
				}
			} else if boundary, has := params["boundary"]; m == "multipart/form-data" && has {
				r := multipart.NewReader(req.Body, boundary)
				files := map[string]string{}
				u, err := uuid.NewRandom()
				if err != nil {
					res.WriteHeader(500)
					return
				}
				base := fmt.Sprintf("uuid://%s/", u.String())
				opts := ld.NewJsonLdOptions(base)
				opts.DocumentLoader = dl
				opts.ProcessingMode = ld.JsonLd_1_1
				var graph []interface{}
				for {
					if p, err := r.NextPart(); err == io.EOF {
						break
					} else if err != nil {
						res.WriteHeader(400)
						return
					} else if name := p.FormName(); name == req.URL.RawQuery {
						if doc, err := ld.DocumentFromReader(p); err != nil {
							res.WriteHeader(400)
							return
						} else if expanded, err := proc.Expand(doc, opts); err != nil {
							res.WriteHeader(400)
							return
						} else if flattened, err := proc.Flatten(expanded, nil, opts); err != nil {
							res.WriteHeader(400)
							return
						} else {
							graph = flattened.([]interface{})
						}
					} else if c, err := store(p); err != nil {
						res.WriteHeader(400)
						return
					} else {
						id := base + name
						uri := fmt.Sprintf("dweb:/ipfs/%s", c.String())
						files[id] = uri
					}
				}

				walk(graph, files)

				// Convert to RDF
				rdf, err := proc.Normalize(graph, options)
				if s, is := rdf.(string); !is || err != nil {
					res.WriteHeader(400)
					return
				} else if c, err := store(strings.NewReader(s)); err != nil {
					res.WriteHeader(500)
					return
				} else {
					cid = c
					reader = strings.NewReader(s)
				}
			} else {
				res.WriteHeader(415)
				return
			}

			if quads, graphs, err := styx.ParseMessage(reader); err != nil {
				res.WriteHeader(400)
			} else if r := db.HandleMessage(peerID.ID, cid, quads, graphs); res == nil {
				res.WriteHeader(204)
			} else {
				res.WriteHeader(200)

				encoder := json.NewEncoder(res)
				encoder.Encode(r)
			}
			return
		}
		fs.ServeHTTP(res, req)
	})

	log.Printf("Listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
