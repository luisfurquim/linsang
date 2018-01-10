package linsangmysql

import (
   "database/sql"
   _ "github.com/go-sql-driver/mysql"
)

func (hnd *Handle) getNodeByIdReal(nodeId int64) (float64, error) {
   var err   error
   var rows  *sql.Rows
   var value  float64

   Goose.Read.Logf(6,"Checking real node")

   // Fetch the primary key by value
   rows, err = hnd.GetNodeRealQuery.Query(nodeId)
   if err != nil {
      return 0, err
   }
   defer rows.Close()

   Goose.Read.Logf(6,"Real node: %#v", rows)

   for rows.Next() {
      rows.Scan(&value)
      Goose.Read.Logf(6,"Found real node")
      return value, nil
   }

   return 0.0, ErrNotFound
}

