package linsang

func (op Vertex) Compute(store Store, srcNode *Node) (Set, error) {
   if srcNode == nil {
      return store.SetOperationVtx()
   }

   return store.SetOperationVtx(srcNode)
}

