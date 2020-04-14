// Package main
package main

import "unicode/utf8"

func main() {
	s := "hello, 世界"
	for i := 0; i < len(s); i++ {
		r, size := utf8.DecodeRuneInString(s[i:])
	}
}
