package main

import "fmt"

func main() {
	var line int
	fmt.Scanln(&line)
	for i := 1; i <= line; i++ { 	//i 代表行数
		for space := 1; space <= line - i; space++ {
			fmt.Printf(" ")
		}
		for j := 1; j <= 2 * i - 1; j++ {
			if j == 1 || j == 2 * i - 1 || i == line{
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}