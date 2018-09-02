package main

import (
  "fmt"
  "math"
)

func sqrt(x float64) string {
  if x < 0 {
    return sqrt(-x) + "i"
  }
  return fmt.Sprint(math.Sqrt(x))
}
func pow(x, n, lim float64) float64 {
  if v:=math.Pow(x,n); v < lim {
    return v
  }
  return lim
}

func main() {
  fmt.Println(pow(2, 3, 6), pow(2,3,15))
  fmt.Println(sqrt(2), sqrt(-4))
  fmt.Printf("Type of int srqt is %T\n", sqrt(2))
}


