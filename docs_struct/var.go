package main
import "fmt"
func main(){
  // static variable declaration
  var name, country string
  var age int
  var gpa float32

  // variable initialization
  name = "Raihan Shamil"
  country = "Bangladesh"
  age = 25
  gpa = 4.56

  // dynamic variable declaration + initialization
 //  var name = "raihan shamil"
 //  var country = "Bangladesh"
 //  var age = 25
 //  var gpa = 4.56



  fmt.Println(name)
  fmt.Println(name, " is a student")
  fmt.Println(name,"is", age, "years old")
  fmt.Println(name, "has got" , gpa, "/5 in SSC")
  fmt.Print(name,"originally from ", country);
}

