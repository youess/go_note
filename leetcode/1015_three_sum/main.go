//Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0?
//Find all unique triplets in the array which gives the sum of zero.

//Note:

//The solution set must not contain duplicate triplets.
package main

import "fmt"
import "sort"

func threeSum_force(nums []int) [][]int {
    sort.Ints(nums)
    res := make([][]int, 0)
    var i, j, k int
    for i=0; i < len(nums)-2; i++ {
        for j=i+1; j < len(nums)-1; j++ {
            for k=j+1; k < len(nums); k++ {
                s := nums[i] + nums[j] + nums[k]
                if s == 0 {
                    res = append(res, []int{nums[i], nums[j], nums[k]})
                }
            }
        }
    }
    return res
}

/* slice version
func twoSum(nums []int, start int, end int, target int, res *[][]int) {
    l, r := start, end
    for l < r {
        s := nums[l] + nums[r] + target
        if s == 0 {
            *res = append(*res, []int{target, nums[l], nums[r]})
            //res = append(res, []int{target, nums[l], nums[r]})   // 如果不用指针传入，这样是不能修改外面的数据的
            for l < r && nums[l] == nums[l+1] {
                l++        // skip left dumplicate
            }
            for l < r && nums[r] == nums[r-1] {
                r--
            }
            l++
            r--
        } else if s < 0 {
            l++
        } else {
            r--
        }
    }
}

// runtime: 960ms
func threeSum(nums []int) [][]int {
    n := len(nums)
    res := make([][]int, 0)
    if n < 3 {
        return res
    }
    // 进行排序，否则后续不好判断移动方向
    sort.Ints(nums)
    for i := 0; i < n-2; i++ {   // 3 sum
        if i > 0 && nums[i] == nums[i-1] {
            continue      // 跳过重复答案
        }
        twoSum(nums, i+1, n-1, nums[i], &res)
    }
    return res
}
*/

// 没有比上面快，应该是leetcode服务器问题
func threeSum(nums []int) [][]int {
    n := len(nums)
    res := [][]int{}
    if n < 3 {
        return res
    }
    // 进行排序，否则后续不好判断移动方向
    sort.Ints(nums)
    for i := 0; i < n-2; i++ {   // 3 sum
        if i > 0 && nums[i] == nums[i-1] {
            continue      // 跳过重复答案
        }
        l, r := i+1, n-1
        target := nums[i]
        for l < r {
            s := nums[l] + nums[r] + target
            if s == 0 {
                res = append(res, []int{target, nums[l], nums[r]})
                for l < r && nums[l] == nums[l+1] {
                    l++        // skip left dumplicate
                }
                for l < r && nums[r] == nums[r-1] {
                    r--
                }
                l++
                r--
            } else if s < 0 {
                for l < r && nums[l] == nums[l+1] {
                    l++        // skip left dumplicate
                }
                l++
            } else {
                for l < r && nums[r] == nums[r-1] {
                    r--
                }
                r--
            }
        }
    }
    return res
}

func main() {
    // fmt.Println(threeSum_force([]int{-1, 0, 1, 2, -1, -4}))
    fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
