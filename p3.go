package main

import (
  "fmt"
  "math"
)

func main() {
  x := 600851475143
  result := largestPrimeFactorFor(x)
  fmt.Println(result)
}

func largestPrimeFactorFor(x int) int {
  sqrt := int(math.Floor(math.Sqrt(float64(x))))
  var factorQueue []int

  for i := 2; i < sqrt; i++ {
    if x % i == 0 {
      counterpart := x / i
      if isPrime(counterpart) {
        return counterpart
      }
      factorQueue = append(factorQueue, i)
    }
  }
  for i := len(factorQueue) - 1; i >= 0; i-- {
    if isPrime(factorQueue[i]) {
      return factorQueue[i]
    }
  }
  //Question for further research: why is the compiler making me return this 0?
  //Is there a proper way to return a value to signify that the method did not locate a suitable answer?
  //eg in Ruby I would return nil in a situation like this
  return 0
}

func isPrime(x int) bool {
  if x == 2 {
    return true
  }
  sqrt := int(math.Floor(math.Sqrt(float64(x))))
  for i:= 2; i <= sqrt; i++ {
    if x % i == 0 {
      return false
    }
  }
  return true
}