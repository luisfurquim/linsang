package linsangmysql

func (set *MySQLSet) AllBool() ([]bool, error) {
   var res []bool
   var sz    int64
   var err   error
   var value bool

   sz = set.len-(set.curr+set.offset)-1

   if sz <= 0 {
      return nil, nil
   }

   res = []bool{}

   for set.Next() {
      value, err = set.GetBool()
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
