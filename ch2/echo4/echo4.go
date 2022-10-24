// Echo4 prints its command-line arguments
package main

import (
	"flag"
	"fmt"
	"strings"
)

/**
 * @brief flag.Bool创建一个新的对应布尔型标志参数的变量
 * @param name 命令行参数的名字
 * @param value 该参数的默认值
 * @param usage 该标志参数的描述信息
 * @ret 指向对应命令行标志参数变量的指针
 */
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	//! 程序运行时，必须在使用标志参数对应的变量前调用flag.Parse()
	//! 用于更新每个标志参数对应变量的值
	flag.Parse()

	// 调用flag.Args()访问非标志参数的命令行参数
	// 返回值对应一个字符串类型的slice
	fmt.Print(strings.Join(flag.Args(), *sep))

	if !*n {
		fmt.Println()
	}

}
