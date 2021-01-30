package main

import "fmt"

func main() {
	var arr[2][3]int

	arr[1][1] = 5
	fmt.Printf("%p %p\n", &arr[0], &arr[0][0])
	fmt.Printf("%p %p\n", &arr[1], &arr[1][0])
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%v ", arr[i][j])
		}
		fmt.Println()
	}

	var arr1 = [...][5]int{{123},{3,4,5}}
	fmt.Println(arr1)
}
