package main

import "fmt"

const (
	//a = iota
	_ = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
	//b
	//c
)

/*const (
	m = iota
)*/

func main() {
/*	var hodType int
	fmt.Printf("%v\n", hodType)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", hodType == b)*/
	//fmt.Printf("%v\n", b)

	fmt.Printf("%v\n", KB)
	fmt.Printf("%v\n", MB)
	fmt.Printf("%v\n", GB)
}
