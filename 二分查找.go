package main

import "fmt"

func binaryfind(arr *[]int, leftindex int, rightindex int, num int) {

	if leftindex > rightindex {
		fmt.Println("对不起没有要查询的数字")
	}

	middle := (leftindex+rightindex) / 2
	if num < (*arr)[middle] {
		binaryfind(arr, leftindex, middle-1, num)
	} else if num > (*arr)[middle]{
		binaryfind(arr, middle+1, rightindex, num)
	} else {
		fmt.Println("找到了")
	}
}

func sort(arr *[]int, length int)  {
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
}

func main() {
	var length int
	var number int
	fmt.Println("请输入数组的长度")
	fmt.Scanln(&length)
	fmt.Println("请输入数组元素")
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	fmt.Println("请输入要查询的数")
	fmt.Scanln(&number)
	sort(&arr, length)
	fmt.Println(arr)
	binaryfind(&arr, 0, length-1, number)
}
