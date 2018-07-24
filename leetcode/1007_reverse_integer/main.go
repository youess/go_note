//Given a 32-bit signed integer, reverse digits of an integer.

//Example 1:

//Input: 123
//Output: 321
//Example 2:

//Input: -123
//Output: -321
//Example 3:

//Input: 120
//Output: 21
package main

import "fmt"

func pow10(n int) int {
    ret := 1
    for n != 0 {
        ret *= 10
        n--
    }
    return ret
}

// consider negative
// consider overflow of int32
// [-2^31, 2^31-1]
// -2147483648, 2147483647
func reverse(x int) int {

    flag := 0
    if x < 0 {
        flag = 1
        x = -x
    }
    digits := make([]int, 0)
    for x != 0 {
        d := x % 10
        x = x / 10
        digits = append(digits, d)
    }
    //fmt.Printf("%v\n", digits)
    rx := 0
    n := len(digits)
    for i := 0; i < n; i++ {
        rx += digits[i] * pow10(n-i-1)
    }

    if rx > 2147483648 {  // need a bigger int type than int32
        return 0
    }

    if flag == 1 {
        rx = -rx
    }
    //fmt.Println(rx)
    return rx
}

const (
    INT_MAX = 2147483647
    INT_MIN = -2147483648
)

func reverse2(x int) int {
    rx := 0
    for x != 0 {
        pop := x % 10
        x /= 10
        if rx > INT_MAX / 10 || (rx == INT_MAX / 10 && pop > 7) {
            return 0
        }
        if rx < INT_MIN / 10 || (rx == INT_MIN / 10 && pop < -8) {
            return 0
        }
        // no need to store tmp digits
        rx = rx * 10 + pop
    }
    return rx
}

func main() {
    //
    //fmt.Println(pow10(3))
    fmt.Println(reverse2(123))
    fmt.Println(reverse2(210))
    fmt.Println(reverse2(-123))
    fmt.Println(reverse2(1234567809))
}
