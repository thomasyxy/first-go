package main

func main() {
  goIf()
}

func goIf() {
  x := 100
  if x > 0 {
    println("x")
  } else if x < 0 {
    println("-x")
  } else {
    println("0")
  }
}
