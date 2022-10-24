package main

import (
	"fmt"

	"gopl.io/ch2/popcount"
)

func main() {
	var n uint64
	fmt.Scanf("%v", &n)
	fmt.Printf("PopCount: %v\n", popcount.PopCount(n))
	fmt.Printf("PopCount1: %v\n", popcount.PopCount1(n))
	fmt.Printf("PopCount2: %v\n", popcount.PopCount2(n))
	fmt.Printf("PopCount3: %v\n", popcount.PopCount3(n))

	// TODO 测试性能
}
