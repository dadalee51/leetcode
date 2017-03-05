package main

import (
	"fmt"
)

func main() {
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

func nextPermutation(nums []int) {
	rp := -1
	if len(nums) < 2 {
		return
	}
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			//then find right most successor of nums[i-1]
			for j := len(nums) - 1; j > i-1; j-- {
				if nums[j] > nums[i-1] {
					rp = i
					nums[i-1], nums[j] = nums[j], nums[i-1]
					break
				}
			}
			//fmt.Println("after first swap:",nums)
			//then revert the slice from , excl. rp
			if rp > 0 {
				rev(nums[rp:])
			}
			return
		}
	}
	//no incr detected, simple revert entire slice.
	rev(nums)
	return
}

func rev(nums []int) {
	n := len(nums)
	i := 0
	j := n - 1
	//revert the list by swapping from both ends.
	for i < n/2 && j > (n-1)/2 {
		nums[i], nums[j] = nums[j], nums[i]
		i += 1
		j -= 1
	}
}
