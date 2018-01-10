package linsangmysql

import (
   "database/sql"
   "github.com/luisfurquim/linsang"
)

func (op Out) Compute(store linsang.Store, srcNode *linsang.Node) (linsang.Set, error) {
   var err           error
   var pred          string
   var preds       []string
   var predInt       int64
   var nodeType      uint8
   var res          *MySQLSet
   var rows         *sql.Rows
   var hnd          *Handle
   var chkArcType    func(bool) bool

   hnd, res, nodeType, preds, err = operationPreCheck(store,srcNode,op.Predicate)
   if err != nil {
      if err == ErrNotFound {
         Goose.Read.Logf(1,"No source node %#v node", srcNode)
         return res, nil
      }
      Goose.Read.Logf(1,"Error searching object %#v node: %s", op, err)
      return nil, err
   }

   chkArcType = func(arctype bool) bool {
      return arctype != linsang.Directed
   }

   for _, pred = range preds {
      predInt, rows, err = operationDefineQuery(hnd, pred, srcNode, nodeType, hnd.GetLinkSbjQuery, hnd.GetLinkSPHashQuery,2)
      if err != nil {
         if err == ErrNotFound {
            return res, nil
         }
         Goose.Read.Logf(1,"Error searching object %#v node: %s", op, err)
         return nil, err
      }

      defer rows.Close()

      Goose.Read.Logf(6,"Equals links: %#v", rows)

      err = operationExec(hnd, rows, pred, predInt, srcNode, nodeType, res, linsang.Subject, chkArcType)
      if err != nil {
         Goose.Read.Logf(1,"Error fetching value from Output arc: %s", err)
         return nil, err
      }
   }

   return res, nil
}
