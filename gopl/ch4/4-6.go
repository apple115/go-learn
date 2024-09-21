package main

import "fmt"

// removeRepeatSpace ...
// a b  c
func removeRepeatSpace(b []byte) []byte {
	w := 0
	for _, v := range b {
		if v != ' ' || b[w-1] != ' ' {
			// 将当前字符赋值给新位置 w
			b[w] = v
			// 增加 w 的值，为下一个字符腾出空间
			w++
		}
	}
	return b[:w]
}

func main() {
	a := []byte("aa b  c")
	fmt.Println(removeRepeatSpace(a))
}
