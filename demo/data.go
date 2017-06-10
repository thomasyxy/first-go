package main

import (
  "fmt"
)

func sliceData()  {
  x := make([]int, 0, 5)  // 创建容量为5的切片

  for i := 0; i < 8; i++ {
    x = append(x, i)
  }

  fmt.Println(x)
}

func mapData()  {
  m := make(map[string]int)

  m["a"] = 1

  x, ok := m["b"]

  fmt.Println(x, ok)

  delete(m, "a")
}

func main()  {
  sliceData()
  mapData()
}
