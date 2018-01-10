package linsangmysql

import (
   "fmt"
   "time"
   "github.com/luisfurquim/linsang"
)

/*

Fills the set buffer with the next nodes.
Returns how many items where loaded into the buffer.

*/
func (set *MySQLSet) fillBuffer() int {
   var i    int
   var typ  int
   var sval string
   var ival int64
   var rval float64

   if set.rows != nil {
      Goose.Read.Logf(6,"MySQLSet.Next  D")
      set.offset += set.curr
      set.curr    = -1
      if (set.len-set.offset) > BufSize {
         Goose.Read.Logf(6,"MySQLSet.Next  E")
         set.vertex  = make(linsang.Nodes,BufSize)
      } else {
         Goose.Read.Logf(6,"MySQLSet.Next  F sz=%d",set.len-set.offset)
         set.vertex  = make(linsang.Nodes,set.len-set.offset)
      }
      for i=0; i<len(set.vertex); i++ {
         Goose.Read.Logf(6,"MySQLSet.Next  G")
         if !set.rows.Next() {
            Goose.Read.Logf(6,"MySQLSet.Next  H sz=%d",i)
            set.vertex = set.vertex[:i]
            break
         }
         set.rows.Scan(&set.vertex[i].Id,&set.vertex[i].Value,&typ)
         sval = string(set.vertex[i].Value.([]byte))
         switch typ {
            case linsang.BoolNode:
               set.vertex[i].Value = sval != "0"
            case linsang.IntNode:
               fmt.Sscanf(sval,"%d",&ival)
               set.vertex[i].Value = ival
            case linsang.RealNode:
               fmt.Sscanf(sval,"%f",&rval)
               set.vertex[i].Value = rval
            case linsang.StringNode:
               set.vertex[i].Value = sval
            case linsang.DateNode:
               Goose.Read.Logf(1,"===> %d=%#v",typ,sval)
               set.vertex[i].Value, _ = time.Parse(DateForm, sval)
               Goose.Read.Logf(1,"===> %d=%#v",typ,set.vertex[i].Value)
         }
      }
      Goose.Read.Logf(1,"===> Done reading buf %#v",set)
   }

   return i
}

