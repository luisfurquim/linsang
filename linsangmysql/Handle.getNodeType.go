package linsangmysql

import (
   "time"
   "github.com/luisfurquim/linsang"
)

func (hnd *Handle) getNodeType(node interface{}) (uint8, error) {
   switch node.(type) {
      case bool:
         return linsang.BoolNode, nil
      case uint8, uint16, uint32, uint64, int8, int16, int32, int64, uint, int:
         return linsang.IntNode, nil
      case float32, float64:
         return linsang.RealNode, nil
      case []byte, string:
         return linsang.StringNode, nil
      case time.Time:
         return linsang.DateNode, nil
   }

   return 0, ErrInvalidType
}
