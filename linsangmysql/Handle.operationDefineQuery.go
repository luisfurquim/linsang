package linsangmysql

import (
   "database/sql"
   "github.com/luisfurquim/linsang"
)

func (hnd *Handle) operationDefineQuery(pred string, srcNode *linsang.Node, nodeType uint8, vertexQuery, hashQuery  *sql.Stmt, qtParm int) (int64, *sql.Rows, error) {
   var rows    *sql.Rows
   var hash     int64
   var predInt  int64
   var err      error

   if pred == "" {
      if srcNode == nil {
         return 0, nil, ErrOpNotFound
      }
      if qtParm==2 {
         rows, err = vertexQuery.Query(srcNode.Id, nodeType)
      } else if qtParm==4 {
         rows, err = vertexQuery.Query(srcNode.Id, nodeType, srcNode.Id, nodeType)
      }
   } else {
      predInt, err =  hnd.getNodeString(pred)
      if err != nil {
         return 0, nil, err
      }

      if srcNode != nil { // We have a left operand
         hash = int64((hash64shift(uint64(srcNode.Id)) ^ hash64shift(uint64(predInt)) ^ hash64shift(uint64(nodeType)) ^ hash64shift(uint64(predInt<<32))) & 0x7fffffffffffffff)

         Goose.Read.Logf(6,"Checking link equity")

         // Fetch the link by hash composition
         if qtParm==2 {
            rows, err = hashQuery.Query(hash)
         } else if qtParm==4 {
            rows, err = hashQuery.Query(hash,hash)
         }
      } else { // No left operand, just get all matching predicates
         Goose.Read.Logf(6,"Searching by predicate, only")
         rows, err = hnd.GetLinkPredQuery.Query(predInt)
      }
   }

   return predInt, rows, err
}
