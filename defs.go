/*

Inspired on Cayley (https:///github.com/google/cayley/), Linsang implements a graph database engine for the Go programming language (https://golang.org/).



Features

* Vertex data may be of type bool, int64, float64, string and time.Time.


* Searches may use operators like greater than, less than, etc.


* New operators may be custom defined in external packages.


* Graphs may be directed, undirected or even mix edges of either type.



Limitations

* The graph does not need to be connected, but each vertex MUST have, at least one edge.


* It's alpha code yet.

*

*/
package linsang

import (
   "time"
   "github.com/luisfurquim/goose"
)

/*
const (
   In    byte = iota + 1
   Out
   Arc
   _
   Line
   _
   Edge
)
*/

const (
   IntNode  = iota
   RealNode
   StringNode
   BoolNode
   DateNode
)

const (
   Undirected bool = false
   Directed   bool = true
)

const (
   Subject int = iota
   Predicate
   Object
   Label
   SubObj
)

type Node struct {
   Value interface{}
   Id    int64
   Via []interface{}
}

type Nodes []Node

type SetOperation interface {
   Compute(Store, *Node) (Set, error)
}

type Set interface {
   Next()      bool
   GetNode()   (*Node, error)
   Get()       (interface{}, error)
   GetBool()   (bool, error)
   GetInt()    (int64, error)
   GetReal()   (float64, error)
   GetString() (string, error)
   GetDate()   (time.Time, error)
   Close()     error
   Walk(Store, SetOperation, ...SetOperation)  (Set, error)
//   WalkHistory(...SetOperation)  (*Set, error)
   All()       ([]interface{}, error)
   AllBool()   ([]bool, error)
   AllInt()    ([]int64, error)
   AllReal()   ([]float64, error)
   AllString() ([]string, error)
   AllDate()   ([]time.Time, error)
}

type Store interface {
/*
Executes initialization procedures. If some error occurs when initializing it tries to recreate the storage.
If failure persists then it returns the error.
*/
   Open() error

/*
Closes the underlying storage infrastructure
*/
   Close() error


/*
Creates whatever structures the underlying storage needs to begin operating with the graph.
Any already existent content/structure must be preserved and the error returned is nil in such case.
*/
   Create() error



/*
Creates an edge connecting 2 vertices. If some or both vertices does not exist, it/they is/are created.
No vertex is created with redundancy.

If the arctype is linsang.Undirected, it's a line and there is no difference in order of the vertices,
e. g.
   Handle.Link("bob","workswith","alice",nil,linsang.Undirected)
is the same as
   Handle.Link("alice","workswith","bob",nil,linsang.Undirected)
and creates this:
   +-----+  workswith  +-------+
   | bob |-------------| alice |
   +-----+             +-------+

If the arctype is linsang.Directed, it's an arc and then the subject vertex is the source of the arc and
the object vertex is the target of the arc, e. g.
   Handle.Link("bob","follows","alice",nil,linsang.Directed)
means that bob follows alice and alice is followed by bob and creates this:
   +-----+  follows  +-------+
   | bob |---------->| alice |
   +-----+           +-------+


On the other hand, the code below
   Handle.Link("alice","follows","bob",nil,linsang.Directed)
means that alice follows bob and bob is followed by alice and creates this:
   +-----+  follows  +-------+
   | bob |<----------| alice |
   +-----+           +-------+
*/
   Link(subject interface{}, predicate string, object interface{}, label interface{}, arctype bool) error



/*
Walks through the vertices of the graph.
It starts from all nodes and follow the edges according to the operands.

Consider the following graph:

                        worksfor
       +--------------------------------------+
       |                                      |
   +-------+  follows                         |
   | emily |-------------+                    |
   +-------+             |                    |
                         V                    V
   +-----+  follows  +-------+  worksfor  +---------+  status  +----------+
   | bob |---------->| alice |----------->| charlie |--------->| coolboss |
   +-----+           +-------+            +---------+          +----------+
      |
      |    worksfor  +-------+  status  +---------+
      +------------->| david |--------->| badboss |
                     +-------+          +---------+

The following call to Store.Walk

   set, err = Store.Walk(Vertex{[]string{"alice","bob"}},Out{"worksfor"},Out{"status"})

results in badboss and coolboss.


The following call to Store.Walk

   set, err = Store.Walk(Vertex{"coolboss"},In{"Status"},In{"worksfor"})

results in alice and emily.


The following call to Store.Walk

   set, err = Store.Walk(Vertex{"alice"},In{"follows"},Out{"worksfor"})

results in charlie, david.


The following call to Store.Walk

   set, err = Store.Walk(Vertex{"coolboss"},In{"status"},In{"worksfor"},In{"follows"},Out{"follows"})

results in alice.
*/
   Walk(SetOperation, ...SetOperation)  (Set, error)


/*
SetOperation is the low-level, storage dependant, compute method for the most basic operators that computes the check of relations
between nodes. Applications should avoid using it directly and prefer to call the Walk method of either linsang.Store or linsang.Set.
*/
   SetOperation(srcNode *Node, subject, predicate, object, label interface{}, chkArcType func(bool) bool, needSubj, needObj bool) (Set, error)

   SetOperationVtx(values ...interface{}) (Set, error)
}


