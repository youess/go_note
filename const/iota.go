package main

import "fmt"

type ByteSize float64

const (
	_           = iota             // 忽略第一个值,通常iota用来表示递增，第一个值为0，且最好每一行一个声明变量
	KB ByteSize = 1 << (10 * iota) // 1 << (10*1)
	MB                             // 1 << (10*2)
	GB                             // 1 << (10*3)
	TB                             // 1 << (10*4)
	PB                             // 1 << (10*5)
	EB                             // 1 << (10*6)
	ZB                             // 1 << (10*7)
	YB                             // 1 << (10*8)
)

func test_float(f float64) {
	fmt.Printf("this is a function receive float number, %f", f)
}

func test_bytesize(b ByteSize) {
	fmt.Printf("this is a function receive bytesize variable, %f", b)
}

func main() {

	fmt.Printf("KB = %f\n", KB)
	fmt.Printf("MB = %f\n", MB)
	fmt.Printf("GB = %f\n", GB)
	fmt.Printf("TB = %f\n", TB)
	fmt.Printf("PB = %f\n", PB)
	fmt.Printf("EB = %f\n", EB)
	fmt.Printf("ZB = %f\n", ZB)
	fmt.Printf("YB = %f\n", YB)

	fmt.Println()
	test_bytesize(400.0) // 常量转换

	/* 类型检查
	// .\iota.go:41:15: cannot use f (type float64) as type ByteSize in argument to test_bytesize
	f := 1000.0
	test_bytesize(f)
	*/

	// test_float(KB) // .\iota.go:39:12: cannot use KB (type ByteSize) as type float64 in argument to test_float

}
