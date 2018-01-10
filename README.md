Inspired on Cayley (https://github.com/google/cayley/), Linsang implements a graph database engine for the Go programming language (https://golang.org/).

[![GoDoc](https://godoc.org/github.com/luisfurquim/linsang?status.png)](http://godoc.org/github.com/luisfurquim/linsang)


Features

* Vertex data may be of type bool, int64, float64, string and time.Time.


* Searches may use operators like greater than, less than, etc.


* New operators may be custom defined in external packages.


* Graphs may be directed, undirected or even mix edges of either type.



Limitations

* The graph does not need to be connected, but each vertex MUST have, at least one edge.


* It's alpha code yet.


