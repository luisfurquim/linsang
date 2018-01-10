package linsangmysql

import (
   "database/sql"
   _ "github.com/go-sql-driver/mysql"
)

/*

Returns a handle that satisfy the linsang.Store interface connected to a MySQL database.

*/
func NewFromDSN(dsn string) (*Handle, error) {
   var hnd  *Handle
   var conn *sql.DB
   var err   error

   conn, err = sql.Open("mysql",dsn)
   if err != nil {
      Goose.Init.Logf(1,"Error opening database: %s",err)
      return nil, err
   }

   hnd = &Handle{
      dsn: dsn,
      Conn: conn,
   }

//   hnd.Conn.Register("set names utf8")

   return hnd, nil
}


