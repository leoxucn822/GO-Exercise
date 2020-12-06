package main

import (
	"fmt"
	"strings"
)

func makeSuffix(suffix string) func (string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	f := makeSuffix(".jpg")
	fmt.Println(f("Linux"))
	fmt.Println(f("windows.jpg"))
}
