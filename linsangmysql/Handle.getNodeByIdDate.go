package linsangmysql

import (
   "time"
   "database/sql"
   _ "github.com/go-sql-driver/mysql"
)

func (hnd *Handle) getNodeByIdDate(nodeId int64) (time.Time, error) {
   var err   error
   var rows *sql.Rows
   var value time.Time

   Goose.Read.Logf(6,"Checking date node")

   // Fetch the primary key by value
   rows, err = hnd.GetNodeDateIdQuery.Query(nodeId)
   if err != nil {
      return time.Now(), err
   }
   defer rows.Close()

   Goose.Read.Logf(6,"Date node: %#v", rows)

   for rows.Next() {
      rows.Scan(&value)
      Goose.Read.Logf(6,"Found date node")
      return value, nil
   }

   return time.Now(), ErrNotFound
}

