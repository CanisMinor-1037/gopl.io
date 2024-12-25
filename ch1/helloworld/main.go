// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 1.

// Helloworld is our first Go program.
//!+
// go语言的代码通过package组织
// 一个package由位于单个目录下的一个或多个.go文件组成
// 每个源文件都是以一条package声明语句开始, 表示文件所属的包
// main包定义了一个独立可执行的程序, 而不是一个库
package main

// 包声明语句之后, 是import语句, 用于导入程序中需要使用的包
// Go的标准库提供了大量的包, 用于处理常见的任务, 如输入输出、排序以及文本处理等
// fmt包提供了格式化输出、接收输入的函数
// 必须导入恰当的包，缺少必要的包或者导入未使用的包，都会导致编译错误
import "fmt"

// 组成程序的函数、变量、常量、类型的声明语句
// main包的main函数是程序执行时的入口
// main函数内调用fmt包的Println函数, 打印"Hello, 世界"
func main() {
	// Println是其中一个基础函数，可以打印以空格间隔的一个或多个值，并在最后添加一个换行符，从而输出一整行。
	fmt.Println("Hello, 世界") // 不需要在语句的末尾添加分号, 除非一行上有多条语句
}

//!-
