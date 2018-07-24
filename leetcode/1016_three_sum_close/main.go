//Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target. 
//Return the sum of the three integers. You may assume that each input would have exactly one solution.

//Example:

//Given array nums = [-1, 2, 1, -4], and target = 1.

//The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
package main


import "fmt"
import "sort"


func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

//func min(a int, b int) int {
//    if a < b {
//        return a
//    }
//    return b
//}

func threeSumClosest(nums []int, target int) int {
    // 确定数据长度大于等于3

    ms := nums[0] + nums[1] + nums[2]
    length := len(nums)

    sort.Ints(nums)

    for i := 0; i < length-2; i++ {
        l, r := i+1, length-1
        //fmt.Println(i)
        for l < r {
            s := nums[i] + nums[l] + nums[r]
            //fmt.Println(nums[i], nums[l], nums[r], s)
            if abs(s - target) < abs(ms - target) {
                ms = s
            }
            if s < target {
                l++
            } else {
                r--
            }
        }
    }
    return ms

}

func main() {
    //fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
    //fmt.Println(threeSumClosest([]int{0, 1, 2, -3}, 1))
    fmt.Println(threeSumClosest([]int{1, 1, -1, -1, 3}, -1))
}
