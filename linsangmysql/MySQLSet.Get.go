package linsangmysql

/*

Returns the current value of the node set.
Returns the ErrEOS error if all the nodes of the set are already received (similar to the os.EOF error).

*/
func (set *MySQLSet) Get() (interface{}, error) {

   if (set.curr < int64(len(set.vertex))) && (set.curr >= 0) {
      return set.vertex[set.curr].Value, nil
   }

   return nil, ErrEOS
}

