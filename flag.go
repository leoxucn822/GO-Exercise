package main

import (
	"flag"
	"fmt"
)

func main() {
	var user string
	var pwd string
	var host string
	var port int

	// &user  接受用户命令行输入-u 后面参数的值
	// -u 指定参数
	// "" 默认值
	flag.StringVar(&user, "u", "", "用户名默认为空")
	flag.StringVar(&pwd, "p", "", "密码默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名默认为localhost")
	flag.IntVar(&port, "P", 80, "端口默认为80")

	//这里有个非常重要的操作,转换，必须要调用
	flag.Parse()
	fmt.Printf("user=%v, pwd=%v, host=%v, port=%v\n", user, pwd, host, port)

}
