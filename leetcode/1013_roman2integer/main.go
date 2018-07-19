// Package main provides roman to integer
package main

import "fmt"


func romanToInt(s string) int {
	letter2dict := map[string]int{
		"M": 1000,"CM": 900,"D": 500,"CD": 400,"C": 100,
		"XC": 90,"L": 50,"XL": 40,"X": 10,"IX": 9,"V": 5,"IV": 4,"I": 1,
	}
    buf := ""
	res := 0
    for i := 0; i < len(s); i++ {
		c := string(s[i])
		if buf == "" && (c == "I" || c == "X" || c == "C") {
			buf = c
		} else if (buf == "I" && (c == "V" || c == "X")) ||
                    (buf == "X" && (c == "L" || c == "C")) ||
                    (buf == "C" && (c == "D" || c == "M")) {
			buf = buf + c
			res += letter2dict[buf]
            // fmt.Println(2, buf)
			buf = ""
		} else {
            res += letter2dict[buf]
            // fmt.Println(3, buf)
			buf = c
		}
    }
    if buf != "" {
        res += letter2dict[buf]
        // fmt.Println(4, buf)
	}
    return res
}

func main() {
    fmt.Println("III = ", romanToInt("III")) // 3
    fmt.Println("IV = ", romanToInt("IV")) // 4
    fmt.Println("IX = ", romanToInt("IX")) // 9
    fmt.Println("LVIII = ", romanToInt("LVIII")) // 58
    fmt.Println("MCMXCIV = ", romanToInt("MCMXCIV")) // 1994
    fmt.Println("MDCCCLXXXIV = ", romanToInt("MDCCCLXXXIV")) // 1884
}
