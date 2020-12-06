package main

import "fmt"

func main()  {

	var p *int
	var i int
	p = &i
	fmt.Scanf("%d", p)
	fmt.Printf("The value is %d, the memory address is %p\n", *p, p)
}