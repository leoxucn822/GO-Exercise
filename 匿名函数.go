//匿名函数就是没有名字的函数
//调用的方式：
//1.	在定义匿名函数的时候直接调用，这种方式匿名函数只能调用一次
//2.	把匿名函数交给一个变量(函数变量)，该变量在对匿名函数进行调用
package main

import "fmt"

func main() {

	result := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)

	fmt.Println(result)
}