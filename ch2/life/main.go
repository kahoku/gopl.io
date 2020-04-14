// Package main provides ...
package main

import (
	"fmt"
)

func main() {
	t := "a"
	for i := 1; i < 10; i++ {
		t := i * 2
		if i == 9 {
			fmt.Println(t)
		}
	}
	t += "bbb"
	a := 3
	if t == string(a) {
		fmt.Println("equal")
	}
	fmt.Println(t)

}
