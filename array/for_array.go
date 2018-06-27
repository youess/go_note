package main

import "fmt"

func main() {

	numA := []int{10, 20, 30, 40}

	for idx, e := range numA {
		fmt.Printf("idx: %d, val: %d\n", idx, e)
	}

	numB := [][2]int{{1, 2}, {3, 4}, {4, 5}}
	fmt.Println("")
	for idx, e := range numB {
		fmt.Printf("idx: %d, val: %d\n", idx, e)
	}
}
