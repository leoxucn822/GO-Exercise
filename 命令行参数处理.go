package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("命令行的参数有%v\n", len(os.Args))
	for i, v := range os.Args {
		fmt.Printf("%v %v\n", i, v)
	}
}
