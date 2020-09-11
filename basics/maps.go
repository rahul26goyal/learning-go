package main

import "fmt"

func main() {
  testInitialization()

  testIterate()
}

func testIterate() {
  var capital = map[string]string{ "Odisha": "Bhu", "AP": "Hyderabad", "Tn": "Chennai"}
  fmt.Println("Iterating the Maps and priting it,.")
  for k, v := range capital {
    fmt.Println("S:", k, "C:",v)
  }
}

func testInitialization() {
  var capital = map[string]string{ "Odisha": "Bhu", "AP": "Hyderabad", "Tn": "Chennai"}
  fmt.Printf("%#v \n", capital) // # to print thee schema
  fmt.Println("Capital map:", capital)

  var nums = make(map[int]string)
  nums[1] = "one"
  nums[100] = "hundred"
  fmt.Printf("Number: %v \n", nums)
}
