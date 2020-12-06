package main

import (
	"fmt"
)

func main() {

	/*go语言的字符串是一个字节的切片
	可以通过将其内容封装在""来其创建字符串，go语言的字符串是unicode兼容的，并且是UTF-8编码的

	字符串是一些字节的集合
	"", ``, ""的字符会转义，但是``表示多行的字符串

	字符操作的都是编码值

	字节: byte ---> uint8
	中文字符串占用3个字节
	*/

	s1 := "hello中国"
	s2 := "hello world!"

	fmt.Println(s1)
	fmt.Println(s2)

	//字符串的长度获取到的是字节的个数
	fmt.Println(len(s1))

	//获取某个字节
	fmt.Println(s1[0])
	a := 'h'
	b := 104
	fmt.Printf("%c %c %c\n", s1[0], a, b)

	//字符串的遍历
	for i := 0; i < len(s2); i++ {
		fmt.Printf("%c", s2[i])
	}
	fmt.Println()

	for _, value := range s2 {
		fmt.Printf("%c", value)
	}
	fmt.Println()

	slice1 := []byte{65,66,67,68,69}
	s3 := slice1
	fmt.Println(string(s3))
	//字符串不允许修改的

	var str1 string = "hello world"
	r := []rune(str1)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%v ", r[i])
	}
	fmt.Println()
}
