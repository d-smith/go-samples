package main

import (
  "fmt"
  "hash/crc32"
  "crypto/sha1"
  )


func main() {
  h := crc32.NewIEEE()
  h.Write([]byte("test"))
  v := h.Sum32()
  fmt.Println(v)

  h2 := sha1.New()
  h2.Write([]byte("test"))
  bs := h.Sum([]byte{})
  fmt.Println(bs) 
}
