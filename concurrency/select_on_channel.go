/*
The select statement lets a goroutine wait on multiple communication operations.
A select blocks until one of its cases can run, then it executes that case.
It chooses one at random if multiple are ready.
*/

package main

import (
  "fmt"
  "time"
)

/*
A function which will consume inputs from multiple Channel
using the select construct.
*/
func testSelect(activeChannel, quitChennel chan int) {
  print("starting testSelect in go routine..")
  for {
    select {
    case a := <- activeChannel:
      print("Recived on active channel..")
      fmt.Println("Value recied:", a)
      time.Sleep(2 * time.Second)
    case <- quitChennel:
      print("Quit Chennnel reveived...")
      return
    default:
      print("No Op received..executing defalut..sleep 2")
      time.Sleep(2 * time.Second)
    }
  }
}

/*
A function to send inputs to the channel.
*/
func sendInputs(channel, quitChannel chan int) {
  for i := 0; i< 10; i++ {
    print("Send input to channel..")
    //send input to active channel.
    channel <- i
    //time.Sleep(1 * time.Second) // comment this if u want to see the random choise by select.
  }

  time.Sleep(5 * time.Second)
  // send input to Quit Channel.
  print("Sending Quit Signal to Channel...")
  quitChannel <- 123
}

func main() {
  // create 2 channel for our use case.
  numChannle := make(chan int, 15)
  quitChannel := make(chan int, 2)
  // trigger a thread which will dump values into the channel.
  go sendInputs(numChannle, quitChannel)
  // trigger go routine for select o/p
  testSelect(numChannle, quitChannel)
  time.Sleep(5 * time.Second)
  fmt.Println("BYE.")
}

/*
Utiliti to print time along with message.
*/
func print(msg string) {
  fmt.Println(time.Now().Minute(), ":", time.Now().Second(), msg)
}
