package main

import (
	"fmt"
)

/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/
func twoSum(nums []int, target int) []int {
	/*
		d := make(map[int]int)
		for idx, e := range nums {
			d[e] = idx
		}
		// fmt.Println(d)
		solution := make([]int, 2)
		for idx, e := range nums {
			oe, ok := d[target-e]
			if ok && idx != oe {        // 需要检查是否自己相同使用两次
				solution[0] = idx
				solution[1] = oe
				return solution
			}
		}
	*/
	d := make(map[int]int)
	for idx, e := range nums {
		oe, ok := d[target-e]
		// if ok && idx != oe {
		if ok { // no need to check
			return []int{oe, idx}
		} else {
			d[e] = idx
		}
	}

	return nil
}

// 如果需要遍历所有解
func twoSumMultiple(nums []int, target int) [][2]int {

	d := make(map[int]int)
	sol := make([][2]int, 0)
	for idx, e := range nums {
		oe, ok := d[target-e]
		if ok {
			sol = append(sol, [2]int{oe, idx})
		} else {
			d[e] = idx
		}
	}
	if len(sol) > 0 {
		return sol
	}
	return nil
}

func main() {

	nums := []int{2, 7, 11, 15}
	target := 9
	sol := twoSum(nums, target)
	fmt.Println(sol)
	nums = []int{2, 3, 6}
	target = 6
	fmt.Println("")
	fmt.Println(twoSum(nums, target))

	nums = []int{1, 2, 3, 4, 5, 6}
	target = 5
	fmt.Println("")
	fmt.Println(twoSumMultiple(nums, target))
}
