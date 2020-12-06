package main

import "fmt"

func Addplus() func (int) int {
	var n int = 10
	return func (x int) int {
		n = n + x
		return n
	}
}

func main() {
	f := Addplus()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
}
