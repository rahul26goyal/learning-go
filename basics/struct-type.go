
package main

import "fmt"

type User struct {
  name string
  age int
}

func (u User) Greet() {
  fmt.Printf("Hello, %s. Age: %d .\n", u.name, u.age)
}

/*
pass bu value method for User Struct
*/
func (u User) updateName(newName string) {
  u.name = newName
}

/*
pass by reference method for User Struct
*/
func (u *User) updateAge(newAge int) {
  u.age = newAge
}

func main() {
  u1 := User{"Rahul", 23}
  u1.Greet()
  u1.updateName("Rahul Goyal")
  u1.updateAge(50)
  u1.Greet()
}
