package main

import "fmt"
func main(){

  var radius, area float32

  fmt.Printf("Enter Radius : ")
  fmt.Scan(&radius)

  area = 3.1416 * radius * radius
  fmt.Printf("Area of circle: %v\n",area)

  // a program to calculate area of triangle
  // var base, height, area float32

  // fmt.Printf("Base = ")
  // fmt.Scan(&base)

  // fmt.Printf("Height = ")
  // fmt.Scan(&height)

  // area = 0.5 * base * height

  // fmt.Printf("Area of triangle = %v\n", area)
}
