package linsangmysql

func (set *MySQLSet) AllString() ([]string, error) {
   var res []string
   var sz    int64
   var err   error
   var value string

   sz = set.len-(set.curr+set.offset)-1

   if sz <= 0 {
      return nil, nil
   }

   res = []string{}

   for set.Next() {
      value, err = set.GetString()
      if err == ErrWrongType {
         continue
      }

      if err != nil {
         return nil, err
      }

      res = append(res,value)
   }

   return res, nil
}
