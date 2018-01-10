package linsangmysql

import (
//   "os"
//   "fmt"
//   "reflect"
   "database/sql"
   _ "github.com/go-sql-driver/mysql"
)

func (hnd *Handle) getNodeByIdString(nodeId int64) (string, error) {
   var err    error
   var rows  *sql.Rows
   var value  string

   // fetch the string value by primary key and check if it is the requested value or it is a hash collision
   rows, err = hnd.GetNodeStringIdQuery.Query(nodeId)
   if err != nil {
      Goose.Read.Logf(1,"Error fetching string node: %s\n",err)
      return "", err
   }
   defer rows.Close()

   for rows.Next() {
      rows.Scan(&value)
      Goose.Read.Logf(6,"Found string node by Id: %s",value)
      return value, nil
   }

   return "", ErrNotFound
}

