package linsangmysql


/*

Returns the current value of the node set if the value is of type float64, otherwise returns the ErrWrongType error.
Returns the ErrEOS error if all the nodes of the set are already received (similar to the os.EOF error).

*/
func (set *MySQLSet) GetReal() (float64, error) {

   if (set.curr < int64(len(set.vertex))) && (set.curr >= 0) {
      switch set.vertex[set.curr].Value.(type) {
         case float64:
            return set.vertex[set.curr].Value.(float64), nil
         default:
            return 0.0, ErrWrongType
      }
   }

   return 0.0, ErrEOS
}
