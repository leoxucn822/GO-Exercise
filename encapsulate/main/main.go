package main

import (
	"fmt"
	"goproject/encapsulate/model"
)

func main() {
	p := model.NewPerson("Leo")
	fmt.Println(p)
}