package linsang

import (
   "sort"
   "time"
   "reflect"
)

func (n *Nodes) Add(e Node) {
   var pos int

   pos = sort.Search(len(*n), func(i int) bool {
      if reflect.TypeOf(e.Value) == reflect.TypeOf((*n)[i].Value) {
         switch e.Value.(type) {
            case int64:     return (*n)[i].Value.(int64)     >= e.Value.(int64)
            case bool:      return e.Value.(bool)
            case float64:   return (*n)[i].Value.(float64)   >= e.Value.(float64)
            case time.Time: return e.Value.(time.Time).Before((*n)[i].Value.(time.Time))
            case string:    return (*n)[i].Value.(string)    >= e.Value.(string)
         }
      }
      return false
   })
   if pos >= len(*n) || (*n)[pos].Value != e.Value {
      Goose.Logf(7,"Adding %s to set [%s]@%d",e,*n,pos)
      if pos == len(*n) {
         (*n) = append(*n,e)
      } else {
         (*n) = append(*n,Node{})
         copy((*n)[pos+1:],(*n)[pos:])
         (*n)[pos] = e
      }
   } else {
      Goose.Logf(7,"Avoid adding %s to set [%s]@%d",e,*n,pos)
   }
}

