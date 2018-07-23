// https://leetcode.com/problems/longest-common-prefix/solution/
package main

import (
    "fmt"
    "strings"
)

func isCommonPrefix(strs []string, length int) bool {

    str1 := strs[0][:length]
    for _, s := range strs {
        if !strings.HasPrefix(s, str1) {
            return false
        }
    }
    return true

}

func min(i int, j int) int {
    if i < j {
        return i
    }
    return j
}

func LCPBest(strs []string) string {

    if len(strs) == 0 {
        return ""
    }
    minLen := 2147483647   // max of int32
    for _, s := range(strs) {
        minLen = min(minLen, len(s))
    }
    low, high := 1, minLen
    for low <= high {
        mid := (low + high) / 2
        if isCommonPrefix(strs, mid) {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return strs[0][:(low+high)/2]

}


// 另外可以根据构建tries树或者二叉搜索进行构建
func longestCommonPrefix(strs []string) string {

    if len(strs) == 0 {
        return ""
    }

    prefix := make([]string, 0)
    idx := 0
    flag := false

    EXIT:
    for {
        for _, s := range strs {
            if len(s) == 0 {
                return ""
            } else if idx < len(s) && len(prefix) <= idx {
                prefix = append(prefix, string(s[idx]))
                flag = true
                // fmt.Printf("if1 - %s - %d\n", s, idx)
            } else if idx < len(s) && string(s[idx]) == prefix[idx] {
                // fmt.Printf("if2 - %s - %d\n", s, idx)
                continue
            } else {
                // fmt.Printf("if3 - %s - %d\n", s, idx)
                if flag {
                    prefix = prefix[:len(prefix)-1]
                }
                break EXIT
            }
        }
        idx++
        flag = false
    }

    return strings.Join(prefix, "")
}


func main() {
    //prefix := longestCommonPrefix([]string{"abcede", "abc", "abcefj3o"})
    //fmt.Println(prefix)
    //prefix := longestCommonPrefix([]string{"c", "c", "c"})
    prefix := LCPBest([]string{"abcede", "abc", "abcefj3o"})
    //prefix := LCPBest([]string{"c", "c", "c"})
    fmt.Println(prefix)
}
