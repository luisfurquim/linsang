package linsangmysql

import (
//   "os"
//   "fmt"
//   "reflect"
   "database/sql"
   _ "github.com/go-sql-driver/mysql"
)

func (hnd *Handle) getNodeString(value string) (int64, error) {
   var err    error
   var hrows  *sql.Rows
   var vrows  *sql.Rows
   var id     uint64
   var h      int64
   var v      string

   h = hash64shiftString(value)

   // Fetch the primary key by hash(value)
   hrows, err = hnd.GetNodeStringHashQuery.Query(h)
   if err != nil {
      Goose.Read.Logf(1,"Error fetching string node: %s\n",err)
      return 0, err
   }
   defer hrows.Close()

   for hrows.Next() {
      hrows.Scan(&id)

      Goose.Read.Logf(6,"Id: %#v\n",id)

      // fetch the string value by primary key and check if it is the requested value or it is a hash collision
      vrows, err = hnd.GetNodeStringIdQuery.Query(id)
      if err != nil {
         Goose.Read.Logf(1,"Error fetching string node: %s\n",err)
         return 0, err
      }

      for vrows.Next() {
         vrows.Scan(&v)
         if v == value { // Confirmed the value already exists, return its primary key
            vrows.Close()
            Goose.Read.Logf(6,"Found string node by Id: %s",v)
            return int64(id&0x7fffffffffffffff), nil
         }
      }

      vrows.Close()
   }

   return 0, ErrNotFound
}

