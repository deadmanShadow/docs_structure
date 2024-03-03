    // a go program to determine even/odd
	package main

	import "fmt"
	func main(){
	  var number int
	  fmt.Printf("Enter any integer: ")
	  fmt.Scanf("%v",&number)
  
	  if number % 2 == 0 {
		fmt.Printf("Even\n")
	  }	else {
		fmt.Printf("Odd\n")
	  }
	}