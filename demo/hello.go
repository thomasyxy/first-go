package main

func main() {
  var x int32
  var s = "hello world!"
  println(x, s)
  inner()
}

func inner() {
  x := 100
  println(x)
}
