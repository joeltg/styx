package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	ipfs "github.com/ipfs/go-ipfs-api"
	ld "github.com/piprate/json-gold/ld"

	styx "github.com/underlay/styx/db"
	loader "github.com/underlay/styx/loader"
)

var sampleData = []byte(`{
	"@context": {
		"@vocab": "http://schema.org/",
		"xsd": "http://www.w3.org/2001/XMLSchema#",
		"prov": "http://www.w3.org/ns/prov#",
		"prov:generatedAtTime": {
			"@type": "xsd:dateTime"
		},
		"birthDate": {
			"@type": "xsd:date"
		}
	},
	"prov:generatedAtTime": "2019-07-24T16:46:05.751Z",
	"@graph": {
		"@type": "Person",
		"name": ["John Doe", "Johnny Doe"],
		"birthDate": "1996-02-02",
		"knows": {
			"@id": "http://people.com/jane",
			"@type": "Person",
			"name": "Jane Doe",
			"familyName": { "@value": "Doe", "@language": "en" },
			"birthDate": "1995-01-01"
		}
	}
}`)

func TestIngest(t *testing.T) {
	// Replace at your leisure
	sh := ipfs.NewShell(defaultHost)

	if !sh.IsUp() {
		t.Error("IPFS Daemon not running")
		return
	}

	peerID, err := sh.ID()
	if err != nil {
		t.Error(err)
		return
	}

	// Remove old db
	fmt.Println("removing path", styx.DefaultPath)
	if err := os.RemoveAll(styx.DefaultPath); err != nil {
		t.Error(err)
		return
	}

	dl := loader.NewShellDocumentLoader(sh)

	store := styx.NewHTTPDocumentStore(sh)

	db, err := styx.OpenDB(styx.DefaultPath, peerID.ID, dl, store)
	if err != nil {
		t.Error(err)
		return
	}

	defer db.Close()

	var data map[string]interface{}
	if err = json.Unmarshal(sampleData, &data); err != nil {
		t.Error(err)
		return
	}

	if err = db.IngestJSONLd(data); err != nil {
		t.Error(err)
		return
	}

	if err = db.Log(); err != nil {
		t.Error(err)
	}
}

func testQuery(query []byte) error {
	// Replace at your leisure
	sh := ipfs.NewShell(defaultHost)

	if !sh.IsUp() {
		return fmt.Errorf("IPFS Daemon not running")
	} else if err := os.RemoveAll(styx.DefaultPath); err != nil {
		return err
	}

	peerID, err := sh.ID()
	if err != nil {
		return err
	}

	dl := loader.NewShellDocumentLoader(sh)

	store := styx.NewHTTPDocumentStore(sh)

	db, err := styx.OpenDB(styx.DefaultPath, peerID.ID, dl, store)
	if err != nil {
		return err
	}

	defer db.Close()

	var data map[string]interface{}
	if err := json.Unmarshal(sampleData, &data); err != nil {
		return err
	}

	if err := db.IngestJSONLd(data); err != nil {
		return err
	}

	db.Log()

	var queryData map[string]interface{}
	if err := json.Unmarshal(query, &queryData); err != nil {
		return err
	}

	fmt.Println("got query data")

	proc := ld.NewJsonLdProcessor()
	stringOptions := styx.GetStringOptions(dl)
	rdf, err := proc.ToRDF(queryData, stringOptions)
	if err != nil {
		return err
	}

	size := uint32(len(rdf.(string)))
	fmt.Println("About to put")
	mh, err := store.Put(strings.NewReader(rdf.(string)))
	if err != nil {
		return err
	}

	fmt.Println("Hash", mh.B58String())
	reader, err := store.Get(mh)
	quads, _, err := styx.ParseMessage(reader)
	if err != nil {
		return err
	}

	fmt.Println("--- query graph ---")
	for _, quad := range quads {
		fmt.Printf(
			"  %s %s %s",
			quad.Subject.GetValue(),
			quad.Predicate.GetValue(),
			quad.Object.GetValue(),
		)
		if quad.Graph != nil {
			fmt.Print(" " + quad.Graph.GetValue())
		}
		fmt.Print("\n")
	}

	result, err := db.HandleMessage(mh, size)
	if err != nil {
		return err
	}
	api := ld.NewJsonLdApi()
	normalized, err := api.Normalize(result, stringOptions)
	fmt.Println("Result:")
	fmt.Println(normalized)
	s, err := sh.Add(strings.NewReader(normalized.(string)), ipfs.RawLeaves(true))
	fmt.Printf("http://localhost:8000?%s\n", s)
	return err
}

