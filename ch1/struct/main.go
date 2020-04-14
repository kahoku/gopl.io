//结构体介绍
/*
// 以下创建变量的几种方式
var v = Vertex{1, 2} //值结构体变量
var v = Vertex{X: 1, Y: 2}  //指定key创建
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
	"encoding/json"
)

type Vertex struct {
	X, Y float64
}


var a Vertex

//声明方式枚举
func declarCase() {
	e := new(Vertex)  // 同 var a *Vertext = new(Vertex)
	c := Vertex{1, 2} //值结构体变量
	d := Vertex{X: 1, Y: 2} //指定key创建

    //b := Vertex // type Vertex is not an expression
	b := Vertex{} // 和 a 一样

	//f := &Vertex// type Vertex is not an expression
	f := &Vertex{}

	g := &Vertex{3,4}//同 var g *Vertex = &Vertex{3,4}
	h := []Vertex{{1,2},{5,2},{5,5}} // 初始化元素是结构体的slice

	fmt.Println(a,b,c,d,e,f,g,h)
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
	//methodCall()
	//inMapCase()
   //declarCase()
   //useTag()

   conversType()
}

//方法介绍
func methodCall() {
	vv := &Vertex{1, 2}
	v := Vertex{3, 4}
	//访问结构体成员
	vv.X = 4
	fmt.Printf("vv.X has changed : %v \n", vv)

	// 方法调用
	abs := vv.Abs()
	fmt.Printf("vv.Abs=%9.2f \n", abs)

	//在值变量上调用接受者指针方式,go会转指针
	v.add(10)
	fmt.Printf("v has added %v \n", v)
}

type Test struct {
	X,Y float64
}

func conversType() {
	var a *Vertex
	t := &Test{1,2}

	a = (*Vertex)(t)
	fmt.Println(a)

}
//与map结合
func inMapCase() {
	type user struct{ name string  }
	/*
	当 map 因扩张而重新哈希时，各键值项存储位置都会发生改变。
	因此，map 被设计成 not addressable。
	类似 m[1].name 这种期望透过原 value 指针修改成员的行为自然会被禁 。
	*/
	m := map[int]user{ //
		1: {"Tom"},
	}

	fmt.Println(m[1].name)
	// m[1].name = "Tom"
	// ./main.go:16:12: cannot assign to struct field m[1].name in map
	fmt.Println(m)

	// 正确做法是完整替换 value 或使用指针。
	u := m[1]
	u.name = "Tom"
	m[1] = u // 替换 value。

	m2 := map[int]*user{
		1: &user{"user1"},
	}

	m2[1].name = "Jack" // 返回的是指针复制品。透过指针修改原对象是允许的。
	fmt.Printf("m2:%#v \n", m2)
}

//json与struct
func useTag() {
	type User struct {
		UserId   int
		UserName string

	}

	//标签是类型的组成部分
	type UserTag struct {
		UserId   int    `json:"user_id" bson:"user_id"`
		UserName string `json:"user_name" bson:"user_name"`

	}

	u := &User{UserId: 1, UserName: "Murphy"}
	j, _ := json.Marshal(u)
	fmt.Printf("struct User echo : %v\n", string(j))

	u_tag := &UserTag{UserId: 1, UserName: "Murphy"}
	j_tag, _ := json.Marshal(u_tag)
	fmt.Printf("struct UserTag echo : %v\n", string(j_tag))

}

