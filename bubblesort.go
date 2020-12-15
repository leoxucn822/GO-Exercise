package main

import "fmt"

func bubblesort(arr *[]int, length int)  {
	var temp int
	for j := 0; j < length; j++ {
		for i := 0; i < length-1-j; i++ {
			if (*arr)[i] > (*arr)[i + 1] {
				temp = (*arr)[i]
				(*arr)[i] = (*arr)[i + 1]
				(*arr)[i + 1] = temp
			}
		}
	}
	fmt.Println(*arr)
}

func main() {
	var length int
	fmt.Scanln(&length)
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	bubblesort(&arr, length)
}