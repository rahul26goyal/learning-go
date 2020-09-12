/*
A goroutine is a lightweight thread managed by the Go runtime.
*/
package main

import (
  "fmt"
  "time"
)

func simple(msg string) {
  for i := 0; i < 10; i++ {
    fmt.Println("Message: ", i, ":", msg)
    time.Sleep(1000 * time.Millisecond)
  }
}
func main() {
  go simple("Go Routine..") // executes in a separate thread.
  time.Sleep(10000 * time.Millisecond) // waits here..
  simple("Go Non-Routine....") // finally this is executed..
}
