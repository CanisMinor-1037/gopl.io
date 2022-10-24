package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%gF = %gC\n", freezingF, ftoc(freezingF))
	fmt.Printf("%gF = %gC\n", boilingF, ftoc(boilingF))
}

func ftoc(f float32) (c float32) {
	c = (f - 32) * 5 / 9
	return
}
