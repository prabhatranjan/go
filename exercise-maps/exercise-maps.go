/*
Implement WordCount. It should return a map of the counts of each “word” in the string s.
The wc.Test function runs a test suite against the provided function and prints success or failure.
*/

package main
import (
  "golang.org/x/tour/wc"
  "strings"
  )

func WordCount(s string) map[string]int{
  words := strings.Fields(s)
  m := make(map[string]int)
  for i := range words{
    m[words[i]]=m[words[i]]+1;
  }
  return m
}

func main(){
  wc.Test(WordCount)
}
