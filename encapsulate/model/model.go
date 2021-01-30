package model

import "fmt"

type person struct {
	Name string
	age int		//其他包无法访问
	salary float64		//	其他包无法访问
}

//写一个工厂模式的函数，相当于构造函数
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

//为了访问age和salary，我们需要编写一套get和set的方法

func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("年龄范围不对!")
	}
}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetSalary(salary float64) {
	if salary > 3000 && salary < 30000 {
		p.salary = salary
	} else {
		fmt.Println("薪水范围不对!")
	}
}

func (p *person) GetSalary() float64 {
	return p.salary
}