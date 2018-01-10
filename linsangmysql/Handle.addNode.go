package linsangmysql

import (
   "time"
   "reflect"
   "github.com/luisfurquim/linsang"
)

func (hnd *Handle) addNode(value interface{}) (int64,  uint8, error) {
   var v    int64
   var err  error

   switch value.(type) {
      case bool:
         if value.(bool) {
            return 1, linsang.BoolNode, nil
         }
         return 0, linsang.BoolNode, nil
      case uint8:
         return int64(value.(uint8)), linsang.IntNode, nil
      case uint16:
         return int64(value.(uint16)), linsang.IntNode, nil
      case uint32:
         return int64(value.(uint32)), linsang.IntNode, nil
      case int8:
         return int64(value.(int8)), linsang.IntNode, nil
      case int16:
         return int64(value.(int16)), linsang.IntNode, nil
      case int32:
         return int64(value.(int32)), linsang.IntNode, nil
      case int64:
         return value.(int64), linsang.IntNode, nil
      case uint:
         return int64(value.(uint)), linsang.IntNode, nil
      case int:
         if reflect.TypeOf(value.(int)).Size() == 8 {
            return 0, 0, ErrInvalidType
         }
         return int64(value.(int)), linsang.IntNode, nil
      case float32:
         v, err = hnd.addNodeReal(float64(value.(float32)))
         if err != nil {
            return 0, 0, err
         }
         return v, linsang.RealNode, nil
      case float64:
         v, err = hnd.addNodeReal(value.(float64))
         if err != nil {
            return 0, 0, err
         }
         return v, linsang.RealNode, nil
      case []byte:
         v, err = hnd.addNodeString(string(value.([]byte)))
         if err != nil {
            return 0, 0, err
         }
         return v, linsang.StringNode, nil
      case string:
         v, err = hnd.addNodeString(value.(string))
         if err != nil {
            return 0, 0, err
         }
         return v, linsang.StringNode, nil
      case time.Time:
         v, err = hnd.addNodeDate(value.(time.Time))
         if err != nil {
            return 0, 0, err
         }
         return v, linsang.DateNode, nil
   }

   return 0, 0, ErrInvalidType
}
