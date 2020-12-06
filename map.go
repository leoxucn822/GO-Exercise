package main

import (
	"fmt"
	"sort"
)

func main() {

	var map1 map[int]string
	var map2 = make(map[int]string)
	var map3 = map[int]string{1:"go", 2:"C", 3:"C++", 4:"Java"}

	//map1[0] = "hello" // if map is nil, can't assignment
	fmt.Printf("%v\n", map1)
	fmt.Printf("%v\n", map2)
	fmt.Printf("%v\n", map3)

	fmt.Println(map1 == nil)
	fmt.Printf("%v\n", map3[40]) //print the none

	v,k := map3[40]

	if k {
		fmt.Printf("对应的数值是: %v\n", v)
	}else {
		fmt.Printf("操作的值不存在，获取的值是零值 %v\n", v)
	}

	var keys []int
	for k, _ := range map3 {
		keys = append(keys, k)
	}
	fmt.Println(keys)	//打印出来是无序的
	sort.Ints(keys)
	fmt.Println(keys)
}
