package main

import "fmt"

type student struct {
	name string
	score float64
}

type primary struct {
	student
}

type junior struct {
	student
}

type senior struct {
	student
}

func (stu *student) GetScore(name string) {
	fmt.Println(stu.score)
}

func (p *primary) PrintStatus() {
	fmt.Println("小学生正在考试中")
}

func (j *junior) PrintStatus() {
	fmt.Println("中学生正在考试中")
}

func (s *senior) PrintStatus() {
	fmt.Println("高中生正在考试中")
}

func main() {
	//stu := student{
	//	name : "Leo Xu",
	//	score : 90,
	//}
	//stu.GetScore(stu.name)

	xiaoxuesheng := primary{}
	xiaoxuesheng.name = "Leo Xu"
	xiaoxuesheng.score = 90
	xiaoxuesheng.GetScore("Leo XU")
	xiaoxuesheng.PrintStatus()
}