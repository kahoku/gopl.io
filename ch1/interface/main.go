package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

//接口介绍
/**

 */

//定义接口
type Awesomizer interface {
	Awesomize() string
}

//类型不声明实现了某个接口
type Foo struct{}

//通过实现接口中的方法，隐式实现了接口
func (foo Foo) Awesomize() string {
	return "Awesome!"
}

/*
//实现了io.Writer接口
func (foo Foo) Write(data []byte) (n int, err error) {
	fmt.Printf("foo.write %v", data)
	return len(data), nil
}
*/

func main() {

	//判断实现接口的朴素方式
	var _ Awesomizer = (*Foo)(nil)
	//fmt.Printf("k:%#v", k)

	//f := &Foo{}
	//f.Write([]byte{1, 2, 3})
	//condition()

}

//实现接口条件
func condition() {
	var w io.Writer       //声明一个nil接口
	w = os.Stdout         // OK: *os.File has Write method
	w = new(bytes.Buffer) // OK: *bytes.Buffer has Write method
	//w = time.Second       // compile error: time.Duration lacks Write method

	fmt.Printf("w:%#v \n", w)
	var rwc io.ReadWriteCloser
	rwc = os.Stdout // OK: *os.File has Read, Write, Close methods
	//	rwc = new(bytes.Buffer) // compile error: *bytes.Buffer lacks Close method
	fmt.Printf("w:%#v \n", rwc)

}
