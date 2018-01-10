package linsangmysql

import (
   "database/sql"
   _ "github.com/go-sql-driver/mysql"
   "github.com/luisfurquim/linsang"
)

func (hnd *Handle) getNodeInt(value int64) (int64, error) {
   var err    error
   var rows  *sql.Rows
   var chk    uint64

   Goose.Read.Logf(6,"Checking int node")

   // Fetch the primary key by value
   rows, err = hnd.ExistSubjectOrObjectQuery.Query(value,linsang.IntNode,value,linsang.IntNode)
   if err != nil {
      return 0, err
   }
   defer rows.Close()

   Goose.Read.Logf(6,"Int node: %#v", rows)

   for rows.Next() {
      rows.Scan(&chk)
      Goose.Read.Logf(6,"Found real node")
      if chk==1 {
         return value, nil
      }
   }

   return 0, ErrNotFound
}

