package linsangmysql

import (
//   "fmt"
   "database/sql"
   _ "github.com/go-sql-driver/mysql"
)

/*

Create the tables and indexes in an already opened MySQL database.
If the tables/indexes already exists it doesn't destroy its contents.

*/
func (hnd *Handle) Create() error {
   var err  error
   var tx  *sql.Tx
   var linkTable, nodeStrTable, nodeDoubleTable, nodeDateTable, linkIndex, nodeIndex sql.Result

   tx, err = hnd.Conn.Begin()
   if err != nil {
      Goose.Init.Logf(1,"Cannot begin transaction: %s", err)
      return err
   }

   linkTable, err = tx.Exec(`
      CREATE TABLE IF NOT EXISTS links (
         sp_hash BIGINT NOT NULL,
         so_hash BIGINT NOT NULL,
         op_hash BIGINT NOT NULL,
         subject_id BIGINT NOT NULL,
         predicate_id BIGINT NOT NULL,
         object_id BIGINT NOT NULL,
         label_id BIGINT,
         subject_type TINYINT UNSIGNED NOT NULL,
         object_type TINYINT UNSIGNED NOT NULL,
         label_type TINYINT UNSIGNED,
         arctype BOOL NOT NULL,
         UNIQUE(arctype, predicate_id, object_id, object_type, subject_id, subject_type, label_id, label_type)
      );
   `)
   if err != nil {
      Goose.Init.Logf(1,"Cannot create link table: %s", err)
      Goose.Init.Logf(6,"Link table not created: %#v", linkTable)
      tx.Rollback()
      return err
   }

   nodeStrTable, err = tx.Exec(`
      CREATE TABLE IF NOT EXISTS nodes_string (
         id SERIAL PRIMARY KEY,
         hash BIGINT NOT NULL,
         value TEXT NOT NULL
      )
   `)
   if err != nil {
      Goose.Init.Logf(1,"Cannot create string table: %s", err)
      Goose.Init.Logf(6,"Node string table not created: %#v", nodeStrTable)
      tx.Rollback()
      return err
   }

   nodeDoubleTable, err = tx.Exec(`
      CREATE TABLE IF NOT EXISTS nodes_double (
         id SERIAL PRIMARY KEY,
         value DOUBLE NOT NULL
      )
   `)
   if err != nil {
      Goose.Init.Logf(1,"Cannot create node double table: %s", err)
      Goose.Init.Logf(6,"Node double table not created: %#v", nodeDoubleTable)
      tx.Rollback()
      return err
   }

   nodeDateTable, err = tx.Exec(`
      CREATE TABLE IF NOT EXISTS nodes_date (
         id SERIAL PRIMARY KEY,
         value DATETIME NOT NULL
      )
   `)
   if err != nil {
      Goose.Init.Logf(1,"Cannot create node date table: %s", err)
      Goose.Init.Logf(6,"Node date table not created: %#v", nodeDateTable)
      tx.Rollback()
      return err
   }

   tx.Exec(`DROP INDEX sp_index ON links;`)
   linkIndex, err = tx.Exec(`CREATE INDEX sp_index ON links (sp_hash,arctype);`)
   if err != nil {
      Goose.Init.Logf(1,"Cannot create link indexes: %s", err)
      Goose.Init.Logf(6,"Link indexes not created: %#v", linkIndex)
      tx.Rollback()
      return err
   }

   tx.Exec(`DROP INDEX so_index ON links;`)
   linkIndex, err = tx.Exec(`CREATE INDEX so_index ON links (so_hash,arctype);`)
   if err != nil {
      Goose.Init.Logf(1,"Cannot create link indexes: %s", err)
      Goose.Init.Logf(6,"Link indexes not created: %#v", linkIndex)
      tx.Rollback()
      return err
   }

   tx.Exec(`DROP INDEX op_index ON links;`)
   linkIndex, err = tx.Exec(`CREATE INDEX op_index ON links (op_hash,arctype);`)
   if err != nil {
      Goose.Init.Logf(1,"Cannot create link indexes: %s", err)
      Goose.Init.Logf(6,"Link indexes not created: %#v", linkIndex)
      tx.Rollback()
      return err
   }

   tx.Exec(`DROP INDEX s_index ON links;`)
   linkIndex, err = tx.Exec(`CREATE INDEX s_index ON links (subject_id,subject_type,arctype);`)
   if err != nil {
      Goose.Init.Logf(1,"Cannot create link indexes: %s", err)
      Goose.Init.Logf(6,"Link indexes not created: %#v", linkIndex)
      tx.Rollback()
      return err
   }

   tx.Exec(`DROP INDEX o_index ON links;`)
   linkIndex, err = tx.Exec(`CREATE INDEX o_index ON links (object_id,object_type,arctype);`)
   if err != nil {
      Goose.Init.Logf(1,"Cannot create link indexes: %s", err)
      Goose.Init.Logf(6,"Link indexes not created: %#v", linkIndex)
      tx.Rollback()
      return err
   }

   tx.Exec(`DROP INDEX p_index ON links;`)
   linkIndex, err = tx.Exec(`CREATE INDEX p_index ON links (predicate_id,arctype);`)
   if err != nil {
      Goose.Init.Logf(1,"Cannot create link indexes: %s", err)
      Goose.Init.Logf(6,"Link indexes not created: %#v", linkIndex)
      tx.Rollback()
      return err
   }

   tx.Exec(`DROP INDEX nsh_index ON nodes_string;`)
   nodeIndex, err = tx.Exec(`CREATE INDEX nsh_index ON nodes_string (hash);`)
   if err != nil {
      Goose.Init.Logf(1,"Cannot create nsh_index: %s", err)
      Goose.Init.Logf(6,"Nsh_index not created: %#v", nodeIndex)
      tx.Rollback()
      return err
   }


   tx.Exec(`DROP INDEX nrv_index ON nodes_double;`)
   nodeIndex, err = tx.Exec(`CREATE INDEX nrv_index ON nodes_double (value);`)
   if err != nil {
      Goose.Init.Logf(1,"Cannot create nrv_index: %s", err)
      Goose.Init.Logf(6,"Nrv_index not created: %#v", nodeIndex)
      tx.Rollback()
      return err
   }


   tx.Exec(`DROP INDEX ndv_index ON nodes_date;`)
   nodeIndex, err = tx.Exec(`CREATE INDEX ndv_index ON nodes_date (value);`)
   if err != nil {
      Goose.Init.Logf(1,"Cannot create ndv_index: %s", err)
      Goose.Init.Logf(6,"Ndv_index not created: %#v", nodeIndex)
      tx.Rollback()
      return err
   }

   err = tx.Commit()
   if err != nil {
      Goose.Init.Logf(1,"Couldn't commit creation transaction: %s", err)
      return err
   }

   return nil
}


