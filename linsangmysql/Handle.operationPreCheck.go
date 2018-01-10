package linsangmysql

import (
   "github.com/luisfurquim/linsang"
)

func (hnd *Handle) operationPreCheck(srcNode *linsang.Node, predicate interface{}) (*MySQLSet, uint8, []string, error) {
   var err           error
   var nodeType      uint8
   var preds       []string
   var res           MySQLSet

   switch predicate.(type) {
      case nil:
         preds = []string{""}
      case string:
         preds = []string{predicate.(string)}
      case []string:
         preds = predicate.([]string)
         if len(preds) == 0 {
            preds = []string{""}
         }
      default:
         return nil, 0, nil, ErrInvalidPred
   }

   res.curr   = -1
   res.vertex = make(linsang.Nodes,0,16)

   if srcNode != nil { // We have a left operand
      srcNode.Id, nodeType, err = hnd.getNode(srcNode.Value)
      // The check below could be ignored if we just resturned the err in the last func statemente
      // But it could open a future bug if, in the future, some developer decides to add some code
      // inside this block
      if err != nil {
         return &res, nodeType, preds, err
      }
   }

   return &res, nodeType, preds, nil
}
