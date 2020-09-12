/*
Channels are a typed conduit through which you can send and receive
values with the channel operator, <-.
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.
*/

package main

import (
  "fmt"
  "time"
)

func main() {
  arr := []int{1,2,3,4,5}
  // declare a channel
  channel := make(chan int)
  now()
  // start a thread to sum.
  go sum(arr, channel)

  fmt.Println("Routine submitted...waiting on result..")
  // wait for the response from channel
  result := <- channel

  now()
  fmt.Println("Result received: Sum = ", result)

}

func sum(arr []int, channel chan int) {
  sum := 0
  for _,v := range arr {
    sum += v
  }
  fmt.Println("Sleeping for 5 sec in therad...")
  time.Sleep(5 * time.Second)
  fmt.Println("Returning sum", sum)
  channel <- sum // send sum to the channel.
}

func now() {
  fmt.Println("now:", time.Now())
}
