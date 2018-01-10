package linsangmysql

/*

Returns the current value of the node set if the value is of type boole, otherwise returns the ErrWrongType error.
Returns the ErrEOS error if all the nodes of the set are already received (similar to the os.EOF error).

*/
func (set *MySQLSet) GetBool() (bool, error) {

   if (set.curr < int64(len(set.vertex))) && (set.curr >= 0) {
      switch set.vertex[set.curr].Value.(type) {
         case bool:
            return set.vertex[set.curr].Value.(bool), nil
         default:
            return false, ErrWrongType
      }
   }



   return false, ErrEOS
}

