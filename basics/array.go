package main

import "fmt"

func main() {
  //testDeclarations()
  // testSliceing()
  //testRange()

  testExercise()

}

func testExercise() {
  //Given a list of names, you need to organize each name within a slice based on its length.
  var names = []string{"Katrina", "Evan", "Neil", "Adam", "Martin", "Matt",
	"Emma", "Isabella", "Emily", "Madison",
	"Ava", "Olivia", "Sophia", "Abigail",
	"Elizabeth", "Chloe", "Samantha",
	"Addison", "Natalie", "Mia", "Alexis"}
  result := make([][]string, 100) // assuming the max lenght as 100 for any name
  max := 0
  for _, v := range names {
    len := len(v)
    result[len] = append(result[len], v)
    if len > max {
      max = len
    }
  }
  max = max + 1 // incease indeex
  fmt.Printf("Outpur : %v \n", result[:max])
}

func testRange() {
  arr := []int{1,2,3,4,5}
  fmt.Println("Test testRange Operations....", arr)
  for i,v := range arr {
    fmt.Println(i,v)
  }

  a2 := make([]int, 5)
  fmt.Println("Empty array", a2)
  val := 10
  for i := range a2 {
    a2[i] = val
    val = val + 10
  }
  fmt.Println("Filled array", a2)
}

func testSliceing() {
  arr := []int{1,2,3,4,5,6,7}
  fmt.Println("Test Slicing Operations....", arr)

  fmt.Println("Slice 1:3", arr[1:3])
  fmt.Println("Slice first :3", arr[:3])
  fmt.Println("Slice from 2nd:", arr[2:])
}

func testDeclarations() {
  var a1[5] int
  fmt.Println(a1)

  a1[1] = 10
  a1[2] = 20
  fmt.Println("modifiled array", a1)
  //fmt.Printf("Array a1 %s \n", a1)
  //fmt.Printf("Array a1 %q \n", a1)

  a2 := [3]string{"a", "b", "c"}
  fmt.Println("String array", a2)
  // fmt.Println("String array", a2[3]) -- index out of Bound.

  a3 := [...]int{1,2,3,4,5}  // implic lenght in allocated.
  fmt.Println("3.array", a3, ", Lenght:", len(a3))
  // a3[5]= 12 - out of index as the size if fixed.
}
