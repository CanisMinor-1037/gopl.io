// 编写一个非递归版本的comma函数，使用bytes.Buffer代替字符串链接操作
package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer
	l := len(s)
	i := l % 3

	fmt.Fprintf(&buf, "%s", s[:i])

	for ; i < l; i += 3 {
		if i != 0 {
			buf.WriteByte(',')
		}
		fmt.Fprintf(&buf, "%s", s[i:i+3])
	}
	return buf.String()
}

// test comma
func main() {
	fmt.Println(comma("1234"))
	fmt.Println(comma("1234567"))
	fmt.Println(comma("123456789"))
	fmt.Println(comma("1234567891"))
}
