package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	fmt.Println(testdigui(n))
}

func testdigui(n int) int {
	if n == 1{
		return 3
	} else {
		return 2 * testdigui(n-1) + 1
	}
}
