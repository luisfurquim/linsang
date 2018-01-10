package linsangmysql

import (
   "strings"
)

func hash64shiftString(key ...string) int64 {
   var h     uint64
   var buf []byte
   var i     int

   for i=0; i<len(key); i++ {
      if (len(key[i])%8) != 0 {
         key[i] += strings.Repeat(" ",8 - (len(key[i])%8))
      }

      for buf=[]byte(key[i]); len(buf)>0; buf=buf[8:] {
         h ^= hash64shift((uint64(buf[7])<<56) | (uint64(buf[6])<<48) | (uint64(buf[5])<<40) | (uint64(buf[4])<<32) | (uint64(buf[3])<<24) | (uint64(buf[2])<<16) | (uint64(buf[1])<<8) | uint64(buf[0]))
      }
   }

   return int64(h&0x7fffffffffffffff)
}
