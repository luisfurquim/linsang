package linsangmysql

import (
   "github.com/luisfurquim/linsang"
)

func (hnd *Handle) getNodeById(nodeId int64, nodeType uint8) (interface{}, error) {
   var str string
   var err error

   Goose.Read.Logf(3,"Getting node by id: %d/%d", nodeId, nodeType)

   switch nodeType {
      case linsang.BoolNode:
         return nodeId!=0, nil
      case linsang.IntNode:
         return nodeId, nil
      case linsang.RealNode:
         return hnd.getNodeByIdReal(nodeId)
      case linsang.StringNode:
         str, err = hnd.getNodeByIdString(nodeId)
         Goose.Read.Logf(3,"Got string: %s", str)
         return str, err
      case linsang.DateNode:
         return hnd.getNodeByIdDate(nodeId)
   }

   return nil, ErrInvalidType
}
