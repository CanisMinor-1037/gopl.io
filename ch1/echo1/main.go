// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 4.
//!+

// Echo1 prints its command-line arguments.
package main

// os包以跨平台的方式, 提供了一些与操作系统交互的函数和变量。
// 程序的命令行参数可通过os包的Args变量获取
// os.Args变量是一个字符串切片
// 切片是Go语言中的一种数据结构, 用于表示一个可变长度的序列
// 切片可以看作是数组的一个动态视图, 它提供了更灵活的元素访问方式
// 切片由三个部分组成: 指针、长度和容量
// 指针指向底层数组中的第一个元素, 长度表示切片中元素的个数, 容量表示切片底层数组的总长度
import (
	"fmt"
	"os"
)

func main() {
	// var声明两个string类型的变量s和sep, 并隐式初始化为空字符串""
	var s, sep string
	// os.Args[0]是程序的名称, 而os.Args[1:]是程序的参数
	// i是循环索引变量，在for循环的初始化语句中定义
	// :=是Go语言的短变量声明运算符, 用于声明并初始化变量, 赋予适当的类型
	// go语言只有for循环一种循环结构, 包含初始化语句、条件表达式、后置语句三个部分，都可选
	for i := 1; i < len(os.Args); i++ {
		// 赋值语句，将s的旧值与sep和os.Args[i]连接, 并赋值给s
		// +=是Go语言的复合赋值运算符, 用于将右边的表达式与左边的变量相加, 并将结果赋值给左边的变量
		// 当参数数量庞大时, 使用+=连接字符串的开销会很大
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

//!-
