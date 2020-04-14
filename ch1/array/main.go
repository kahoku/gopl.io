//数组介绍
package main

import "fmt"

func main() {

	//showDeclar()
	showPtrCnt()
}

//如何声明
func showDeclar() {
	var a [10]int // 声明包含10个整形元素的数据
	a[3] = 42     // 赋值42
	i := a[3]     // 读取第三个元素
	fmt.Printf("i:%v \n", i)

	// 声明并初始化
	var b = [2]int{1, 2}
	//	b := [2]int{1, 2} //简短声明
	//	b := [...]int{1, 2} //长度自动计算

	months := [...]string{1: "January" /* ... */, 12: "December"} // 0默认空

	fmt.Printf("a:%v;b:%v;month:%v", a, b, months)
}

//相同底层内容
func showPtrCnt() {

	// 声明第一个包含 3 个元素的指向字符串的指针数组
	var array1 [3]*string
	// 声明第二个包含 3 个元素的指向字符串的指针数组
	// 使用字符串指针初始化这个数组
	array2 := [3]*string{new(string), new(string), new(string)}
	// 使用颜色为每个元素赋值
	*array2[0] = "Red"
	*array2[1] = "Blue"
	*array2[2] = "Green"

	// 将 array2 复制给 array1,他们存储的地址一样
	array1 = array2

	*array1[0] = "Redold"

	fmt.Printf("array1:%v\n", array1)
	fmt.Printf("array2:%v;array2[0]:%v", array2, *array2[0])

}
