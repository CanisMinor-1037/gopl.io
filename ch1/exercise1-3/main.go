package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// 使用time包来测量执行时间
	start := time.Now()
	inefficientEcho()
	inefficientTime := time.Since(start).Microseconds()

	start = time.Now()
	efficientEcho()
	efficientTime := time.Since(start).Microseconds()

	fmt.Printf("低效版本执行时间: %v\n", inefficientTime)
	fmt.Printf("高效版本执行时间: %v\n", efficientTime)
	fmt.Printf("时间差异: %v\n", inefficientTime-efficientTime)
}

// inefficientEcho 是一个低效的实现版本
func inefficientEcho() {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// efficientEcho 是一个高效的实现版本
func efficientEcho() {
	s := strings.Join(os.Args[1:], " ")
	fmt.Println(s)
}
