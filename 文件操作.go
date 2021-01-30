package main

import (
	"log"
	"os"
)

func main() {

	file, error := os.Open("/etc/passwd1")
	if error != nil {
		log.Fatal("can't find the file!")
	}
	log.Fatal("%v", file.Name())

	defer file.Close()
}
