package main

// for the same type of variable, we can ignore the fist one, just keep the last one is ok
//func add(a int, b int) int {
/*func add(a,b int) int {
	return a + b
}

func main() {
	fmt.Println("the sum of a and b is", add(3,5))
}*/

/*func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}*/

/*func sum(x,y int) (z int) {
	z = x + y
	return
}*/

/*func sum(x,y int) int{
	z := x + y
	return z
}

func main()  {
	fmt.Println("The sum of two number is", sum(2,5))
}*/

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for math.Pow(z, z) - x > 1e-10 {
		z -= (z * z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}



