package main

import "fmt"

func main() {
 var a [2]int
 p := [6]int{5,6,7,8,9}
 fmt.Println(a,p)
 s := "my club中文"
 for i, r := range s {
	fmt.Println(i, r)
 }
 
}

