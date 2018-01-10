package linsangmysql

import (
   "time"
   "database/sql"
   _ "github.com/go-sql-driver/mysql"
)

func (hnd *Handle) getNodeDate(value time.Time) (int64, error) {
   var err   error
   var rows *sql.Rows
   var id    uint64

   Goose.Read.Logf(6,"Checking date node")

   // Fetch the primary key by value
   rows, err = hnd.GetNodeDateQuery.Query(value)
   if err != nil {
      return 0, err
   }
   defer rows.Close()

   Goose.Read.Logf(6,"Date node: %#v", rows)

   for rows.Next() {
      rows.Scan(&id)
      Goose.Read.Logf(6,"Found date node")
      return int64(id&0x7fffffffffffffff), nil
   }

   return 0, ErrNotFound
}

