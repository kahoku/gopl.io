package main

import (
	"fmt"
	"gopl.io/ch2/scope"
)

//包级别变量，在包内可见
var y int

func main() {
	scope.GetFirst()
	scoperange()
	//fmt.Printf("%q", scope.GetSec())
	//fmt.Println(y)
}

//range创建的作用域
func scoperange() {
	x := "hello"
	y = 100
	//range 创建了隐式作用域
	for _, x := range x {
		x := x + 'A' - 'a'
		fmt.Printf("%c \n", x) // "HELLO" (one letter per iteration)
	}
	
	fmt.Println(x)

}
