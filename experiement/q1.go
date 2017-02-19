package main

import (
	"fmt"
)

func main(){
	fmt.Println(len("AAA\nAAA"))
	fmt.Println("AAA\nAAA")

	fmt.Println(len("BBB\tBBB"))
	fmt.Println("BBB\tBBB")
	
	fmt.Println([]string{"hello", "world"})
	
	fmt.Printf("%-6s%6s ( )\n", "12", "345")
	
	fmt.Printf("%d\n", 14%10 + 14/10)
	
}