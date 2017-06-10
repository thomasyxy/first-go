package main

func main() {
  // goIf()
  // goSwitch()
  goFor()
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

func goFor() {
  for i := 0; i < 5; i++ {
    println(i)
  }

  for i := 4; i > 0; i-- {
    println(i)
  }

  x := 0
  for x < 5 {
    println(x)
    x++
  }

  y := 4
  for {
    println(y)
    y--

    if y < 0 {
      break
    }
  }

  z := []int{100, 101, 102}
  for i, n := range z {
    println(i, ":", n)
  }
}
