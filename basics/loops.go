package main

import "fmt"

func main() {
   simpleFor()

   whileLoop()
}

func whileLoop() {
  fmt.Println("Syntax for while type condition...")
  i := 0
  sum := 0
  for i < 10 {
    sum += i
    i = i + 1
  }
  fmt.Println("Sum=", sum)

}


func simpleFor() {
  for i := 0; i < 10; i++ {
    fmt.Println("Ii=", i)
  }

  sum := 1
  for ; sum < 1000; {
   sum += sum
 }
 fmt.Println("sum=", sum)
}