type Equ   struct {
   Predicate string
   Value     interface{}
   Direction byte
}


/*
Value, one of:

* a bool, any integer, any float, non empty string a non empty []byte or time.Time: The vertex value to search

* a list of values of the types above: The vertices to search

If computed via linsang.Store.Walk object it searches on all nodes.

If computed via linsang.Set.Walk object it searches only in the nodes belonging to the set.

*/
type Vertex struct {
   Value     interface{}
}

type V Vertex

/*
Predicate is one of:

* an empty string: All arcs pointing into this node

* a non empty string: The arc name to follow into this node

* a list of strings: The arcs to follow into this node

Same as Out, but in the other direction.

Consider the following graph:

                        worksfor
       +--------------------------------------+
       |                                      |
   +-------+  follows                         |
   | emily |-------------+                    |
   +-------+             |                    |
                         V                    V
   +-----+  follows  +-------+  worksfor  +---------+  status  +----------+
   | bob |---------->| alice |----------->| charlie |--------->| coolboss |
   +-----+           +-------+            +---------+          +----------+
      |
      |    worksfor  +-------+  status  +---------+
      +------------->| david |--------->| badboss |
                     +-------+          +---------+

If computed via Store.Walk object it starts from all nodes and follow the arcs with predicates defined by the operand(s) to their subjects.
The following line

   set, err = Store.Walk(In{"Status"})

results in charlie and david.



If computed via Set.Walk object it starts from the nodes on this set and follow the arcs with predicates defined by the operand(s) to their subjects.
The following lines

   set, err = Store.Walk(Vertex{"coolboss"})
   set.Walk(In{"Status"})

results in charlie. And the lines

   set, err = Store.Walk(Vertex{"coolboss"})
   set.Walk(In{"Status"},In{"worksfor"})

results in alice and emily.





*/
type In struct {
   Predicate interface{}
}




/*
Predicate is one of:

* an empty string: All arcs pointing out from this node

* a non empty string: The arc name to follow out from this node

* a list of strings: The arcs to follow out from this node

Same as In, but in the other direction.

Consider the following graph:

                        worksfor
       +--------------------------------------+
       |                                      |
   +-------+  follows                         |
   | emily |-------------+                    |
   +-------+             |                    |
                         V                    V
   +-----+  follows  +-------+  worksfor  +---------+  status  +----------+
   | bob |---------->| alice |----------->| charlie |--------->| coolboss |
   +-----+           +-------+            +---------+          +----------+
      |
      |    worksfor  +-------+  status  +---------+
      +------------->| david |--------->| badboss |
                     +-------+          +---------+



If computed via Store.Walk object it starts from all nodes and follow the outgoing arcs with predicates defined by the operand(s) to their subjects.
The following line

   set, err = Store.Walk(Out{"Status"})

results in coolboss and badboss.




If computed via Set.Walk object it starts from the nodes on this set and follow the arcs with predicates defined by the operand(s) to their subjects.
The following lines

   set, err = Store.Walk(Vertex{"charlie"})
   set.Walk(Out{"Status"})

results in coolboss. And the lines

   set, err = Store.Walk(Vertex{[]string{"alice","emily"}})
   set.Walk(Out{"worksfor"},Out{"Status"})

results in coolboss.

*/
type Out struct {
   Predicate interface{}
}

/*
Predicate is one of:

* an empty string: All arcs pointing into or out from this node

* a non empty string: The arc name to follow into or out from this node

* a list of strings: The arcs to follow into or out from this node

Same as In/Out, but both directions.

If computed via linsang.Store.Walk object it starts from all nodes and follow the arcs with predicates defined by the operand(s) to their related vertices.

If computed via linsang.Set.Walk object it starts from the nodes on this set and follow the arcs with predicates defined by the operand(s) to their related vertices.

*/
type Arc struct {
   Predicate interface{}
}

/*
Predicate is one of:

* an empty string: All lines connected to this node

* a non empty string: The line name to follow from this node

* a list of strings: The lines to follow from this node

It retrieves vertices connected via undirected relations.

If computed via linsang.Store.Walk object it starts from all nodes and follow the lines with predicates defined by the operand(s) to their related vertices.

If computed via linsang.Set.Walk object it starts from the nodes on this set and follow the lines with predicates defined by the operand(s) to their related vertices.

*/
type Line struct {
   Predicate interface{}
}

/*
Predicate is one of:

* an empty string: All lines and arcs pointing into or out from this node

* a non empty string: The line or arc name to follow into or out from this node

* a list of strings: The lines or arcs to follow into or out from this node

It gives any relation: undirected (Line) or directed (Arc of all directions, e. g. In or Out).

If computed via linsang.Store.Walk object it starts from all nodes and follow the arcs with predicates defined by the operand(s) to their related vertices.

If computed via linsang.Set.Walk object it starts from the nodes on this set and follow the arcs with predicates defined by the operand(s) to their related vertices.

*/
type Edge struct {
   Predicate interface{}
}




type Diff  interface{}
type Grt   interface{}
type GEq   interface{}
type Less  interface{}
type LEq   interface{}
type And   interface{}
type Or    interface{}
type Xor   interface{}
//type Not   struct{}
type Match string




var Goose goose.Alert
