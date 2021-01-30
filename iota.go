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

const (
	n1 = iota //0
	n2 = 100  //100
	n3 = iota //2
	n4        //3
)
const n5 = iota //0

func main() {
/*	var hodType int
	fmt.Printf("%v\n", hodType)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", hodType == b)*/
	//fmt.Printf("%v\n", b)

	fmt.Printf("%v\n", KB)
	fmt.Printf("%v\n", MB)
	fmt.Printf("%v\n", GB)

	fmt.Printf("%v\n", n1)
	fmt.Printf("%v\n", n2)
	fmt.Printf("%v\n", n3)
	fmt.Printf("%v\n", n4)
	fmt.Printf("%v\n", n5)
}
