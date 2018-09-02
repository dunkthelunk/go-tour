package main

import "fmt"

//should be used only in short functions
func split(sum int) (x, y int) {
  x = sum - 1
  y = sum
  return
}

func main() {
  fmt.Println(split(24))
}
