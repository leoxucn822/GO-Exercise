package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	str, error := ioutil.ReadFile("/etc/passwd")
	if error != nil {
		log.Fatal("the file isn't here!")
	}
	fmt.Println(string(str))
}
