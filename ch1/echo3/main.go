// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
)

// !+
func main() {
	// 使用strings包的Join函数, 将os.Args[1:]中的元素连接成一个字符串, 并用空格分隔
	fmt.Println(strings.Join(os.Args[1:], " "))
}

//!-
