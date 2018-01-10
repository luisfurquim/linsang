package linsang

import (
   "fmt"
   "reflect"
)

func (n Node) String() string {
   var s string
   var v interface{}
   var f string

   Goose.Logf(6,"\\\\\\\\\\\\\\\\\\ Node.String: %#v",n)

   switch reflect.TypeOf(n.Value).Kind() {
      case reflect.Bool:    f = "%t"
      case reflect.Int64:   f = "%d"
      case reflect.Float64: f = "%f"
      default:              f = "%q"
   }

   if n.Via != nil {
      for _, v = range n.Via {
         s += fmt.Sprintf(f+":",v)
      }
      return fmt.Sprintf("%d=" + f + "=>%s",n.Id,n.Value,s[:len(s)-1])
   }

   return fmt.Sprintf("%d="+f,n.Id,n.Value)
}

