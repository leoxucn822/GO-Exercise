package main

import (
	"fmt"
	"encoding/json"
)

type Monster struct {
	Name string
	Age int
	Birthday string
	Sal float64
	Skill string
}

func testStruct() {
	monster := Monster{
		Name: "Leo",
		Age: 21,
		Birthday: "2020-12-01",
		Sal: 22.22,
		Skill: "Go",
	}

	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化失败, error是%v\n", err)
	}

	//输出序列化的结果
	fmt.Printf("monster序列化的结果%v\n", string(data))
}

func testMap() {
	var a map[string]interface{}
	//使用map之前需要make
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 18

	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化失败, error是%v\n", err)
	}

	//输出序列化的结果
	fmt.Printf("a序列化的结果%v\n", string(data))
}

func testSlice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "jackie"
	m1["sex"] = "Male"
	m1["address"] = "苏州"
	slice = append(slice, m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "cherry"
	m1["sex"] = "Female"
	m1["address"] = "上海"
	slice = append(slice, m2)

	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化失败, error是%v\n", err)
	}

	//输出序列化的结果
	fmt.Printf("a序列化的结果%v\n", string(data))
}

func main() {
	//演示struct, map, 切片的序列化
	testStruct()
	testMap()
	testSlice()
}
