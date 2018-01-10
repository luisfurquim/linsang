package linsangmysql


/*

Returns an array with all the values in the node set.

*/
func (set MySQLSet) All() ([]interface{}, error) {
   var res []interface{}
   var sz    int64
   var err   error
   var value interface{}
   var i     int

   sz = set.len-(set.curr+set.offset)-1

   if sz <= 0 {
      return nil, nil
   }

   res = make([]interface{},sz)

   for set.Next() {
      value, err = set.Get()
      if err != nil {
         return nil, err
      }
      res[i] = value
      i++
   }

   return res, nil
}
