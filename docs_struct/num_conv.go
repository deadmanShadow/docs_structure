package main

import "fmt"
func main(){


  // string formatting
  // var name = "raihanshamil"
  // fmt.Printf("%s\n",name)
  // fmt.Printf("%q\n",name)

	// floating number formatting
  // var number = 3.1416
  // fmt.Printf("%.2f\n",number)

  var decimalNumber int

  fmt.Printf("decimal number= ")
  fmt.Scanf("%v", &decimalNumber)

  fmt.Printf("Binary Number = %b\n",decimalNumber)
  fmt.Printf("Octal Number = %o\n",decimalNumber)
  fmt.Printf("Hexa Number = %x\n",decimalNumber)

}