func TestSPO(t *testing.T) {
	if err := testQuery([]byte(`{
	"@context": { "@vocab": "http://schema.org/" },
	"@type": "http://underlay.mit.edu/ns#Query",
	"@graph": {
		"@id": "http://people.com/jane",
		"name": { }
	}
}`)); err != nil {
		t.Error(err)
	}
}

func TestOPS(t *testing.T) {
	if err := testQuery([]byte(`{
	"@context": { "@vocab": "http://schema.org/" },
	"@type": "http://underlay.mit.edu/ns#Query",
	"@graph": {
		"name": "Jane Doe"
	}
}`)); err != nil {
		t.Error(err)
	}
}

func TestSimpleQuery(t *testing.T) {
	if err := testQuery([]byte(`{
	"@context": {
		"@vocab": "http://schema.org/"
	},
	"@type": "http://underlay.mit.edu/ns#Query",
	"@graph": {
		"@type": "Person",
		"birthDate": { },
		"knows": {
			"name": "Jane Doe"
		}
	}
}`)); err != nil {
		t.Error(err)
	}
}

// func TestEntityQuery(t *testing.T) {
// 	if err := testQuery([]byte(`{
// 	"@context": {
// 		"@vocab": "http://schema.org/",
// 		"prov": "http://www.w3.org/ns/prov#",
// 		"u": "http://underlay.mit.edu/ns#"
// 	},
// 	"@type": "u:Query",
// 	"@graph": {
// 		"@type": "prov:Entity",
// 		"u:satisfies": {
// 			"@graph": {
// 				"@type": "Person",
// 				"birthDate": { },
// 				"knows": {
// 					"name": "Jane Doe"
// 				}
// 			}
// 		}
// 	}
// }`)); err != nil {
// 		t.Error(err)
// 	}
// }

func TestBundleQuery(t *testing.T) {
	if err := testQuery([]byte(`{
	"@context": {
		"@vocab": "http://schema.org/",
		"dcterms": "http://purl.org/dc/terms/",
		"prov": "http://www.w3.org/ns/prov#",
		"u": "http://underlay.mit.edu/ns#",
		"u:index": { "@container": "@list" },
		"u:domain": { "@container": "@list" }
	},
	"@type": "u:Query",
	"@graph": {
		"@type": "prov:Entity",
		"dcterms:extent": 1,
		"u:domain": [],
		"u:index": [],
		"u:satisfies": {
			"@graph":  {
				"@type": "Person",
				"birthDate": { },
				"knows": {
					"name": "Jane Doe"
				}
			}
		}
	}
}`)); err != nil {
		t.Error(err)
	}
}

func TestGraphQuery(t *testing.T) {
	if err := testQuery([]byte(`{
	"@context": {
		"dcterms": "http://purl.org/dc/terms/",
		"prov": "http://www.w3.org/ns/prov#",
		"u": "http://underlay.mit.edu/ns#",
		"u:index": { "@container": "@list" }
	},
	"@type": "u:Query",
	"@graph": {
		"@type": "prov:Bundle",
		"dcterms:extent": 3,
		"u:index": [],
		"u:enumerates": {
			"@graph": {
				"@type": "u:Graph"
			}
		}
	}
}`)); err != nil {
		t.Error(err)
	}
}

func TestGraphIndexQuery(t *testing.T) {
	if err := testQuery([]byte(`{
	"@context": {
		"dcterms": "http://purl.org/dc/terms/",
		"prov": "http://www.w3.org/ns/prov#",
		"rdf": "http://www.w3.org/1999/02/22-rdf-syntax-ns#",
		"u": "http://underlay.mit.edu/ns#",
		"u:index": { "@container": "@list" }
	},
	"@type": "u:Query",
	"@graph": {
		"@type": "prov:Bundle",
		"dcterms:extent": 3,
		"u:index": {
			"@id": "_:b0",
			"rdf:value": { "@id": "ul:/ipfs/QmRyaXPZpXxXBcdrikHTjnLr2w6rQK9bChsB7V1bUZv1er#_:c14n0" }
		},
		"u:enumerates": {
			"@graph": {
				"@id": "_:b0",
				"@type": "u:Graph"
			}
		}
	}
}`)); err != nil {
		t.Error(err)
	}
}

func TestDomainQuery(t *testing.T) {
	if err := testQuery([]byte(`{
	"@context": {
		"@vocab": "http://schema.org/",
		"dcterms": "http://purl.org/dc/terms/",
		"prov": "http://www.w3.org/ns/prov#",
		"u": "http://underlay.mit.edu/ns#",
		"u:index": { "@container": "@list" }
	},
	"@type": "u:Query",
	"@graph": {
		"@type": "prov:Bundle",
		"dcterms:extent": 3,
		"u:index": { "@id": "_:b0" },
		"u:enumerates": {
			"@graph": {
				"@id": "_:b0",
				"@type": "Person",
				"name": { "@id": "_:b1" }
			}
		}
	}
}`)); err != nil {
		t.Error(err)
	}
}

