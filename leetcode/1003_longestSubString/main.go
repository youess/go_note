package main

import (
	"fmt"
)

/*

Given a string, find the length of the longest substring without repeating characters.

Examples:

Given "abcabcbb", the answer is "abc", which the length is 3.

Given "bbbbb", the answer is "b", with the length of 1.

Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/
func lengthOfLongestSubstring(s string) int {

	maxLen := 0
	alpha := make(map[byte]int) // or array with letter or character encode

	n := len(s)
	for i, j := 0, 0; j < n; j++ { // [i, j]
		val, ok := alpha[s[j]]
		if ok && i < val { // 当遇到重复的值，更新其实下标
			i = val
		}
		if maxLen < j-i+1 {
			maxLen = j - i + 1
		}
		alpha[s[j]] = j + 1
	}
	/* wrong: not consider the duplicate case
	alpha := make(map[rune]int)
	for idx, e := range s {
		pre_idx, ok := alpha[e]
		if !ok {
			if maxLen < idx-pre_idx+1 {
				maxLen = idx - pre_idx + 1
			}
			alpha[e] = idx
		}
	}
	*/
	/* best solution:
	location := [256]int{}
	for i:= 0;i < len(location); i++ {
	    location[i] = -1
	}
	maxLen := 0
	left := 0
	for j := 0; j < len(s); j++ {
	    if location[s[j]]  >= left {
	        left = location[s[j]] +1
	    } else if j - left + 1 >maxLen {
	        maxLen = j - left +1
	    }
	    location[s[j]] =j
	}
	return maxLen
	*/
	return maxLen
}

func main() {
	fmt.Printf("abcabcbb max substring length(3): %d\n", lengthOfLongestSubstring("abcabcbb"))
	fmt.Printf("bbbbb max substring length(1): %d\n", lengthOfLongestSubstring("bbbbb"))
	fmt.Printf("pwwkew max substring length(3): %d\n", lengthOfLongestSubstring("pwwkew"))
}
