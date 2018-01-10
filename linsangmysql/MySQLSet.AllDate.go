package linsangmysql

import (
   "time"
)

func (set *MySQLSet) AllDate() ([]time.Time, error) {
   var res []time.Time
   var sz    int64
   var err   error
   var value time.Time

   sz = set.len-(set.curr+set.offset)-1

   if sz <= 0 {
      return nil, nil
   }

   res = []time.Time{}

   for set.Next() {
      value, err = set.GetDate()
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
