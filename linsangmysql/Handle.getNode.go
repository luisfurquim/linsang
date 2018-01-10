package linsangmysql

import (
   "time"
   "reflect"
   "github.com/luisfurquim/linsang"
)

func (hnd *Handle) getNode(value interface{}) (int64,  uint8, error) {
   var v    int64
   var err  error

   switch value.(type) {
      case bool:
         v, err = hnd.getNodeBool(value.(bool))
         return v, linsang.BoolNode, err
      case uint8:
         v, err = hnd.getNodeInt(int64(value.(uint8)))
         return v, linsang.IntNode, err
      case uint16:
         v, err = hnd.getNodeInt(int64(value.(uint16)))
         return v, linsang.IntNode, err
      case uint32:
         v, err = hnd.getNodeInt(int64(value.(uint32)))
         return v, linsang.IntNode, err
      case int8:
         v, err = hnd.getNodeInt(int64(value.(int8)))
         return v, linsang.IntNode, err
      case int16:
         v, err = hnd.getNodeInt(int64(value.(int16)))
         return v, linsang.IntNode, err
      case int32:
         v, err = hnd.getNodeInt(int64(value.(int32)))
         return v, linsang.IntNode, err
      case int64:
         v, err = hnd.getNodeInt(value.(int64))
         return v, linsang.IntNode, err
      case uint:
         if reflect.TypeOf(value.(uint)).Size() == 8 {
            return 0, 0, ErrInvalidType
         }
         v, err = hnd.getNodeInt(int64(value.(uint)))
         return v, linsang.IntNode, err
      case int:
         v, err = hnd.getNodeInt(int64(value.(int)))
         return v, linsang.IntNode, err
      case float32:
         v, err = hnd.getNodeReal(float64(value.(float32)))
         return v, linsang.RealNode, err
      case float64:
         v, err = hnd.getNodeReal(value.(float64))
         return v, linsang.RealNode, err
      case []byte:
         v, err = hnd.getNodeString(string(value.([]byte)))
         return v, linsang.StringNode, err
      case string:
         v, err = hnd.getNodeString(value.(string))
         return v, linsang.StringNode, err
      case time.Time:
         v, err = hnd.getNodeDate(value.(time.Time))
         return v, linsang.DateNode, err
   }

   return 0, 0, ErrInvalidType
}