func TestIndexQuery(t *testing.T) {
	if err := testQuery([]byte(`{
	"@context": {
		"@vocab": "http://schema.org/",
		"dcterms": "http://purl.org/dc/terms/",
		"prov": "http://www.w3.org/ns/prov#",
		"rdf": "http://www.w3.org/1999/02/22-rdf-syntax-ns#",
		"u": "http://underlay.mit.edu/ns#",
		"u:index": { "@container": "@list" }
	},
	"@type": "u:Query",
	"@graph": {
		"@type": "prov:Bundle",
		"dcterms:extent": 2,
		"u:index": {
			"@id": "_:b0",
			"rdf:value": { "@id": "ul:/ipfs/QmRyaXPZpXxXBcdrikHTjnLr2w6rQK9bChsB7V1bUZv1er#_:c14n1" }
		},
		"u:enumerates": {
			"@graph": {
				"@id": "_:b0",
				"@type": "Person",
				"name": { "@id": "_:b1" }
			}
		}
	}
}`)); err != nil {
		t.Error(err)
	}
}

func TestIndexQuery2(t *testing.T) {
	if err := testQuery([]byte(`{
	"@context": {
		"dcterms": "http://purl.org/dc/terms/",
		"prov": "http://www.w3.org/ns/prov#",
		"u": "http://underlay.mit.edu/ns#",
		"u:index": { "@container": "@list" }
	},
	"@type": "u:Query",
	"@graph": {
		"@type": "prov:Bundle",
		"dcterms:extent": 2,
		"u:index": { "@id": "_:b0" },
		"u:enumerates": {
			"@graph": {
				"@id": "ul:/ipfs/QmRyaXPZpXxXBcdrikHTjnLr2w6rQK9bChsB7V1bUZv1er#_:c14n1",
				"_:b0": {}
			}
		}
	}
}`)); err != nil {
		t.Error(err)
	}
}

func TestNT(t *testing.T) {
	// Replace at your leisure
	sh := ipfs.NewShell(defaultHost)
	if !sh.IsUp() {
		t.Error("IPFS Daemon not running")
		return
	}

	peerID, err := sh.ID()
	if err != nil {
		t.Error(err)
		return
	}

	// Remove old db
	fmt.Println("removing path", styx.DefaultPath)
	if err := os.RemoveAll(styx.DefaultPath); err != nil {
		t.Error(err)
		return
	}

	dl := loader.NewShellDocumentLoader(sh)
	store := styx.NewHTTPDocumentStore(sh)

	db, err := styx.OpenDB(styx.DefaultPath, peerID.ID, dl, store)
	if err != nil {
		t.Error(err)
		return
	}

	defer db.Close()

	// names, err := openFile("/samples/nt/names.json", dl, store)
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }

	// if err = db.IngestJSONLd(names, dl, store); err != nil {
	// 	t.Error(err)
	// 	return
	// }

	individuals, err := openFile("/samples/nt/individuals.json", dl)
	if err != nil {
		t.Error(err)
		return
	}

	if err = db.IngestJSONLd(individuals); err != nil {
		t.Error(err)
		return
	}

	query, err := openFile("/samples/nt/small.json", dl)
	if err != nil {
		t.Error(err)
		return
	}

	documentLoader := loader.NewShellDocumentLoader(sh)

	proc := ld.NewJsonLdProcessor()
	stringOptions := styx.GetStringOptions(documentLoader)
	rdf, err := proc.ToRDF(query, stringOptions)

	if err != nil {
		t.Error(err)
	}

	quads, _, err := styx.ParseMessage(strings.NewReader(rdf.(string)))
	if err != nil {
		t.Error(err)
	}

	fmt.Println("--- query graph ---")
	for _, quad := range quads {
		fmt.Printf(
			"  %s %s %s\n",
			quad.Subject.GetValue(),
			quad.Predicate.GetValue(),
			quad.Object.GetValue(),
		)
	}

	reader := strings.NewReader(rdf.(string))
	size := uint32(len(rdf.(string)))
	mh, err := store.Put(reader)
	if err != nil {
		t.Error(err)
	}

	result, err := db.HandleMessage(mh, size)
	api := ld.NewJsonLdApi()
	normalized, err := api.Normalize(result, stringOptions)
	fmt.Println("Result:")
	fmt.Println(normalized)
}

func openFile(path string, dl ld.DocumentLoader) (doc map[string]interface{}, err error) {
	var dir string
	if dir, err = os.Getwd(); err != nil {
		return
	}

	var data []byte
	if data, err = ioutil.ReadFile(dir + path); err != nil {
		return
	}

	err = json.Unmarshal(data, &doc)
	return
}
