package main

import "fmt"

func main()  {
	a := "hello world"
	b := [] byte(a)
	c := [] rune(a)
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
}
