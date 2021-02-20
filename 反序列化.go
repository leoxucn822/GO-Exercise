package main

import (
	"encoding/json"
	"fmt"
)

type Monster1 struct {
	Name string
	Age int
	Birthday string
	Sal float64
	Skill string
}

func UnMarshal() {
	str := "{\"Name\":\"Leo\",\"Age\":21,\"Birthday\":\"2020-12-01\",\"Sal\":22.22,\"Skill\":\"Go\"}"
	var monster Monster1
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("error is %v\n", err)
	}
	fmt.Printf("反序列化后monster的值: %v\n", monster)
}

func main() {
	UnMarshal()
}
