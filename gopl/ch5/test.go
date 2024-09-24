package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup


	// 模拟 worker goroutine 的循环
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			ch <- val
		}(i)
	}

	// 等待所有 worker goroutine 完成
	go func() {
		wg.Wait()
		close(ch)
	}()

	// 模拟 sizes 通道的循环
	for v := range ch {
		fmt.Println(v)
	}
}
