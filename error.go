package main

import "fmt"

func divide(num1 int, num2 int) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println(num1 / num2)
}

func main() {
	var a int
	var b int
	fmt.Scanln(&a, &b)
	divide(a,b)
	fmt.Println("main下面的代码")
}
