//如果一个类型实现了String()这个方法，那么fmt.Println默认会调用这个变量的String()方法进行输出
package main

import "fmt"

type MethodUtils struct {
	length int
	width int
}

func (utils *MethodUtils) Print(length int, width int) {
	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

func main() {
	var utils MethodUtils
	fmt.Scan(&utils.length, &utils.width)
	utils.Print(utils.length, utils.width)
}
