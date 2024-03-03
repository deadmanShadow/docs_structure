package main

import "fmt"
func main(){
	var number1, number2, number3 int
	fmt.Printf("Enter 3 numbers: ")
	fmt.Scanf("%v %v %v",&number1, &number2, &number3)

	if number1 > number2 {
		if number1 > number3 {
			fmt.Printf("%v is the largest number\n", number1)
		}else{
			fmt.Printf("%v is the largest number\n", number3)
		}
	}	else if number2 > number1 {
		if number2 > number3 {
			fmt.Printf("%v is the largest number\n", number2)
		}else{
			fmt.Printf("%v is the largest number\n", number3)
		}
	}	else {
		fmt.Printf("Numbers are equal\n")
	}
}