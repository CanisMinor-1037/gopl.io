// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	//for _, arg := range os.Args[1:] {
	for _, arg := range os.Args[:] { // 修改echo程序，使其能够打印os.Args[0]
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	// 修改echo程序，使其打印每个参数的索引和值，每个一行
	for index, value := range os.Args[:] {
		fmt.Println(index, value)
	}
}
