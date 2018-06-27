package main

import (
	"fmt"
	"strings"
)

func main() {

	sc := make([]string, 0)
	sc = append(sc, "abcd") // append mean append element at last
	sc = append(sc, "efgh")
	fmt.Printf("%s\n", strings.Join(sc, "_")) // abcd_efgh

	sc = make([]string, 2) // not use slice len with append
	sc = append(sc, "abcd")
	sc = append(sc, "efgh")
	fmt.Printf("%s\n", strings.Join(sc, "_")) // __abcd_efgh

	sc = make([]string, 2)
	sc[0] = "abcd"
	sc[1] = "efgh"
	fmt.Printf("%s\n", strings.Join(sc, "_")) // abcd_efgh

}
