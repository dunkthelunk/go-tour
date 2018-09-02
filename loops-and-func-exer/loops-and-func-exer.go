package main

import (
  "fmt"
)
func abs(x float64) float64 {
  if x > 0.0 {
    return x;
  }
  return -x;
}
func Sqrt2(x float64) float64 {
  znew := 1.0
  zold := 0.0
  for abs(znew - zold) > 0.00000001 {
    zold = znew
    znew = zold - (zold*zold - x)/(2*zold)
//    fmt.Printf("%g\n", znew)
  }
  return znew
}

func Sqrt(x float64) float64 {
  z := 1.0
  for i:=1; i < 11; i++ {
    z -= (z*z - x) / (2*z)
   // fmt.Printf("%g\n", z)
  }
  return z
}

func main() {
  fmt.Println(Sqrt(2))
  fmt.Println(Sqrt2(2))
}


