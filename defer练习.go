package main

import "fmt"

func sum(num1 int, num2 int) int {
	defer fmt.Println("the first number: ", num1)
	defer fmt.Println("the second number: ", num2)
	num1++
	num2++
	res := num1 + num2
	fmt.Println("the sum of the numbers ", res)
	fmt.Println("the new number of num1 ", num1)
	fmt.Println("the new number of num2 ", num2)
	return res

}

func main() {
	res := sum(10, 20)
	fmt.Println(res)
}
