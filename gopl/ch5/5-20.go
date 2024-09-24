package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	ch <- 1 //错误 死锁
	// close(ch) //错误，关了，读不出来了
	// go func() { ch <- 1 }() //正确，必须单独在一个goroutine中使用
	// var a int
	<-ch
	 close(ch) //错误，关了，读不出来了
	// fmt.Println(<-ch)
	fmt.Println("End")
}
