package linsangmysql

import (
   "github.com/luisfurquim/linsang"
)

/*

Returns the current node of the set.
Returns the ErrEOS error if all the nodes of the set are already received (similar to the os.EOF error).

*/
func (set *MySQLSet) GetNode() (*linsang.Node, error) {

   Goose.Read.Logf(6,"GetNode from : %#v",set)
   if (set.curr < int64(len(set.vertex))) && (set.curr >= 0) {
      return &set.vertex[set.curr], nil
   }

   return nil, ErrEOS
}

