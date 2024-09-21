package main

import "fmt"

// rotate ...
// 1 2 3 4 5 6
// 3 2 1 6 5 4
// 4 5 6 1 2 3
func rotate(s []int, k int) {
	k = k % len(s)
	reverse(s[:len(s)-k])
	reverse(s[len(s)-k:])
	reverse(s)
}

// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	rotate(s, 3)
	fmt.Println(s)
}
