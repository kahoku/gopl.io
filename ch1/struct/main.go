//结构体介绍
/*
// 以下创建变量的几种方式
var v = Vertex{1, 2}//值结构体变量
var v = Vertex{X: 1, Y: 2} //指定key创建
var v = []Vertex{{1,2},{5,2},{5,5}} // 初始化元素是结构体的slice
vv := &Vertex {1, 2}//指针结构体变量，等价于 vv:= new(Vertex); *vv=Vertex{1,2}
//匿名结构体
point := struct {
X, Y int

}{1, 2}

*/
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//结构体上可以定义接受者值类型的方法.
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//定义指针接受者类型的方法
func (v *Vertex) add(n float64) {
	v.X += n
	v.Y += n
}

func main() {
	methodCall()

}

func methodCall() {
	vv := &Vertex{1, 2}
	v := Vertex{3, 4}
	//访问结构体成员
	vv.X = 4
	fmt.Printf("vv.X has changed : %v \n", vv)

	// 方法调用
	abs := vv.Abs()
	fmt.Printf("vv.Abs=%9.2f \n", abs)

	//在值变量上调用接受者指针方式,go会自动转指针
	v.add(10)
	fmt.Printf("v has added %v \n", v)
}
