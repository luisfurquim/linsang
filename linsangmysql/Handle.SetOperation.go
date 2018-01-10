package linsangmysql

import (
   "database/sql"
   "github.com/luisfurquim/linsang"
)

// SetOperation is the low-level, storage dependant, compute method for the most basic operators that computes the check of relations
// between nodes. Applications should avoid using it directly and prefer to call the Walk method of either linsang.Store or linsang.Set.
func (hnd *Handle) SetOperation(srcNode *linsang.Node, subject, predicate, object, label interface{}, chkArcType func(bool) bool, needSubj, needObj bool) (linsang.Set, error) {
   var err           error
   var pred          string
   var preds       []string
   var predInt       int64
   var nodeType      uint8
   var res          *MySQLSet
   var rows         *sql.Rows
   var singleQuery  *sql.Stmt
   var comboQuery   *sql.Stmt
   var qtParm        int
   var WhichVtx      int

   // Do some conversion types e data fetching common to all SetOperations
   res, nodeType, preds, err = hnd.operationPreCheck(srcNode,predicate)
   if err != nil {
      if err == ErrNotFound {
         Goose.Read.Logf(1,"No source node %#v node", srcNode)
         return res, nil
      }
      Goose.Read.Logf(1,"Error searching object %s, %s, %s, %s node: %s", subject, predicate, object, label, err)
      return nil, err
   }

   if needSubj && needObj {
      singleQuery = hnd.GetLinkSbjObjQuery
      comboQuery  = hnd.GetLinkSPOPHashQuery
      qtParm      = 4
      WhichVtx    = linsang.SubObj
   } else if needObj {
      singleQuery = hnd.GetLinkSbjQuery
      comboQuery  = hnd.GetLinkSPHashQuery
      qtParm      = 2
      WhichVtx    = linsang.Subject
   } else if needSubj {
      singleQuery = hnd.GetLinkObjQuery
      comboQuery  = hnd.GetLinkOPHashQuery
      qtParm      = 2
      WhichVtx    = linsang.Object
   } else {
      Goose.Read.Logf(1,"%s: %s", ErrUndefOp, err)
      return nil, ErrUndefOp
   }

   for _, pred = range preds {
      predInt, rows, err = hnd.operationDefineQuery(pred, srcNode, nodeType, singleQuery, comboQuery,qtParm)
      if err != nil {
         if err == ErrNotFound {
            return res, nil
         }
         Goose.Read.Logf(1,"Error searching object %s, %s, %s, %s node: %s", subject, predicate, object, label, err)
         return nil, err
      }

      defer rows.Close()

      Goose.Read.Logf(6,"Equals links: %#v", rows)

      err = hnd.operationExec(rows, pred, predInt, srcNode, nodeType, res, WhichVtx, chkArcType)
      if err != nil {
         Goose.Read.Logf(1,"Error fetching value from Input arc: %s", err)
         return nil, err
      }
   }

   return res, nil
}
