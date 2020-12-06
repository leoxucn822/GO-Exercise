package main

import "fmt"

func main() {
	var x int
	var y string
	var z bool

	var a int = 42
	var b string = "James Bond"
	var c bool = true

	fmt.Println(x,y,z)
	s := fmt.Sprintf("%v %v %v", a,b,c)
	fmt.Printf("%T %v", s, s)
}
