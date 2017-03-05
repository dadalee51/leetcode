package main

import (
	"fmt"
)

func main() {
	c := findPairs([]int{3,1,4,1,5}, 2)
	fmt.Println(c)
	c = findPairs([]int{1, 2, 3, 4, 5}, 1)
	fmt.Println(c)
	c = findPairs([]int{1, 3,1,5,4}, 0)
	fmt.Println(c)
	c = findPairs([]int{1, 1,1,1,1}, 0)
	fmt.Println(c)
}

func findPairs(nums []int, k int) int {
	c := 0
	i := len(nums) - 1
	m := make(map[int]int)

	for i >= 0 {
		for j := 0; j < i; j++ {
			if nums[i] == nums[j] {
				//edge case for when k is 0, we add the repeated elements to self maps. 
				if k==0{ 
					m[nums[i]]=nums[i]
				}
				nums = append(nums[:i], nums[i+1:]...)
				//fmt.Println("new nums:", nums)
				break
			}
		}
		i -= 1
	}
	i = len(nums) - 1
	for i >= 0 {
		for j := 0; j < i; j++ {
			kd := nums[i]-nums[j]
			if kd < 0 { kd = -kd}
			if kd == k {
				//fmt.Print("[",nums[i],nums[j],"]")
				c++
			}
		}
		i -= 1
	}
	//count unique item in maps
	if k==0 {
		c+=len(m)
	}
	
	return c
}
