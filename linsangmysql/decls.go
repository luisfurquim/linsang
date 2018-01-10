package linsangmysql

import (
   "errors"
   "github.com/luisfurquim/goose"
)

var Goose struct {
   Init     goose.Alert
   Read     goose.Alert
   Write    goose.Alert
   Search   goose.Alert
}

var (
   BufSize        int64 = 1024
   ErrInvalidType       = errors.New("Invalid type")
   ErrNotFound          = errors.New("Not found")
   ErrOpNotFound        = errors.New("Operand not found")
   ErrEOS               = errors.New("End of set")
   ErrWrongType         = errors.New("Wrong type")
   ErrInvalidPred       = errors.New("Invalid predicate")
   ErrUndefOp           = errors.New("Error operation not defined")
)
