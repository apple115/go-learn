package main

import (
	"ch2/popcount"
	"fmt"
)

func main() {
	var a uint64 = 30
	b := popcount.PopCount(a)
	fmt.Println(b)
}
