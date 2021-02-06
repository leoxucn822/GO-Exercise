package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Stat("/etc/passwd")
	if err == nil {
		fmt.Println("文件存在")
	} else if os.IsNotExist(err) {
		fmt.Println("文件不存在")
	}
}