package linsangmysql

import (
   "database/sql"
   "github.com/luisfurquim/linsang"
)

func (hnd Handle) SetOperationVtx(values ...interface{}) (linsang.Set, error) {
   var err           error
   var nodeInt       int64
   var res           MySQLSet
   var value         interface{}
   var row          *sql.Row

   res.curr   = -1
   res.vertex = linsang.Nodes{}

   // If caller is asking for all nodes we cannot assume that it can fit entirely
   // in memory. So, we just put the query result in res.rows
   if len(values) == 0 {
      res.rows, err = hnd.GetNodeAllQuery.Query()
      if err != nil {
         Goose.Read.Logf(1,"Error searching all nodes: %s", err)
         return nil, err
      }
      row = hnd.GetNodeCountAllQuery.QueryRow()
      row.Scan(&res.len)
      return &res, nil
   }

   for _, value = range values {
      nodeInt, _, err = hnd.getNode(value)
      if err != nil {
         if err == ErrNotFound {
            return &res, nil
         }
         Goose.Read.Logf(1,"Error searching object %#v node: %s", value)
         return nil, err
      }

      res.len++
      res.vertex.Add(linsang.Node{
         Value: value,
         Id:    nodeInt,
      })
   }

   return &res, nil
}
