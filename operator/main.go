package main

import (
	"fmt"
	"math"
)

func main() {

	a := 15
	fmt.Printf("(a):%d, (a/=10):%d, (a%%10):%d", a, a/10, a%10) // (a):15, (a/=10):1, (a%10):5

	a++
	fmt.Printf("\n(a++):%d", a) // 16

	a += 4
	fmt.Printf("\n(a+=3):%d", a) // 20

	a *= 2
	fmt.Printf("\n(a*=2):%d", a) // 40

	a /= 10
	fmt.Printf("\n(a*=2):%d", a) // 4

	fmt.Printf("\n(math.Pow(2, float64(a))):%.3f", math.Pow(2, float64(a)))

}
