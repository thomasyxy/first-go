package main

import (
  "strconv"
  "fmt"
)

func main()  {

  fun1()

  fun2()
}

func fun1()  {
  a, _ := strconv.ParseInt("1100100", 2, 32)
  b, _ := strconv.ParseInt("0144", 8, 32)
  c, _ := strconv.ParseInt("64", 16, 32)

  println(a, b, c)

  println("Ob" + strconv.FormatInt(a, 2))
  println("0" + strconv.FormatInt(a, 8))
  println("0x" + strconv.FormatInt(a, 16))
}

func fun2()  {
  var a float32 = 1.1234567899
  var b float32 = 1.12345678
  var c float32 = 1.123456781

  println(a, b, c)
  println(a == b, a == c)
  fmt.Println("%v %v, %v\n", a, b, c)
}
