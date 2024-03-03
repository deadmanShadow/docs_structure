package main

import "fmt"

func change (val int) {
  val = 8
}

func changeWihPointer (val *int) {
  *val = 8
}

func main(){
  x:=10
  change(x)
  fmt.Println(x)

  // call by reference
  changeWihPointer(&x)
  fmt.Println(x)
}