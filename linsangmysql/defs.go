/*

This package Implements the interfaces needed by the linsang package in order to store/access the graph database in a MySQL server.

*/
package linsangmysql

import (
   "database/sql"
   _ "github.com/go-sql-driver/mysql"
   "github.com/luisfurquim/linsang"
)

const ErrMySQLInsertDuplicate uint16 = 0x426

/*

This type satisfies the linsang.Store interface.
It also has the connection to the MySQL database server and many SQL queries already compiled via Prepare statement.

*/
type Handle struct {
   dsn                         string
   raddr, user, passwd         string
   db                        []string
   Conn                       *sql.DB
   AddNodeStringQuery         *sql.Stmt
   AddNodeRealQuery           *sql.Stmt
   AddNodeDateQuery           *sql.Stmt
   AddLinkQuery               *sql.Stmt
   GetNodeAllQuery            *sql.Stmt
   GetNodeCountAllQuery       *sql.Stmt
   GetNodeStringHashQuery     *sql.Stmt
   GetNodeStringIdQuery       *sql.Stmt
   GetNodeRealQuery           *sql.Stmt
   GetNodeRealIdQuery         *sql.Stmt
   GetNodeDateQuery           *sql.Stmt
   GetNodeDateIdQuery         *sql.Stmt
   GetLinkSPHashQuery         *sql.Stmt
   GetLinkOPHashQuery         *sql.Stmt
   GetLinkSOHashQuery         *sql.Stmt
   GetLinkSPOPHashQuery       *sql.Stmt
   GetLinkSbjQuery            *sql.Stmt
   GetLinkObjQuery            *sql.Stmt
   GetLinkPredQuery           *sql.Stmt
   GetLinkSbjObjQuery         *sql.Stmt
   GetLinkLblQuery            *sql.Stmt
   ExistSubjectQuery          *sql.Stmt
   ExistObjectQuery           *sql.Stmt
   ExistSubjectOrObjectQuery  *sql.Stmt
}

type MySQLSet struct {
   curr     int64
   offset   int64
   len      int64
   vertex   linsang.Nodes
   rows    *sql.Rows
}

/*
Value, one of:

* a bool, any integer, any float, non empty string a non empty []byte or time.Time: The vertex value to search

* a list of values of the types above: The vertices to search

If computed via linsang.Store.Walk object it searches on all nodes.

If computed via linsang.Set.Walk object it searches only in the nodes belonging to the set.

*/
type Vertex struct {
   Value     interface{}
}

type V Vertex

const DateForm = "Jan 2, 2006 at 3:04pm (MST)"


