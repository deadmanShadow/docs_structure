package main

import "fmt"

 type Student struct{
	name string
	age int
	id int
  }


func main(){

   <!-- type Student struct{
	name string
	age int
	id int
  } -->

  rahim := Student{"Shamil", 25,1302020017}
  rahim.age = 30
  fmt.Println(rahim)
  fmt.Println(rahim.id)
  fmt.Println(rahim.name)
  fmt.Println(rahim.age)
}

// passing struct in a function
package main

import "fmt"

type Student struct {
  name string
  age int
  id int
}

func displayInfo(s Student) {
  fmt.Println(s.name)
  fmt.Println(s.age)
}
func main() {
  s1 := Student{"Shamil", 25, 1302020017}
  displayInfo(s1)
}

// if we need to change data
package main

import "fmt"

type Student struct {
  name string
  age int
  id int
}

func (x *Student) increaseAge(val int) {
  x.age += val
}

func main() {
  s1 := Student{"Shamil", 25}
  s1.increaseAge(1)
  fmt.Println(s1.age)
}

// timer app
// Employee management app