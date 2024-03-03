      // with if only
      package main

      import "fmt"
      func main(){
        // a program to determine positive / negative / zero
        number := 10
        if number > 0 {
          fmt.Printf("Positive\n")
        }
        if number < 0 {
          fmt.Printf("Negative\n")
        }
        if number == 0 {
          fmt.Printf("Zero\n")
        }
      }

      // with if, else if
      package main

      import "fmt"
      func main(){
        // a program to determine positive / negative / zero
        number := 10
        if number > 0 {
          fmt.Printf("Positive\n")
        }	else if number < 0 {
          fmt.Printf("Negative\n")
        } else if number == 0 {
          fmt.Printf("Zero\n")
        }
      }

      // with if, else if, else
      package main

      import "fmt"
      func main(){
        // a program to determine positive / negative / zero
        number := 10
        if number > 0 {
          fmt.Printf("Positive\n")
        }	else if number < 0 {
          fmt.Printf("Negative\n")
        } else {
          fmt.Printf("Zero\n")
        }
      }