/*

This is just a helper package. Consider you imported the "github.com/luisfurquim/linsang/linsang"
package and make a query like this:


   set, err = hnd.Walk(linsang.Vertex{"coolboss"},linsang.In{"status"},linsang.In{"worksfor"},linsang.In{"follows"},linsang.Out{"follows"})


If you consider that the repetition of the namespace identifier (linsang) before the symbols is causing
too much visual polution, you could dot import this package like

import . "github.com/luisfurquim/linsang/linsang/setop"

And then, you may rewrite the code above this way:

   set, err = hnd.Walk(Vertex{"coolboss"},In{"status"},In{"worksfor"},In{"follows"},Out{"follows"})


*/
package setop

import (
   "github.com/luisfurquim/linsang"
)


// Wrapper to linsang.Vertex
type Vertex linsang.Vertex
func (op Vertex) Compute(s linsang.Store, n *linsang.Node) (linsang.Set, error) {
   return linsang.Vertex(op).Compute(s, n)
}


// Wrapper to linsang.Vertex
type V Vertex
func (op V) Compute(s linsang.Store, n *linsang.Node) (linsang.Set, error) {
   return linsang.Vertex(op).Compute(s, n)
}


// Wrapper to linsang.In
type In linsang.In
func (op In) Compute(s linsang.Store, n *linsang.Node) (linsang.Set, error) {
   return linsang.In(op).Compute(s, n)
}

// Wrapper to linsang.Out
type Out linsang.Out
func (op Out) Compute(s linsang.Store, n *linsang.Node) (linsang.Set, error) {
   return linsang.Out(op).Compute(s, n)
}

// Wrapper to linsang.Arc
type Arc linsang.Arc
func (op Arc) Compute(s linsang.Store, n *linsang.Node) (linsang.Set, error) {
   return linsang.Arc(op).Compute(s, n)
}

// Wrapper to linsang.Line
type Line linsang.Line
func (op Line) Compute(s linsang.Store, n *linsang.Node) (linsang.Set, error) {
   return linsang.Line(op).Compute(s, n)
}

// Wrapper to linsang.Edge
type Edge linsang.Edge
func (op Edge) Compute(s linsang.Store, n *linsang.Node) (linsang.Set, error) {
   return linsang.Edge(op).Compute(s, n)
}





