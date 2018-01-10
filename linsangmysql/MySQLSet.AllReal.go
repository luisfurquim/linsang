package linsangmysql

func (set *MySQLSet) AllReal() ([]float64, error) {
   var res []float64
   var sz    int64
   var err   error
   var value float64

   sz = set.len-(set.curr+set.offset)-1

   if sz <= 0 {
      return nil, nil
   }

   res = []float64{}

   for set.Next() {
      value, err = set.GetReal()
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
