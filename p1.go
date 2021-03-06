//If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

//Find the sum of all the multiples of 3 or 5 below 1000.

package main

import (
  "fmt"
)

func main() {
  var total int

  for i := 3; i < 1000; i += 3 {
    total += i
  }
  for i := 5; i < 1000; i += 5 {
    if i % 3 != 0 {
      total += i
    }
  }

  fmt.Println(total)
}