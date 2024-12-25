// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	// 短变量声明语句可以用在函数内部，不能用于包变量
	s, sep := "", ""
	// for循环的range格式可以遍历数组、切片、字符串、map和channel
	// 在range格式中, 第一个返回值是元素的索引, 第二个返回值是元素的值
	// 使用_来忽略索引, _是Go语言的空白标识符, 用于表示一个未命名的变量
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

//!-
