package main

import "fmt"

func main() {
	label:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 2 {
				break label
			}
			fmt.Println(j)
		}
	}
}