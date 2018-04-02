package main
import "fmt"

func divide(num int)(x,y int){
  x = num * 2/3
  y = num - x
  return
}

func main(){
  fmt.Println(divide(24))
}
