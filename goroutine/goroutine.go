package main

import (
	"time"
	"fmt"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(index int) {
			for{
				a[index]++
				fmt.Printf("Hello from goroutine %d\n", index)
				// 交出控制权
				//runtime.Gosched()
			}
		}(i)
	}

	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
