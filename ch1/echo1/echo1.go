// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	//for i := 1; i < len(os.Args); i++ {
	for i := 0; i < len(os.Args); i++ { // 修改echo程序，使其能够打印os.Args[0]
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
