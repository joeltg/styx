# User Manual

## Table of Contents

- [Data Model](#Data_Model)
- [Installing](#Installing)
- [Writing Data](#Writing_Data)
- [Reading Data](#Reading_Data)

## Data Model

A basic understanding of RDF is assumed, which can be gathered from the [RDF Primer](https://www.w3.org/TR/rdf11-primer/).

Styx is a [Triplestore](https://en.wikipedia.org/wiki/Triplestore), which is a kind of graph database that stores and retrieves data in the form of subject-predicate-object "triples". At their most verbose, triples look like this:

```
_:b0 <http://schema.org/description> "The Empire State Building is a 102-story landmark in New York City." .
_:b0 <http://schema.org/name> "The Empire State Building" .
_:b0 <http://schema.org/image> <http://www.civil.usherbrooke.ca/cours/gci215a/empire-state-building.jpg> .
_:b0 <http://schema.org/geo> _:b1 .
_:b1 <http://schema.org/latitude> "4.075E1"^^<http://www.w3.org/2001/XMLSchema#double> .
_:b1 <http://schema.org/longitude> "7.398E1"^^<http://www.w3.org/2001/XMLSchema#double> .
```

This syntax is called [N-Triples](https://en.wikipedia.org/wiki/N-Triples) (`application/n-triples`), but there are also other, more developer-friendly ones like [JSON-LD](https://json-ld.org/):

```json
{
	"@context": { "@vocab": "http://schema.org/" },
	"name": "The Empire State Building",
	"description": "The Empire State Building is a 102-story landmark in New York City.",
	"image": {
		"@id": "http://www.civil.usherbrooke.ca/cours/gci215a/empire-state-building.jpg"
	},
	"geo": {
		"latitude": 40.75,
		"longitude": 73.98
	}
}
```

Both the JSON-LD and N-Triples here encode the same graph:

![graph](Screenshot_2019-09-27%20Styx%20Directory.png)

### URIs, Blank Nodes, and Literals

There are three kinds of terms that can occupy the three subject/predicate/object positions:

- **URIs** (`<http://schema.org/image>`, `<http://www.civil.usherbrooke.ca/...`) are used for absolute identifiers. All predicates must be URIs, typically taken from some vocabulary like [schema.org](https://schema.org/) or [foaf](http://xmlns.com/foaf/spec/).
- **Blank nodes** (`_:b0`, `_:b1`) are anonymous nodes. They're given labels that start with `_:` for serialization purposes, but those labels aren't significant outside each individual graph (i.e. if two different files both have a blank node labelled `_:b0`, they're **not** taken to mean the same thing). In JSON-LD, all JSON objects are implicitly blank nodes unless they're given an explicit `@id` property.
- **Literals** (`"The Empire State Building"`, `"4.075E1"^^<http://www.w3.org/2001/XMLSchema#double>`) are data primitives composed of a string _value_ and a URI _datatype_. The datatype is typically taken from the [XML Schema Definition Datatypes](https://www.w3.org/TR/xmlschema11-2/), but doesn't have to be. If there's no datatype given, it's assumed to be `<http://www.w3.org/2001/XMLSchema#string>`. Literals can only appear as the object of triples, never the subject or predicate.

### Datasets and Named Graphs

A set of triples (like the first example) is called an _RDF Graph_, and it's the simplest type of graph data "container". Another kind of "graph data container" is an _RDF Dataset_, which is an unordered collection of labeled RDF Graphs, plus an unlabeled "default graph" (again, the [RDF Primer](https://www.w3.org/TR/2014/REC-rdf11-concepts-20140225/#section-dataset) is the best place to learn about RDF concepts).

RDF Datasets are serialized by extending the N-Triples format into an "N-Quads" format (with MIME `application/n-quads`), where statement has an optional fourth term to indicate which graph within the dataset it belongs to.

```
_:b0 <http://www.w3.org/ns/prov#generatedAtTime> "2019-09-27T19:35:40.734Z"^^<http://www.w3.org/2001/XMLSchema#dateTime> .
_:b1 <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://xmlns.com/foaf/0.1/Person> _:b0 .
_:b1 <http://xmlns.com/foaf/0.1/knows> _:b2 _:b0 .
_:b1 <http://xmlns.com/foaf/0.1/name> "Andy Roddick" _:b0 .
_:b2 <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://xmlns.com/foaf/0.1/Person> _:b0 .
_:b2 <http://xmlns.com/foaf/0.1/name> "Tiger Woods" _:b0 .
```

Here, there's a fourth term `_:b0` on all of the lines except the first one, which states a property of the blank node as a subject. The meaning of this dataset is that there's a collection of graph data saying "a person named Andy Roddick knows a person named Tiger Woods", and that _this collection of data_ was generated on 09 September 2019.

In JSON-LD, named graphs are represented with a special `@graph` syntax:

```json
{
	"@context": {
		"@vocab": "http://xmlns.com/foaf/0.1/",
		"prov": "http://www.w3.org/ns/prov#",
		"xsd": "http://www.w3.org/2001/XMLSchema#"
	},
	"generatedAtTime": {
		"@value": "2019-09-27T19:35:40.734Z",
		"@type": "xsd:dateTime"
	},
	"@graph": {
		"@type": "Person",
		"name": "Andy Roddick",
		"knows": {
			"@type": "Person",
			"name": "Tiger Woods"
		}
	}
}
```

This interpretation of the "meaning" of the graph label - where the graph label is used as a means of reification - is not official. The RDF working group couldn't come to concensus on the semantics of RDF Datasets even though they standardized the syntax. More background is available in the working group note [On Semantics of RDF Datasets](https://www.w3.org/TR/rdf11-datasets/), but the emerging concensus in ensuing years is that this approach (using the graph label to refer to the graph) is the most sane interpretation.

As a result, **Styx restricts RDF Datasets to named graphs with blank graph names**. This means you can't insert datasets with URI graph labels - Styx won't allow it. This was a deliberate decision to strengthen these chosen semantics: the graph label referent is always the named graph itself, and only as represented in the dataset. This is less clear if graph labels can be URIs, which suggest global scope or universality.

### Content Addressing

Styx is built to play well with the distributed web, in particular [IPFS](https://ipfs.io/).

When you write graphs or datasets to Styx, it remembers which

## Installing

Styx is written in Go, so you need to [have Go installed](https://golang.org/doc/install). Then you should be able to install it with

```
go get github.com/underlay/styx
```

Styx is just a really fancy index for graph data that is ultimately stored in an IPFS repo. This IPFS repo doesn't have to be online or anything (although that's the responsible thing to do if you're storing data other people might want!), it just has to follow the [IPFS HTTP API spec](https://docs.ipfs.io/reference/api/http/) for file storage and retrieval.

You can do this [installing a prebuilt IPFS binary](https://docs.ipfs.io/guides/guides/install/), but you can also [build it from source](https://github.com/ipfs/go-ipfs#build-from-source) or even run [js-ipfs with Node](https://github.com/ipfs/go-ipfs#build-from-source). You'll need to initialize a repo somewhere and start the daemon in the background.

### Environment Variables

- `IPFS_HOST` (`localhost:5001`): the HTTP API endpoint of your IPFS daemon (just the hostname, port, and pathname - not the protocol)
- `STYX_PATH` (`/tmp/styx`): the absolute path of a directory that Styx will store data in (and look for existing data in)
- `STYX_PORT` (`8086`): the localhost port that Styx will serve its WebUI from

If you have an IPFS daemon running, then you shouldn't have to set any of these - just running `styx` will open a new database at `/tmp/styx` and start serving the WebUI at [`http://localhost:8086/`](http://localhost:8086/)

```
~ $ styx
2019/09/27 12:14:21 Opening badger database at /tmp/styx
badger 2019/09/27 12:14:21 INFO: All 1 tables opened in 0s
badger 2019/09/27 12:14:21 INFO: Replaying file id: 0 at offset: 7431
badger 2019/09/27 12:14:21 INFO: Replay took: 1.645545ms
badger 2019/09/27 12:14:21 DEBUG: Value log discard stats empty
2019/09/27 12:14:21 Listening on port 8086
```

Styx is build on the [Badger](https://github.com/dgraph-io/badger) key/value store, so everything that it writes to `STYX_PATH` will be Badger database files.

## Writing Data

The simplest way to write data to Styx is to send a POST to `http://localhost:8086/` with `Content-Type: application/ld+json` or `Content-Type: application/n-quads`.

```
~ $ cat esb.nt
_:b0 <http://schema.org/description> "The Empire State Building is a 102-story landmark in New York City." .
_:b0 <http://schema.org/name> "The Empire State Building" .
_:b0 <http://schema.org/image> <http://www.civil.usherbrooke.ca/cours/gci215a/empire-state-building.jpg> .
_:b0 <http://schema.org/geo> _:b1 .
_:b1 <http://schema.org/latitude> "4.075E1"^^<http://www.w3.org/2001/XMLSchema#double> .
_:b1 <http://schema.org/longitude> "7.398E1"^^<http://www.w3.org/2001/XMLSchema#double> .
~ $ curl --data-binary @esb.nt -H 'Content-Type: application/n-quads' http://localhost:8086/
null
```

Then if we look at http://localhost:8086/directory/, we'll see that file listed (by content-address!) and we can view its contents:

![directory](Screen%20Shot%202019-09-27%20at%204.40.48%20PM.png)
![download](Screen%20Shot%202019-09-27%20at%204.40.54%20PM.png)
![view](Screen%20Shot%202019-09-27%20at%204.33.23%20PM.png)

## Reading Data

The reason Styx exists is its query interface, which is designed to be both easier to use for simple use cases and more powerful for larger ones than SPARQL, openCypher, GQL, etc. It's still a work in progress.

The query language is built around the concept of subgraph matching. Queries are expressed with the same syntax as the data is: using RDF graphs and datasets. Specifically, whenever Styx receives a dataset with a graph label that is given an RDF type of `http://underlay.mit.edu/ns#Query`, it interprests the associated named graph as a query. Mainly, this means that instead of treating blank nodes as anonyous values, Styx will interpret them as variables to be satisfied, and will search it indices for assignments such that the resulting graph is "true" - that is, the graph exists is a subset of the merged graph of all datasets in the database.

```
~ $ cat query.nt
_:q <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://underlay.mit.edu/ns#Query> .
_:b0 <http://schema.org/description> _:b2_ _:q .
_:b0 <http://schema.org/name> "The Empire State Building" _:q .
_:b0 <http://schema.org/geo> _:b1 _:q .
_:b1 <http://schema.org/latitude> _:b3_ _:q .
_:b1 <http://schema.org/longitude> _:b4 _:q .
~ $ curl --data-binary @query.nt -H 'Content-Type: application/n-quads' http://localhost:8086/
<ul:/ipfs/QmShB9mYxbcKfNGpnFxJGML6vCLYCF6cEMZH2KZTCUPqww#_:b0> <http://schema.org/description> "The Empire State Building is a 102-story landmark in New York City." _:b0 .
<ul:/ipfs/QmShB9mYxbcKfNGpnFxJGML6vCLYCF6cEMZH2KZTCUPqww#_:b0> <http://schema.org/geo> <ul:/ipfs/QmShB9mYxbcKfNGpnFxJGML6vCLYCF6cEMZH2KZTCUPqww#_:b1> _:b0 .
<ul:/ipfs/QmShB9mYxbcKfNGpnFxJGML6vCLYCF6cEMZH2KZTCUPqww#_:b0> <http://schema.org/name> "The Empire State Building" _:b0 .
<ul:/ipfs/QmShB9mYxbcKfNGpnFxJGML6vCLYCF6cEMZH2KZTCUPqww#_:b1> <http://schema.org/latitude> "4.075E1"^^<http://www.w3.org/2001/XMLSchema#double> _:b0 .
<ul:/ipfs/QmShB9mYxbcKfNGpnFxJGML6vCLYCF6cEMZH2KZTCUPqww#_:b1> <http://schema.org/longitude> "7.398E1"^^<http://www.w3.org/2001/XMLSchema#double> _:b0 .
_:b0 <http://underlay.mit.edu/ns#instanceOf> <ul:/ipfs/QmeUKcMCigv1bRsSyiSzEVTqBMoabuLk1hpSVT91PoVT5F#_:q> .
```