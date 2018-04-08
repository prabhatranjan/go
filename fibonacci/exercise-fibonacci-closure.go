/*
Implement a fibonacci function that returns a function (a closure)
that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).
*/

package main
import (
  "fmt"
)

func fibonacci() func() int{
  a,b := 0,1
  var c int
  fmt.Println(a)
  fmt.Println(b)
  return func() int{
    c=a+b
    a = b
    b = c
    return c
  }
}

func main(){
  f := fibonacci()
  for i:=0;i<10;i++{
    fmt.Println(f())
  }
}
