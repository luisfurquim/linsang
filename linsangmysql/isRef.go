package linsangmysql

import (
   "github.com/luisfurquim/linsang"
)

func isRef(n *linsang.Node) bool {
   switch n.Value.(type) {
      case bool, int64:
         return false
   }
   return true
}