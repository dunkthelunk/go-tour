package main

import "fmt"

var i,j int = 1,2

// Outside a function, every statement begins with a keyword
// k := 6 error

func main() {
  k := 3
  c, python, java := true, false, "no!"
  fmt.Println(i, j, k, c, python, java)
}
