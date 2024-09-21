package main

import "fmt"

// reverse ...
func reverse(sp *[5]int) {
	for i, j := 0, len(*sp)-1; i < j; i, j = i+1, j-1 {
		(*sp)[i], (*sp)[j] = (*sp)[j], (*sp)[i]
	}
}



func main() {
	s := [5]int{1, 2, 3, 4, 5}
	a := &s
	reverse(a)
	fmt.Println(s)
}
