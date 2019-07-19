package main

import (
	"fmt"
	// 导入字符串工具包
	"github.com/favorstack/go-example/stringutil"
)

func main() {
	fmt.Println("Hello, World!")

	s := "悠云白雁过南楼"
	fmt.Println(s)

	// 使用反转函数
	fmt.Println(stringutil.Reverse(s))
}
