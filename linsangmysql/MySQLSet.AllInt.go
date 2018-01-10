package linsangmysql

func (set *MySQLSet) AllInt() ([]int64, error) {
   var res []int64
   var sz    int64
   var err   error
   var value int64

   sz = set.len-(set.curr+set.offset)-1

   if sz <= 0 {
      return nil, nil
   }

   res = []int64{}

   for set.Next() {
      value, err = set.GetInt()
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
