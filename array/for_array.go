package main

import "fmt"

func main() {

	nums := []int{10, 20, 30, 40}

	for idx, e := range nums {
		fmt.Printf("idx: %d, val: %d\n", idx, e)
	}
}
