package main

import (
	"io/ioutil"
	"log"
)

func main() {
	filepath1 := "/etc/passwd"
	filepath2 := "/tmp/passwd"

	data, err := ioutil.ReadFile(filepath1)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(filepath2, data, 0666)
	if err != nil {
		log.Fatal(err)
	}
}
