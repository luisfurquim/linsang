package linsangmysql

import (
   "github.com/luisfurquim/linsang"
)

/*

Walks through the vertices of the graph.
It starts from the selected set of nodes and follows the edges according to the operands.

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

The following call to Handle.Walk

   set, err = hnd.Walk(Vertex{[]string{"alice","bob"}})
   set.Walk(Out{"worksfor"},Out{"status"})

the call to set.Walk results in badboss and coolboss.


The following call to Handle.Walk

   set, err = hnd.Walk(Vertex{"coolboss"})
   set.Walk(In{"Status"},In{"worksfor"})

the call to set.Walk results in alice and emily.


The following call to Handle.Walk

   set, err = hnd.Walk(Vertex{"alice"})
   set.Walk(In{"follows"},Out{"worksfor"})

the call to set.Walk results in charlie, david.


The following call to Handle.Walk

   set, err  = hnd.Walk(Vertex{"coolboss"},In{"status"})
   set2, err = set.Walk(In{"worksfor"},In{"follows"})
   set2.Walk(Out{"follows"})

the call to set.Walk results in emily and the call to set2.Walk results in alice.


*/
func (set *MySQLSet) Walk(hnd linsang.Store, op linsang.SetOperation, ops ...linsang.SetOperation)  (linsang.Set, error) {
   var i              int
   var err            error
   var finalResult    MySQLSet
   var partialResult  linsang.Set
   var node          *linsang.Node

   finalResult.curr   = -1
   finalResult.vertex = make([]linsang.Node,0,16)

   for set.Next() {
      node, err = set.GetNode()
      if err != nil {
         return nil, err
      }
      partialResult, err = op.Compute(hnd, node)
      if err != nil {
         return nil, err
      }

      if len(ops) > 0 {
         partialResult, err = partialResult.Walk(hnd,ops[0], ops[1:]...)
         if err != nil {
            return nil, err
         }
      }
      if len(finalResult.vertex) > 0 {
         for partialResult.Next() {
            node, err = partialResult.GetNode()
            if err != nil {
               return nil, err
            }
            for i=0; i<len(finalResult.vertex); i++ {
               if finalResult.vertex[i].Value == node.Value {
                  break
               }
            }
            if i >= len(finalResult.vertex) {
               finalResult.vertex = append(finalResult.vertex,*node)
               finalResult.len++
            }
         }
      } else {
         for partialResult.Next() {
            node, err = partialResult.GetNode()
            if err != nil {
               return nil, err
            }
            finalResult.vertex = append(finalResult.vertex,*node)
            finalResult.len++
         }
      }
   }

   return &finalResult, nil
}