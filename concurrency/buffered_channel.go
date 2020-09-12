/*
In this example, we will learn about Buffered channels in Go.
*/

package main

import (
  "fmt"
  "time"
)

func main() {
  // testSimpleBugger()
  time.Sleep(1 * time.Second)
  //testBufferOverFlow()
  testOverFlowWithRoutine()
  fmt.Println("\n\nGoing aheand..")
  // testEmptyBuffer()
}

func testSimpleBugger() {
  fmt.Println("Channels can be buffered. Provide the buffer length as the second argument to" +
    "make to initialize a buffered channel:")
  // create a buffer channel of size 3
  var bufChannel chan int = make(chan int, 3)
  // insert 2 elements.
  bufChannel <- 1
  bufChannel <- 2
  // pop one element
  fmt.Println("Elemeent: ", <- bufChannel)
  // insert 2 more.
  bufChannel <- 3
  bufChannel <- 4
  // pop 2 elements.
  fmt.Println("Elemeent: ", <- bufChannel)
  fmt.Println("Elemeent: ", <- bufChannel)
}

func testBufferOverFlow() {
  fmt.Println("Sends to a buffered channel block only when the buffer is full.")
  bufChannel := make(chan int, 2)
  bufChannel <- 1
  bufChannel <- 2
  bufChannel <- 3 // this will throw error: atal error: all goroutines are asleep - deadlock!
}

func testOverFlowWithRoutine() {
  fmt.Println("Sends to a buffered channel block only when the buffer is full. But goroutine will wait until the channel is available.")
  bufChannel := make(chan int, 2)
  bufChannel <- 1
  bufChannel <- 2

  someFunc := func() {
       fmt.Println(time.Now().Second(), "trying to push Element:")
       // this will wait until the channel has space i.e. pop in called in channnel.
       bufChannel <- 3
       fmt.Println(time.Now().Second(), "Pushed Element:")
  }
  go someFunc()
  time.Sleep(3 * time.Second)
  fmt.Println(time.Now().Second(), "Element:", <- bufChannel)
  time.Sleep(3 * time.Second)
  fmt.Println(time.Now().Second(), "Element:", <- bufChannel)
  time.Sleep(3 * time.Second)
  fmt.Println(time.Now().Second(), "Element:", <- bufChannel)
}

func testEmptyBuffer() {
  fmt.Println("Receives block when the buffer is empty.")
  bufChannel := make(chan int, 3)
  bufChannel <- 10
  fmt.Println("Element:", <- bufChannel)
  // fatal error: all goroutines are asleep - deadlock!
  fmt.Println("wating for Element:", <- bufChannel)
}
