/*
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

*/
package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	// 确保m是其中较短的数组
	if m > n {
		m, n = n, m
		nums1, nums2 = nums2, nums1
	}
	imin, imax, halfLen := 0, m, (m+n+1)/2
	for imin <= imax {
		i := (imin + imax) / 2
		j := halfLen - i
		if i < imax && nums2[j-1] > nums1[i] {
			// i 太小
			imin++
		} else if i > imin && nums1[i-1] > nums2[j] {
			imax--
		} else {
			maxLeft := 0
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				if nums2[j-1] > nums1[i-1] {
					maxLeft = nums2[j-1]
				} else {
					maxLeft = nums1[i-1]
				}
			}
			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}

			minRight := 0
			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else {
				if nums2[j] > nums1[i] {
					minRight = nums1[i]
				} else {
					minRight = nums2[j]
				}
			}
			return float64(maxLeft+minRight) / 2.0
		}
	}
	return 0.0

}

func findMedianSortedArraysBest(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	total := len1 + len2
	isEven := total%2 == 0
	val := 0.0

	i1 := 0
	i2 := 0

	for i := 1; i <= total/2+1; i++ {
		current := 0

		if i1 == len1 || (i2 != len2 && nums1[i1] > nums2[i2]) {
			current = nums2[i2]
			i2++
		} else {
			current = nums1[i1]
			i1++
		}

		if (isEven && i >= total/2) || (!isEven && i > total/2) {
			val += float64(current)
		}
	}

	if isEven {
		val /= 2
	}

	return val
}

func main() {

	fmt.Printf("[1, 3], [2] median=%f\n", findMedianSortedArrays([]int{1, 3}, []int{2}))

	fmt.Printf("[1, 2], [3, 4] median=%f\n", findMedianSortedArrays([]int{1, 2}, []int{3, 4}))

}
