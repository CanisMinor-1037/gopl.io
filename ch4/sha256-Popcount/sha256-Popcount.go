package main

import (
	"crypto/sha256"
	"fmt"
)

func PopCount(x byte) int {
	var cnt int
	for x != 0 {
		cnt++
		x = x & (x - 1)
	}
	return cnt
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	var cnt1, cnt2 int

	for _, v := range c1 {
		cnt1 += PopCount(v)
	}

	for _, v := range c2 {
		cnt2 += PopCount(v)
	}

	fmt.Printf("%x\n%d\n%x\n%d\n", c1, cnt1, c2, cnt2)
}
