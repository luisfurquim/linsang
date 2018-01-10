package linsangmysql

import (
   "reflect"
   "github.com/luisfurquim/linsang"
)

func (op Vertex) Compute(store linsang.Store, srcNode *linsang.Node) (linsang.Set, error) {
   var err           error
   var nodeInt       int64
   var res           MySQLSet
   var rval          reflect.Value
   var hnd          *Handle

   switch store.(type) {
      case *Handle:
         hnd = store.(*Handle)
      default:
         return nil, ErrWrongType
   }

   if isRef(&linsang.Node{Value:op.Value}) {
      rval = reflect.ValueOf(op.Value)
      if reflect.Zero(reflect.TypeOf(op.Value)) == rval {
         return nil, ErrOpNotFound
      }
   }

   nodeInt, _, err = hnd.getNode(op.Value)
   if err != nil {
      Goose.Read.Logf(1,"Error searching object %#v node: %s", op, err)
      return nil, err
   }

   res.curr   = -1
   res.vertex = linsang.Nodes{
      linsang.Node{
         Value: op.Value,
         Id:    nodeInt,
      },
   }

   return &res, nil

}

