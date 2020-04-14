package main

import (
	"bytes"
	"fmt"
	"io"
)

const debug = true

func main() {
	var buf *bytes.Buffer //*bytes.Buffer的空指针
	//var buf io.Writer
	if debug {
		buf = new(bytes.Buffer) //如果开启调试，初始化输出bufer
	}
	f2(buf) // NOTE: subtly incorrect!
	if debug {
		//使用buf
		fmt.Printf("%v", buf)

	}

}

// 小心out在运行态 为空指针值的非空接口
func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("done!\n"))
	}

}

func f2(out io.Writer) {
	fmt.Printf("f2 out is: %#v, T:%T \n", out, out)

	//类型断言做层转换,out2从非空接口值变为*bytes.Buffer的零值
	if out2, ok := out.(*bytes.Buffer); ok {
		fmt.Printf("out2:%#v T:%T ,ok:%v\n", out2, out2, ok)
		if out2 != nil {
			out2.Write([]byte("done!\n"))
		} else {
			fmt.Printf("out2 is nil: %#v", out2)
		}
	} else {
		fmt.Printf("out2:%#v ok:%v\n", out2, ok)
	}

}
