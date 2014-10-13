package main

import (
  "fmt"
)

func main() {
  var total int

  fibo1 := 0
  fibo2 := 1
  var fibo3 int

  for fibo2 < 4000000 {
    fibo3 = fibo1 + fibo2
    fibo1 = fibo2
    fibo2 = fibo3
    if fibo3 % 2 == 0 {
      total += fibo3
    }
  }


  fmt.Println(total)
}