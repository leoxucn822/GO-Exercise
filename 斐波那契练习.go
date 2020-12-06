package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	fmt.Println(fbn(n))
}

func fbn(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fbn(n-1) + fbn(n-2)
	}
}

