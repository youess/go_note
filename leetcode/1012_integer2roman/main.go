// Package main provides solution of 1012
package main

import (
	"fmt"
	"strings"
)

func intToRoman(num int) string {
    valArr := []int{1000,900,500,400,100,90,50,40,10,9,5,4,1}
    charArr := []string{"M","CM","D","CD","C","XC","L","XL","X","IX","V","IV","I"}

    length := len(valArr)
    result := ""
    i := 0
    for  {
        if i >= length || num == 0 {
            break
        }
        tmp := num - valArr[i]
        if tmp >= 0 {
            num = tmp
            result += charArr[i]
        } else {
            i++
        }
    }

    return result
}


func int2roman(num int) string {
	// romanDict := getIntRomanDict()
	s := make([]string, 0)
	for num > 0 {
		switch {
			case num >= 1000:
				s = append(s, "M")
				num -= 1000
			case num >= 900:
				s = append(s, "CM")
				num -= 900
			case num >= 500:
				s = append(s, "D")
				num -= 500
			case num >= 400:
				s = append(s, "CD")
				num -= 400
			case num >= 100:
				s = append(s, "C")
				num -= 100
			case num >= 90:
				s = append(s, "XC")
				num -= 90
			case num >= 50:
				s = append(s, "L")
				num -= 50
			case num >= 40:
				s = append(s, "XL")
				num -= 40
			case num >= 10:
				s = append(s, "X")
				num -= 10
			case num >= 9:
				s = append(s, "IX")
				num -= 9
			case num >= 5:
				s = append(s, "V")
				num -= 5
			case num >= 4:
				s = append(s, "IV")
				num -= 4
			case num >= 1:
				s = append(s, "I")
				num -= 1
			default:
		}
	}
	return strings.Join(s, "")
}

func main() {
	fmt.Printf("int 3 to %s\n", int2roman(3))
	fmt.Printf("int 4 to %s\n", int2roman(4))
	fmt.Printf("int 9 to %s\n", int2roman(9))
	fmt.Printf("int 58 to %s\n", int2roman(58))
	fmt.Printf("int 1994 to %s\n", int2roman(1994))
}
