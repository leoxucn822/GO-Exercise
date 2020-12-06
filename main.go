package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("pls input the string: ")
	reader := bufio.NewReader(os.Stdin)
	s1,_ := reader.ReadString('\n')
	fmt.Println(s1)
}
