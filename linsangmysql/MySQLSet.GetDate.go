package linsangmysql

import (
   "time"
)

/*

Returns the current value of the node set if the value is of type time.Time, otherwise returns the ErrWrongType error.
Returns the ErrEOS error if all the nodes of the set are already received (similar to the os.EOF error).

*/
func (set *MySQLSet) GetDate() (time.Time, error) {
   var t time.Time

   if (set.curr < int64(len(set.vertex))) && (set.curr >= 0) {
      switch set.vertex[set.curr].Value.(type) {
         case time.Time:
            return set.vertex[set.curr].Value.(time.Time), nil
         default:
            return t, ErrWrongType
      }
   }

   return t, ErrEOS
}
