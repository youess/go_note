package main

import "fmt"

func test1(s []int) {
    s = append(s, 123)
    fmt.Printf("in test1 slice: %v\n", s)
}

func test2(s *[]int) {
    *s = append(*s, 123)
    fmt.Printf("in test2 slice: %v\n", s)
}

func main() {
    a := make([]int, 0)
    a = append(a, []int{3, 4, 5}...)
    fmt.Printf("start slice: %v\n", a)
    test1(a)
    fmt.Printf("after test1 slice in main: %v\n", a)  // didn't change
    test2(&a)
    fmt.Printf("after test2 slice in main: %v\n", a)  // changed
}
