package linsang

func (op Out) Compute(store Store, srcNode *Node) (Set, error) {
   var chkArcType    func(bool) bool

   chkArcType = func(arctype bool) bool {
      return arctype != Directed
   }

   return store.SetOperation(srcNode, nil, op.Predicate, nil, nil, chkArcType, false, true)
}
