package main

import (
	"fmt"

	"gopl.io/ch2/scope"
)

//包级别变量，在包内可见
var y int

func main() {

	scoperange()
	fmt.Printf("%q", scope.GetSec())
	fmt.Println(y)
}

//range创建的作用域
func scoperange() {

	x := "hello"
	y = 100
	for _, x := range x {
		x := x + 'A' - 'a'
		fmt.Printf("%c", x) // "HELLO" (one letter per iteration)

	}
}
