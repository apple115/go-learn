package main

import "fmt"

func testsaa() (result int) {
	defer func() {
		if p := recover(); p != nil {
			result = 32
		}
	}()
	panic("panic")
}

func main() {
	a := testsaa()
	fmt.Println(a)
}
