package strutil

import (
  "math/rand"
  "time"
  "unsafe"
)

const (
  letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
)

const letterLength = len(letterBytes)

func init(){
  rand.Seed(time.Now().Unix())
}

func RandomText(n int) string{
  b := make([]byte,n)
  for i := range b{
    b[i] = letterBytes[rand.Int()%letter]
  }
  return *(*string)(unsafe.Pointer(&b))
}