package linsangmysql

/*

Closes the SQL query that generated this node set.

*/
func (set *MySQLSet) Close() error {
   var err error

   set.len = 0

   if set.vertex != nil {
      set.vertex = nil
   }

   if set.rows != nil {
      err      = set.rows.Close()
      set.rows = nil
   }

   return err
}
