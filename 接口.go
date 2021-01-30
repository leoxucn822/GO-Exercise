package main

import (
	"fmt"
	"sort"
)

type sortarray []int

func (test sortarray) Len() int {
	return len(test)
}

func (test sortarray) Less(i, j int) bool {
	return test[i] < test[j]
}

func (test sortarray) Swap(i, j int) {
	temp := test[i]
	test[i] = test[j]
	test[j] = temp
}

func main() {
	var array = sortarray{-98, -1, 5, 10, 32}
	sort.Sort(array)
	fmt.Println(array)
}

