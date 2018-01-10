package linsangmysql


import (
   "database/sql"
   "github.com/go-sql-driver/mysql"
   "github.com/luisfurquim/linsang"
)


/*

Creates an edge connecting 2 vertices. If some or both vertices does not exist, it/they is/are created.
No vertex is created with redundancy.

If the arctype is linsang.Undirected, it's a line and there is no difference in order of the vertices,
e. g.
   Handle.Link("bob","workswith","alice",nil,linsang.Undirected)
is the same as
   Handle.Link("alice","workswith","bob",nil,linsang.Undirected)
and creates this:
   +-----+  workswith  +-------+
   | bob |-------------| alice |
   +-----+             +-------+

If the arctype is linsang.Directed, it's an arc and then the subject vertex is the source of the arc and
the object vertex is the target of the arc, e. g.
   Handle.Link("bob","follows","alice",nil,linsang.Directed)
means that bob follows alice and alice is followed by bob and creates this:
   +-----+  follows  +-------+
   | bob |---------->| alice |
   +-----+           +-------+


On the other hand, the code below
   Handle.Link("alice","follows","bob",nil,linsang.Directed)
means that alice follows bob and bob is followed by alice and creates this:
   +-----+  follows  +-------+
   | bob |<----------| alice |
   +-----+           +-------+


*/
func (hnd *Handle) Link(subject interface{}, predicate string, object interface{}, label interface{}, arctype bool) error {
   var err                       error
   var tx                       *sql.Tx
   var sbjInt, objInt, predInt   int64
   var lblInt                    int64
   var sbjType, objType, lblType uint8
   var myerr                    *mysql.MySQLError
   var res                       sql.Result


   tx, err = hnd.Conn.Begin()
   if err != nil {
      Goose.Write.Logf(1,"Cannot begin transaction: %s",err)
      return err
   }

   sbjInt, sbjType, err = hnd.addNode(subject)
   if err != nil {
      Goose.Write.Logf(1,"Cannot add subject node: %s",err)
      tx.Rollback()
      return err
   }

   predInt, err = hnd.addNodeString(predicate)
   if err != nil {
      Goose.Write.Logf(1,"Cannot add predicate node: %s",err)
      tx.Rollback()
      return err
   }

   objInt, objType, err = hnd.addNode(object)
   if err != nil {
      Goose.Write.Logf(1,"Cannot add object node: %s",err)
      tx.Rollback()
      return err
   }

   lblInt, lblType, _ = hnd.addNode(label) // We deliberately ignore the error, thus making label optional

   if arctype == linsang.Undirected {
      if (sbjInt > objInt) || ((sbjInt == objInt) && (sbjType > objType)) {
         // For undirected graphs there is no order in the (subject, object) pair.
         // So to avoid duplicates, we order them to force a MySQL error caused by the UNIQUE clause of the links table
         sbjInt,  objInt  = objInt,  sbjInt
         sbjType, objType = objType, sbjType
      }
   }

   res, err = hnd.AddLinkQuery.Exec(
      int64((hash64shift(uint64(sbjInt)) ^ hash64shift(uint64(predInt)) ^ hash64shift(uint64(sbjType)) ^ hash64shift(uint64(predInt<<32))) & 0x7fffffffffffffff), // sp_hash
      int64((hash64shift(uint64(sbjInt)) ^ hash64shift(uint64(objInt))  ^ hash64shift(uint64(sbjType)) ^ hash64shift(uint64(objType))  ^ hash64shift(uint64(objInt<<32)))  & 0x7fffffffffffffff), // so_hash
      int64((hash64shift(uint64(objInt)) ^ hash64shift(uint64(predInt)) ^ hash64shift(uint64(objType)) ^ hash64shift(uint64(predInt<<32))) & 0x7fffffffffffffff), // op_hash
      sbjInt,   // subject_id
      predInt,  // predicate_id
      objInt,   // object_id
      lblInt,   // label_id
      sbjType,  // subject_type
      objType,  // object_type
      lblType,  // label_type
      arctype)  // arctype
   if err != nil {
      Goose.Write.Logf(6,"Insert link error: %#v",err)

      switch err.(type) {
         case *mysql.MySQLError:
            myerr = err.(*mysql.MySQLError)
            if myerr.Number != ErrMySQLInsertDuplicate {
               Goose.Write.Logf(1,"Cannot add link: %s",err)
               tx.Rollback()
               return err
            }
         default:
            Goose.Write.Logf(1,"Cannot add link: %s",err)
            tx.Rollback()
            return err
      }

   }

   Goose.Write.Logf(6,"res: %#v\n\n",res)

   err = tx.Commit()
   if err != nil {
      Goose.Init.Logf(1,"Couldn't commit creation transaction: %s", err)
      return err
   }

   Goose.Write.Logf(3,"LINKED: %#v %s %#v",subject,predicate,object)
   return nil

}
