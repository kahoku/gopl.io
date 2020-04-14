package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i: //当i为0时，选择改分支向ch发送0,缓冲区为1
		}
	}
}
