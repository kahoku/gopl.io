// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 228.

// Pipeline1 demonstrates an infinite 3-stage pipeline.
package main

import "fmt"
import "os"

//!+没有关闭通知会一直执行计算
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; ; x++ {
			if(x > 100) {
				close(naturals)
				os.Exit(4)
				break
			}else{
				naturals <- x
			}
		}
	}()

	// Squarer
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	// Printer (in main goroutine)
	for {
		fmt.Println(<-squares)
	}
}

//!-
