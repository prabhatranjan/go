package main

import "fmt"

func main(){
  arr := []int{2,3,5,7,11,13}
  fmt.Println("Full Array is \n",arr)

  s := arr[0:4]
  printSlice(s);
}

func printSlice(s []int){
  fmt.Printf("Length = %d, Capacity = %d, %v\n",len(s),cap(s),s);
}
