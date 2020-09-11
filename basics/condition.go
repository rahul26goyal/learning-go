package main

import "fmt"

func main() {
   testEvenOdd(3)
   testEvenOdd(30)

   testGrade(100)
   testGrade(10)

   testSwitch(1)
   testSwitch(10)
}

func testEvenOdd(x int) {
  if x %2 == 0 {
    fmt.Println("Even")
    return
  }
  fmt.Println("Odd")
}


func testGrade(x int) {
  grade := ""
  if x > 90 {
    grade = "A"
  } else if x > 80 {
    grade = "B"
  } else if x > 70 {
    grade = "C"
  } else {
    grade = "D"
  }
  fmt.Printf("For Marks: %d the grade is: %s\n", x, grade)
}

func testSwitch(cond int) {
  fmt.Println("Testing Switch Case in Go")
  switch cond {
  case 0:
    fmt.Println("Case 0")
  case 1:
    fmt.Println("Case 1")
  default:
    fmt.Println("Case deefault")
  }
}
