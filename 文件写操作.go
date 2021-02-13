package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	str, error := os.OpenFile("/tmp/test", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0644)
	if error != nil {
		log.Fatal("there is a issue while open the files!")
	}
	defer str.Close()
	write := bufio.NewWriter(str)
	//有些编辑器识别\r有些识别\n
	write.WriteString("hello go programming language!\r\n")
	write.Flush()
}
