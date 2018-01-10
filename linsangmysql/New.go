package linsangmysql

import (
   "fmt"
   "database/sql"
   _ "github.com/go-sql-driver/mysql"
)

/*

Returns a handle that satisfy the linsang.Store interface connected to a MySQL database.

*/
func New(raddr, user, passwd string, db ...string) (*Handle, error) {
   var hnd  *Handle
   var conn *sql.DB
   var err   error

   conn, err = sql.Open("mysql",fmt.Sprintf("%s:%s@tcp(%s)/%s",user, passwd, raddr, db[0]))
   if err != nil {
      Goose.Init.Logf(1,"Error opening database: %s",err)
      return nil, err
   }

   hnd = &Handle{
      raddr:      raddr,
      user:       user,
      passwd:     passwd,
      db:         db,
      Conn:       conn,
   }

//   hnd.Conn.Register("set names utf8")

   return hnd, nil
}


