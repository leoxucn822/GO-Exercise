package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, error := os.Open("/etc/passwd")
	if error != nil {
		log.Fatal("the file isn't found!")
	}

	reader := bufio.NewReader(file)
	for {
		str, error := reader.ReadString('\n')
		if error == io.EOF {
			break
		}
		fmt.Print(str)
	}

	defer file.Close()
}