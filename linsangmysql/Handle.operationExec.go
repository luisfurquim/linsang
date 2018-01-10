package linsangmysql

import (
   "database/sql"
   "github.com/luisfurquim/linsang"
)

func (hnd *Handle) operationExec(rows *sql.Rows, pred string, predInt int64, srcNode *linsang.Node, nodeType uint8, res *MySQLSet, vtx int, chkArcType func(bool) bool) error {
   var err           error
   var id         [4]int64
   var typ        [4]uint8
   var relVtx        int
   var arctype       bool
   var node          linsang.Node
   var bothVtx       bool

   switch vtx {
      case linsang.Subject:
         relVtx = linsang.Object
      case linsang.Object:
         relVtx = linsang.Subject
      case linsang.SubObj:
         vtx    = linsang.Subject
         relVtx = linsang.Object
         bothVtx = true
   }

   for rows.Next() {
      rows.Scan(&id[0],&id[1],&id[2],&id[3],&typ[0],&typ[2],&typ[3],&arctype)
      Goose.Read.Logf(6,"sbj_id:%d,pred_id:%d,obj_id:%d,lbl_id:%d,sbj_type:%d,obj_type:%d,lbl_type:%d,arctype:%t",id[0],id[1],id[2],id[3],typ[0],typ[2],typ[3],arctype)

      if chkArcType(arctype) {
         continue
      }

      if srcNode != nil { // The cases where we have a left operand...
         if pred!="" && predInt != id[1] {
            Goose.Read.Logf(7,"Predicate hash collision")
            continue
         }

         if bothVtx {
            if (id[relVtx]==srcNode.Id) && (typ[relVtx]==nodeType) {
               node.Value, err = hnd.getNodeById(id[vtx],typ[vtx])
               node.Id         = id[vtx]
            } else if (id[vtx]==srcNode.Id) && (typ[vtx]==nodeType) {
               node.Value, err = hnd.getNodeById(id[relVtx],typ[relVtx])
               node.Id         = id[relVtx]
            } else {
               // Just a hash collision
               continue
            }
         } else {
            if (id[vtx]!=srcNode.Id) || (typ[vtx]!=nodeType) {
               continue
            }
            node.Value, err = hnd.getNodeById(id[relVtx],typ[relVtx])
            node.Id         = id[relVtx]
         }

      } else {
         // The cases whith no left operand
         if bothVtx {
            // We have to get either vertices of the arc
            node.Value, err = hnd.getNodeById(id[vtx],typ[vtx])
            node.Id         = id[vtx]
            if err != nil {
               Goose.Read.Logf(1,"Error fetching subject value from arc: %s", err)
               return err
            }
            res.vertex.Add(node)
            res.len++

            node.Value, err = hnd.getNodeById(id[relVtx],typ[relVtx])
            node.Id         = id[relVtx]

         } else {
            node.Value, err = hnd.getNodeById(id[relVtx],typ[relVtx])
            node.Id         = id[relVtx]
         }
      }

      if err != nil {
         Goose.Read.Logf(1,"Error fetching value from Input arc: %s", err)
         return err
      }

      res.vertex.Add(node)
      res.len++

      Goose.Read.Logf(6,"Found node %s",node)
   }

   return nil

}
