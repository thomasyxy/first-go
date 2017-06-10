package main

import (
  "errors"
  "fmt"
)

func div(a, b int) (int, error) {
  if b == 0 {
    return 0, errors.New("division by zero")
  }

  return a / b, nil
}

func close(x int) func() {
  return func() {
    println(x)
  }
}

func delay(a, b int) {
  defer println("dispose...")
  println(a / b)
}

func main() {
  a, b := 10, 2
  c, err := div(a, b)

  fmt.Println(c, err)

  x := 100
  f := close(x)
  f()

  // 报错panic: runtime error: integer divide by zero
  delay(10, 0)
}
