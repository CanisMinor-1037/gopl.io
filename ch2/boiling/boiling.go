// boiling.go打印水的沸点
package main

import "fmt"

const bF = 212.0

func main() {
	var boilingF float32 = bF
	var boilingC float32 = (boilingF - 32) * 5 / 9
	fmt.Printf("%gF, %gC\n", boilingF, boilingC)
}
