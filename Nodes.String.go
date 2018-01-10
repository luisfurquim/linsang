package linsang

import (
   "fmt"
)

func (n Nodes) String() string {
   var s string
   var v Node

   if len(n) == 0 {
      return ""
   }

   Goose.Logf(6,">>>>>>>>>>>> Nodes.String: %#v",n)

   for _, v = range n {
      s += fmt.Sprintf("%s ",v)
   }

   Goose.Logf(6,"<<<<<<<<<<<< Nodes.String: %#v",n)

   return s[:len(s)-1]
}

