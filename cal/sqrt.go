package main

import (
  "fmt"
  "math"
)

func Sqrt(x float64) float64 {
  z := 1.0
  d := 1000.0

  for {
    old := z
    z = z - (z * z - x) / (2 * z)

    if uint(z * d) == uint(old * d) {
      fmt.Printf("Info: old= %v z=%v\n", old, z)
      return z
    }
  }
  return z
}

func main()  {
  fmt.Printf("Sqrt out: %v\n", Sqrt(2.0))
  fmt.Printf("math.Sqrt out: %v", math.Sqrt(2.0))
}
