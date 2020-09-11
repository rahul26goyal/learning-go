/*
An interface type is defined by a set of methods.
A value of interface type can hold any value that implements those methods.
*/

package main

import "fmt"

/*
Define an interface with a method which should be implemented
*/
type Greeter interface {
  Greet()
}

type User struct {
  name string
  age int
}

/*
Implement the Interface required by Greeter to become a child.
*/
func (u *User) Greet() {
  fmt.Printf("Hello, %s. Age: %d .\n", u.name, u.age)
}

type Employee struct {
  name string
  age int
  dept string
}

func (u *Employee) Greet() {
  fmt.Printf("Hello, %s. Age: %d  dept: %s.\n", u.name, u.age, u.dept)
}


/*
A function which takes an interface as argument and calls the method.
*/
func DoGreet(obj Greeter) {
  obj.Greet()
}

func main() {
  u1 := &User{"Rahul", 23}

  u1.Greet()
  DoGreet(u1)

  e1 := &Employee{"Rahul", 23, "R&D"}
  e1.Greet()
  DoGreet(e1)
}
