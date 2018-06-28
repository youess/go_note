package main

import "fmt"

func main() {
	a, b := 0, 0
	for a < 10 {
		b += a
		a++
	}
	fmt.Printf("Final sum result is: %d", b) // 45
}
