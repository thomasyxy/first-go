package main

func main() {
  // goIf()
  goSwitch()
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

func goSwitch() {
  x := 100

  switch {
    case x > 0:
      println("x")
    case x < 0:
      println("-x")
    default:
      println("0")
  }
}
