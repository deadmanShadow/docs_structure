  // a go program to spell digits with switch statement
  package main

  import "fmt"
  func main(){
	  var digit int
	  fmt.Printf("Enter a digit (0-9): ")
	  fmt.Scanf("%v",&digit)
  
	  switch digit {
	  case 0:
		  fmt.Printf("Zero\n")
	  case 1:
		  fmt.Printf("One\n")
	  case 2:
		  fmt.Printf("Two\n")
	  case 3:
		  fmt.Printf("Three\n")
	  case 4:
		  fmt.Printf("Four\n")
	  case 5:
		  fmt.Printf("Five\n")
	  case 6:
		  fmt.Printf("Six\n")
	  case 7:
		  fmt.Printf("Seven\n")
	  case 8:
		  fmt.Printf("Eight\n")
	  case 9:
		  fmt.Printf("Nine\n")
  
	// multiple cases possible
	// case 0,1,2,3,4:
	//   fmt.Printf("less than 5" )
	  default:
		  fmt.Printf("Not a digit\n")
	  }
  
  }
  

//   same as c, c++, java, js but only difference you do not use break statement. only run the mateched statement