package main

import "fmt"

var globalVar int = 10

func main()  {
	globalVar = 20
	fmt.Printf("%d\n", globalVar)
}

