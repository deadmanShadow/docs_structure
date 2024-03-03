package main

import "fmt"

func square (number int) int{
  return number * number
}

func displayMessage(country string){
  fmt.Printf("I love my motherland: %v\n", country)
}

func main(){
  displayMessage("Bangladesh")
  displayMessage("India")
  fmt.Printf("Square of 6 = %v\n",square(6))
  fmt.Printf("Square of 7 = %v\n",square(7))
}