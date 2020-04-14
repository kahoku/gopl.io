package main

import (
	"fmt"
	"time"
)

func cc() {
	fmt.Println(222)
}
func run() {
	i := 0
	go cc()
	for {
		fmt.Println(1111)
		i++
	}
}

func main() {
	for i := 0; i <= 20; i++ {
		go run()

	}
	time.Sleep(time.Second)
	for {
		fmt.Println("这句代码会执行吗?")

	}

}
