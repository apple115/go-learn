package main

import "fmt"

// remove ...
func removeRepeatWords(s []string) []string {
	w := 0 //从第一个元素开始
	for _, v := range s {
		if w == 0 || s[w-1] != v { //检查当前元素 v 是否与前一个元素不同
			s[w] = v
			w++
		}
	}
	return s[:w]
}

func main() {
	s := []string{"a", "a", "a", "b", "b", "c", "c", "d", "e", "f", "f", "g"}
	//
	fmt.Println(removeRepeatWords(s))

}
