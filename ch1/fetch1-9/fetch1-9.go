// Fetch prints the content found at a URL
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url) // 创建HTTP请求, resp得到访问的请求结果

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body) // resp.Body包括一个可读的服务器响应流
		// ioutil.ReadAll()从response中读取到全部内容，将其结果保存到变量b中
		resp.Body.Close() // 关闭resp的Body流，防止资源泄露
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(0)
		}
		fmt.Printf("%s", b) // 将结果b写入到标准输出流中

		// TODO: 打印HTTP协议的状态码
		fmt.Printf("\n%v\n", resp.Status)
	}
}
