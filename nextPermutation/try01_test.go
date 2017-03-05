package main

import (
	"fmt"
	"testing"
)



func TestNextPerm(t *testing.T) {
	a := []int{0, 1, 2, 5, 3, 3, 0}
	//the next lexigraphical perm should be 0,1,3,0,2,3,5
	//see https://www.nayuki.io/page/next-lexicographical-permutation-algorithm
	nextPermutation(a)
	fmt.Println(a)
	b := []int{0, 1, 2, 3, 4, 5}
	nextPermutation(b)
	fmt.Println(b)
	b = []int{1,1,1,1,1,5}
	nextPermutation(b)
	fmt.Println(b)
}