package linsangmysql

/*

Moves the node pointer to the next node on the set.
Returns true if the next node exists.
Returns false if all the nodes are already received.

*/
func (set *MySQLSet) Next() bool {
   var i    int

   set.curr++

   Goose.Read.Logf(6,"MySQLSet.Next  A")

   if set.curr < int64(len(set.vertex)) {
      Goose.Read.Logf(6,"MySQLSet.Next  B")
      return true
   }

   Goose.Read.Logf(6,"MySQLSet.Next  C curr=%d, |vertex|=%d",set.curr, int64(len(set.vertex)))

   if (set.curr+set.offset) < set.len {
      i        = set.fillBuffer()
      set.curr = 0
      if i > 0 {
         return true
      }
   }

   return false
}

