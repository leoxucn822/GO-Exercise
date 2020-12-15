package main

import "fmt"

func main() {

	//string 底层其实是一个byte数组，所以可以通过切片进行处理
	str := "中国你好"

	//使用切片来获取中文字符
	slice1 := str[6:]
	fmt.Println(len(str))
	fmt.Println(slice1)

	var s1 []rune = []rune(str)
	var s2 []byte = []byte(str)

	fmt.Println(s1)
	fmt.Println(s2)
}
