package main

import (
	"errors"
	"fmt"
)

//使用error.New或者panic的内置函数
//error.New("错误说明")会返回一个error类型的值，表示一个错误
//panic内置函数，接受一个interface{}类型的值(也就是任何值了)作为参数。可以接受error类型的变量，输出错误信息，并且退出程序

func readconf(name string) (err error) {
	if name == "init.config" {
		return nil
	} else {
		return errors.New("can't find the file")
	}
}

func test_readconf(name string) {
	err := readconf(name)
	if err != nil {
		panic(err)
	}
	fmt.Println("file be found")
}

func main() {
	//test_readconf("hello world")
	test_readconf("init.config")
	fmt.Println("main function go on!")
}
