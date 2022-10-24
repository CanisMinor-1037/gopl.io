package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var calSha256 = flag.Bool("sha256", true, "calculate sha256")
var calSha384 = flag.Bool("sha384", false, "calculate sha384")
var calSha512 = flag.Bool("sha512", false, "calculate sha512")
var input = flag.String("input", "", "data to be processed")

func main() {
	flag.Parse()

	if *calSha256 {
		fmt.Printf("sha256: %x\n", sha256.Sum256([]byte(*input)))
	}
	if *calSha384 {
		fmt.Printf("sha384: %x\n", sha512.Sum384([]byte(*input)))
	}
	if *calSha512 {
		fmt.Printf("sha512: %x\n", sha512.Sum512([]byte(*input)))
	}
}
