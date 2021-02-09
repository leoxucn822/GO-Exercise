package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var strcount int
	var spacecount int
	var number int
	var otherchar int
	file, err := os.Open("/tmp/test")
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else {
			for _, v := range str {
				switch {
				case v >= 'a' && v <= 'z':
					fallthrough
				case v >= 'A' && v <= 'Z':
					strcount++
				case v == ' ' || v == '\t':
					spacecount++
				case v > '0' && v < '9':
					number++
				default:
					otherchar++
				}
			}
		}
	}
	fmt.Printf("%v %v %v %v\n", strcount, spacecount, number, otherchar)
	defer file.Close()
}
