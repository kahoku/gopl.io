//切片介绍
/*
	var a []int                               // 去掉了长度
	var a = []int{1, 2, 3, 4}                 // 声明初始化
	a := []int{1, 2, 3, 4}                    // 简短声明
	chars := []string{0: "a", 2: "c", 1: "b"} // 带索引声明

	//使用切片创建切片，共享同一数组
	var b = a[1:4]       // slice from index 1 to 3
	var b = a[:3]        // missing low index implies 0
	var b = a[3:]        // missing high index implies len(a)
	a = append(a, 17, 3) // append items to slice a
	c := append(a, b...) // concatenate slices a and b

	// 通过make生成
	a = make([]byte, 5, 5) // 第一个是长度，第二个为容量
	a = make([]byte, 5)    // 容量可以省略

	// 通过数组生成
	x := [3]string{"苹果", "香蕉", "鸭梨"}
	s := x[:] // s为一个slice并且 引用了x的存储内容

*/
package main

import "fmt"

func main() {
	//shareCnt()
	assignCap()
}

//共享数据
func shareCnt() {
	// 创建一个整型切片
	// 其长度和容量都是 5 个元素
	slice := []int{10, 20, 30, 40, 50}
	// 创建一个新切片
	// 其长度为 2 个元素，容量为 4 个元素
	newSlice := slice[1:3]

	// 修改 newSlice 索引为 1 的元素
	newSlice[1] = 35

	// 同时也修改了原来的 slice 的索引为 2 的元素
	fmt.Printf("share slice[1]:%v", slice[2])

	// 使用原有的容量来分配一个新元素
	// 将新元素赋值为 60
	newSlice = append(newSlice, 60)

}

//指定容量
func assignCap() {
	// 创建字符串切片
	// 其长度和容量都是 5 个元素
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	// 对第三个元素做切片，并限制容量
	// 其长度和容量都是 1 个元素
	slice := source[2:3:3]
	// 向 slice 追加新字符串
	slice = append(slice, "Kiwi")

	//不会改变原来的数据
	fmt.Printf("assignCap source:%v \n", source)
	fmt.Printf("assignCap slice:%v", slice)

}

//for-range

func forRangeSlice() {
	// 创建一个整型切片
	// 其长度和容量都是 4 个元素
	slice := []int{10, 20, 30, 40}
	// 迭代每个元素，并显示值和地址,值地址最终是一样的，
	//range 创建了每个元素的副本
	for index, value := range slice {
		fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n",
			value, &value, &slice[index])

	}

}
