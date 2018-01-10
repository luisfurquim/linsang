package linsangmysql

import (
   "database/sql"
   _ "github.com/go-sql-driver/mysql"
   "github.com/luisfurquim/linsang"
)

func (hnd *Handle) getNodeBool(value bool) (int64, error) {
   var err    error
   var rows  *sql.Rows
   var chk    uint64
   var v      int64

   Goose.Read.Logf(6,"Checking bool node")
   if value {
      v = 1
   }

   // Fetch the primary key by value
   rows, err = hnd.ExistSubjectOrObjectQuery.Query(v,linsang.BoolNode,v,linsang.BoolNode)
   if err != nil {
      return 0, err
   }
   defer rows.Close()

   Goose.Read.Logf(6,"Bool node: %#v", rows)

   for rows.Next() {
      rows.Scan(&chk)
      Goose.Read.Logf(6,"Found bool node")
      if chk==1 {
         return v, nil
      }
   }

   return 0, ErrNotFound
}

