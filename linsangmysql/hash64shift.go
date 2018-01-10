package linsangmysql

func hash64shift(key uint64) uint64 {
   key = (^key) + (key << 21)
   key = key ^ ((key >> 24))
   key = (key + (key << 3)) + (key << 8)
   key = key ^ (key >> 14)
   key = (key + (key << 2)) + (key << 4)
   key = key ^ (key >> 28)
   key = key + (key << 31)
   return key
}

