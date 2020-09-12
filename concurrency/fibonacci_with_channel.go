/*
In this example we will show how we can use a channel to pass values
from the thread via a chennel.
*/

package main

import (
  "fmt"
  "time"
)

func fibinacci(n int, channel chan int) {
  a := 0
  b := 1
  if n == 0 {
    return
  } else if n == 1 {
    channel <- a
  }  else if n == 2 {
    channel <- a
    channel <- b
  } else {
    time.Sleep(2 * time.Second) // dummy initial sleep.
    for i := 1; i<=n; i++ {
      c := a + b
      fmt.Println("Sending c:", a)
      channel <- a // this will be blocking until the chennel has space.
      time.Sleep(200 * time.Millisecond) // adding this sleep just to check if value is poped.
      a = b
      b = c
    }
  }
  time.Sleep(200 * time.Millisecond)
  // close the channel, else thee for loop will not end .
  close(channel)
}

func main() {
  n := 10
  // create a channle with buffer space of 5.
  var channel chan int = make(chan int, 5)
  // trigger the go routine.
  go fibinacci(n, channel)
  // start reading the values from channel.
  for i := range channel {
    // time.Sleep(5 * time.Second) // uncomment to check the blocking in thread.
    fmt.Println("Response:", i)
  }
}
