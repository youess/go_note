package main

import "fmt"
//import "reflect"

const BitsPerWord = 32 << (^uint(0) >> 63)  // either 32 or 64

const (
    MaxInt  = 1<<(BitsPerWord-1) - 1 // either 1<<31 - 1 or 1<<63 - 1
    MinInt  = -MaxInt - 1            // either -1 << 31 or -1 << 63
    MaxUint = 1<<BitsPerWord - 1     // either 1<<32 - 1 or 1<<64 - 1
)

func main() {

    fmt.Printf("uint(0): %d\n", uint(0))
    fmt.Printf("^uint(0): %d\n", ^uint(0))
    fmt.Printf("bits per word : %d\n", BitsPerWord)
    fmt.Printf("max int: %d\n", MaxInt)
    fmt.Printf("min int: %d\n", MinInt)
    fmt.Printf("max uint: %d\n", uint64(MaxUint))
    /*
    uint(0): 0
    ^uint(0): 18446744073709551615
    bits per word : 64
    max int: 9223372036854775807
    min int: -9223372036854775808
    max uint: 18446744073709551615
    */

}
