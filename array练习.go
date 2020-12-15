package main

import "fmt"

func test01(arr *[3]int) {
	(*arr)[0] = 88
	arr[1] = 99
	fmt.Printf("%p %p\n", arr, &arr)
}

func main() {
	var arr = [3]int{11, 22, 33}
	fmt.Printf("%p\n", &arr)
	test01(&arr)
	fmt.Println(arr)

	var arr1 [3]int = [...]int {1,2,3}
	fmt.Println(arr1)

	var arr2 [3]int = [3]int{2,3,4}
	fmt.Println(arr2)

}