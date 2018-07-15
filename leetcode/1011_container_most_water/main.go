// Given n non-negative integers a1, a2, ..., an, where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

// Note: You may not slant the container and n is at least 2.

package main

import "fmt"

func max(e1 int, e2 int) int {
	if e1 > e2 {
		return e1
	} else {
		return e2
	}
}

func min(e1 int, e2 int) int {
	if e2 > e1 {
		return e1
	} else {
		return e2
	}
}

// 用暴力求解，O(n**2)
func maxAreaBruteForce(height []int) int {
	maxA := 0
	n := len(height)
	li, hi := 0, 0
	for li = 0; li < n-1; li++ {
		for hi = li+1; hi < n; hi++ {
			tmpA := min(height[li], height[hi]) * (hi - li)
			maxA = max(maxA, tmpA)
		}
	}
	return maxA
}

// 双指针方法，想象成一个二次矩阵的一半，然后不断比较两两的面积
// 反证法: 假如存在另外一个比这个解更大的面积，o_li, o_hi
// 由于之前计算了最大面积, li和hi肯定访问过o_li或o_hi中的一个。
// [... a1 ... b1 ... a2 ... b2 ...]
func maxArea(height []int) int {
	maxA := 0
	li, hi := 0, len(height)-1
	for li < hi {
		tmpA := min(height[hi], height[li]) * (hi - li)
		maxA = max(maxA, tmpA)
		if height[hi] > height[li] {
			li++
		} else {
			hi--
		}
	}
	return maxA
}


func main() {
	testHeight := []int{1, 3, 9, 5, 10}
	m := maxArea(testHeight)
	// m := maxAreaBruteForce(testHeight)
	fmt.Printf("max area: %d\n", m)
}
