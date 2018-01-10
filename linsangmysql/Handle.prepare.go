package linsangmysql

import (
   "fmt"
   "database/sql"
   _ "github.com/go-sql-driver/mysql"
   "github.com/luisfurquim/linsang"
)

/*

Execute internal initialization of SQL queries and MUST be executed before accessing the graphs stored in the database
AND after the database creation (which is done by calling the Handle.Create() method).

*/
func (hnd  *Handle) prepare() error {
   var err   error
   var qry   string
   var stmt *sql.Stmt

/*
   qry = fmt.Sprintf(
      "SELECT DISTINCT * FROM (" +
         "SELECT DISTINCT subject_id id, subject_id value, %d typ  FROM links2 WHERE subject_type = %d " +
            "UNION ALL " +
         "SELECT DISTINCT subject_id id, subject_id value, %d typ  FROM links2 WHERE subject_type = %d " +
            "UNION ALL " +
         "SELECT DISTINCT object_id id, object_id value, %d typ  FROM links2 WHERE object_type = %d " +
            "UNION ALL " +
         "SELECT DISTINCT object_id id, object_id value, %d typ  FROM links2 WHERE object_type = %d " +
            "UNION ALL " +
         "SELECT id, value, 2 typ from nodes_string " +
            "UNION ALL " +
         "SELECT id, value, 1 typ from nodes_double " +
            "UNION ALL " +
         "SELECT id, value, 4 typ from nodes_date " +
      ") AS allnodes",
      linsang.BoolNode,linsang.BoolNode,linsang.IntNode,linsang.IntNode,linsang.BoolNode,linsang.BoolNode,linsang.IntNode,linsang.IntNode)
*/

   qry = fmt.Sprintf(
      "SELECT DISTINCT * FROM (" +
         "SELECT DISTINCT subject_id id, subject_id value, %d typ  FROM links WHERE subject_type = %d " +
            "UNION ALL " +
         "SELECT DISTINCT subject_id id, subject_id value, %d typ  FROM links WHERE subject_type = %d " +
            "UNION ALL " +
         "SELECT DISTINCT object_id id, object_id value, %d typ  FROM links WHERE object_type = %d " +
            "UNION ALL " +
         "SELECT DISTINCT object_id id, object_id value, %d typ  FROM links WHERE object_type = %d " +
            "UNION ALL " +
         "SELECT id, value, 2 typ from nodes_string " +
            "UNION ALL " +
         "SELECT id, value, 1 typ from nodes_double " +
            "UNION ALL " +
         "SELECT id, value, 4 typ from nodes_date " +
      ") AS allnodes",
      linsang.BoolNode,linsang.BoolNode,linsang.IntNode,linsang.IntNode,linsang.BoolNode,linsang.BoolNode,linsang.IntNode,linsang.IntNode)

   stmt, err = hnd.Conn.Prepare(qry)
   if err != nil {
      return err
   }

   hnd.GetNodeAllQuery = stmt


   stmt, err = hnd.Conn.Prepare("SELECT COUNT(typ) c FROM (" + qry + ") as cnt")
   if err != nil {
      return err
   }

   hnd.GetNodeCountAllQuery = stmt


   stmt, err = hnd.Conn.Prepare(
      "SELECT subject_id,predicate_id,object_id,label_id,subject_type,object_type,label_type,arctype FROM links WHERE sp_hash = ?")
   if err != nil {
      return err
   }

   hnd.GetLinkSPHashQuery = stmt


   stmt, err = hnd.Conn.Prepare(
      "SELECT subject_id,predicate_id,object_id,label_id,subject_type,object_type,label_type,arctype FROM links WHERE op_hash = ?")
   if err != nil {
      return err
   }

   hnd.GetLinkOPHashQuery = stmt


   stmt, err = hnd.Conn.Prepare(
      "SELECT subject_id,predicate_id,object_id,label_id,subject_type,object_type,label_type,arctype FROM links WHERE so_hash = ?")
   if err != nil {
      return err
   }

   hnd.GetLinkSOHashQuery = stmt


   stmt, err = hnd.Conn.Prepare(
      "SELECT subject_id,predicate_id,object_id,label_id,subject_type,object_type,label_type,arctype FROM links WHERE sp_hash=? or op_hash=?")
   if err != nil {
      return err
   }

   hnd.GetLinkSPOPHashQuery = stmt


   stmt, err = hnd.Conn.Prepare(
      "SELECT subject_id,predicate_id,object_id,label_id,subject_type,object_type,label_type,arctype FROM links WHERE subject_id=? and subject_type=?")
   if err != nil {
      return err
   }

   hnd.GetLinkSbjQuery = stmt


   stmt, err = hnd.Conn.Prepare(
      "SELECT subject_id,predicate_id,object_id,label_id,subject_type,object_type,label_type,arctype FROM links WHERE predicate_id=?")
   if err != nil {
      return err
   }

   hnd.GetLinkPredQuery = stmt


   stmt, err = hnd.Conn.Prepare(
      "SELECT subject_id,predicate_id,object_id,label_id,subject_type,object_type,label_type,arctype FROM links WHERE object_id=? and object_type=?")
   if err != nil {
      return err
   }

   hnd.GetLinkObjQuery = stmt


   stmt, err = hnd.Conn.Prepare(
      "SELECT subject_id,predicate_id,object_id,label_id,subject_type,object_type,label_type,arctype FROM links WHERE subject_id=? and subject_type=? " +
      "UNION ALL " +
      "SELECT subject_id,predicate_id,object_id,label_id,subject_type,object_type,label_type,arctype FROM links WHERE object_id=? and object_type=?")
   if err != nil {
      return err
   }

   hnd.GetLinkSbjObjQuery = stmt


   stmt, err = hnd.Conn.Prepare(
      "SELECT subject_id,predicate_id,object_id,subject_type,object_type,label_type,arctype FROM links WHERE label_id = ?")
   if err != nil {
      return err
   }

   hnd.GetLinkLblQuery = stmt


   stmt, err = hnd.Conn.Prepare(
      "INSERT INTO links (" +
      "sp_hash,so_hash,op_hash,subject_id,predicate_id,object_id,label_id,subject_type,object_type,label_type,arctype" +
      ") values (?,?,?,?,?,?,?,?,?,?,?)")
   if err != nil {
      return err
   }

   hnd.AddLinkQuery = stmt


   stmt, err = hnd.Conn.Prepare("SELECT 1 FROM links WHERE subject_id=? and subject_type=? limit 0,1")
   if err != nil {
      return err
   }

   hnd.ExistSubjectQuery = stmt


   stmt, err = hnd.Conn.Prepare("SELECT 1 FROM links WHERE object_id=? and object_type=? limit 0,1")
   if err != nil {
      return err
   }

   hnd.ExistObjectQuery = stmt


   stmt, err = hnd.Conn.Prepare(
      "SELECT " +
         "EXISTS (SELECT true FROM links WHERE (subject_id=? AND subject_type=?) LIMIT 1) OR " +
         "EXISTS (SELECT true FROM links WHERE (object_id=?  AND object_type=?)  LIMIT 1)")
   if err != nil {
      return err
   }

   hnd.ExistSubjectOrObjectQuery = stmt



   stmt, err = hnd.Conn.Prepare("SELECT id FROM nodes_string WHERE hash=?")
   if err != nil {
      return err
   }

   hnd.GetNodeStringHashQuery = stmt

   stmt, err = hnd.Conn.Prepare("SELECT value FROM nodes_string WHERE id=?")
   if err != nil {
      return err
   }

   hnd.GetNodeStringIdQuery = stmt

   stmt, err = hnd.Conn.Prepare("INSERT INTO nodes_string (hash,value) values (?,?)")
   if err != nil {
      return err
   }

   hnd.AddNodeStringQuery = stmt



   stmt, err = hnd.Conn.Prepare("SELECT id FROM nodes_double WHERE value=?")
   if err != nil {
      return err
   }

   hnd.GetNodeRealQuery = stmt

   stmt, err = hnd.Conn.Prepare("SELECT value FROM nodes_double WHERE id=?")
   if err != nil {
      return err
   }

   hnd.GetNodeRealIdQuery = stmt

   stmt, err = hnd.Conn.Prepare("INSERT INTO nodes_double (value) values (?)")
   if err != nil {
      return err
   }

   hnd.AddNodeRealQuery = stmt



   stmt, err = hnd.Conn.Prepare("SELECT id FROM nodes_date WHERE value=?")
   if err != nil {
      return err
   }

   hnd.GetNodeDateQuery = stmt


   stmt, err = hnd.Conn.Prepare("SELECT value FROM nodes_date WHERE id=?")
   if err != nil {
      return err
   }

   hnd.GetNodeDateIdQuery = stmt

   stmt, err = hnd.Conn.Prepare("INSERT INTO nodes_date (value) values (?)")
   if err != nil {
      return err
   }

   hnd.AddNodeDateQuery = stmt


   return nil
}


