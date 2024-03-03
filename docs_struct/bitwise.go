package main

import "fmt"
func main(){

  x := 18      // 18 = 10010
  y := 17      // 17 = 10001
  and := x & y   //16  = 10000
  or := x | y   //19  = 10011
  exor := x ^ y   //3  = 00011
  fmt.Printf("x & y = %v\n",and)
  fmt.Printf("x | y = %v\n",or)
  fmt.Printf("x ^ y = %v\n",exor)

}
