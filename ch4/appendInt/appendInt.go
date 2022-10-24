package main

import "fmt"

func main() {
	/*var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)*/
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%2d %v\n", i, cap(y), y)
		x = y
	}
	var z []int
	z = append(z, 1)
	z = append(z, 2, 3)
	z = append(z, 4, 5, 6)
	z = append(z, z...)
	fmt.Println(z)
}

// 专门用于处理[]int类型的slice
func appendInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		// There is room to grow. Extend the slice
		z = x[:zlen]
	} else {
		// There is insufficient space. Allocate a new array
		// Grow by doubling, for amortized linear complexity
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function;
	}
	copy(z[len(x):], y)
	return z
}
