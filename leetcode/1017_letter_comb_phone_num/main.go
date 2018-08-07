package main

import (
    "fmt"
)

func letterCombinations(digits string) []string {
    d := map[string]string{
        "2": "abc",
        "3": "def",
        "4": "ghi",
        "5": "jkl",
        "6": "mno",
        "7": "pqrs",
        "8": "tuv",
        "9": "wxyz",
        "0": "",
        "1": "",
    }
    if len(digits) == 0 {
        return []string{}
    }
    res := []string{""};
    for _, sd := range digits {
        sd := d[string(sd)]
        if len(sd) <= 0 {
            continue
        }
        n := len(res)
        for ri :=0; ri < n; ri++ {
            tmp := string(res[0])
            //fmt.Println(tmp)
            res = res[1:]
            for k :=0; k < len(sd); k++ {
                res = append(res, tmp + string(sd[k]))
            }
        }
    }
    return res
}


func main() {
    s := "2340"
    fmt.Println(letterCombinations(s))
}
