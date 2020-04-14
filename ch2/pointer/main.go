package main

import (
	"fmt"
)

func main() {
	incrTest()
	//equalTest()
	//retPrtTest()
}

func equalTest() {
	x := 1
	p := &x
	p2 := &x
	fmt.Println(p)
	*p = 2
	fmt.Println(x)
	//指针之间也是可以进行相等测试的，只有当它们指向同一个变量或全部是nil时才相等
	fmt.Println(p == p2)
}

func incr(p *int) int {
	*p++ // 只是增加p指向的变量的值，并不改变p指针！！！
	return *p

}

func incrTest() {
	v := 1
	incr(&v)              // side effect: v is now 2
	fmt.Println(incr(&v)) // "3" (and v is 3)

	x := 1
	p := &x         // p, of type *int, points to x
	fmt.Println(*p) // "1"
	*p = 2          //equivalent to x = 2
	fmt.Println(x)  // "2"

}

func retPrtTest() {
	var p = f()
	*p += 4
	fmt.Println(*p)
}

var p = f()

//返回局部变量的地址
func f() *int {
	v := 1
	return &v

}
