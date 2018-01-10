package linsangmysql

import (
   "github.com/luisfurquim/linsang"
)

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

The following call to Handle.Walk

   set, err = hnd.Walk(Vertex{[]string{"alice","bob"}},Out{"worksfor"},Out{"status"})

results in badboss and coolboss.


The following call to Handle.Walk

   set, err = hnd.Walk(Vertex{"coolboss"},In{"Status"},In{"worksfor"})

results in alice and emily.


The following call to Handle.Walk

   set, err = hnd.Walk(Vertex{"alice"},In{"follows"},Out{"worksfor"})

results in charlie, david.


The following call to Handle.Walk

   set, err = hnd.Walk(Vertex{"coolboss"},In{"status"},In{"worksfor"},In{"follows"},Out{"follows"})

results in alice.


*/
func (hnd *Handle) Walk(op linsang.SetOperation, ops ...linsang.SetOperation)  (linsang.Set, error) {
   var err            error
   var finalResult    linsang.Set

   finalResult, err = op.Compute(hnd, nil)
   if err != nil {
      return nil, err
   }

   if len(ops) > 0 {
      finalResult, err = finalResult.Walk(hnd,ops[0], ops[1:]...)
      return finalResult, err
   }

   return finalResult, nil
}
