package linsangmysql

import (
   "fmt"
)

/*

Just used for fmt.Printf("%s",set) purposes.

*/
func (set *MySQLSet) String() string {
   var s string

   Goose.Read.Logf(6,">>>>>>>>>>>>> SET: %#v",set.vertex)
   if (len(set.vertex)==0) && (set.rows!=nil) && ((set.curr+set.offset) < set.len) {
      Goose.Read.Logf(6,"SET: curr=%d, |vertex|=%d",set.curr,len(set.vertex))
      set.fillBuffer()
      Goose.Read.Logf(6,"SET: curr=%d, |vertex|=%d",set.curr,len(set.vertex))
   }

   s = fmt.Sprintf("%s",set.vertex)

   Goose.Read.Logf(6,"<<<<<<<<<<<<< SET: %#v",set.vertex)

   return s
}

