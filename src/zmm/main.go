package main

import (
  "fmt"
  "reflect"
)

func main(){
  fmt.Println("我"[0])

  for k,v := range "我们"{
    fmt.Println(k)
    fmt.Println(reflect.TypeOf(v))
  }
}