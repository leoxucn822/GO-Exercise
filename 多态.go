//多态实现二种方式
//1. 通过函数参数形式把接口传递进去
//2. 创建接口数组形式来实现多态

package main

import "fmt"

type USB interface {
	start()
	stop()
}

type Phone struct {
	name string
}

type Camera struct {
	name string
}

func (p Phone) start() {
	fmt.Println("手机开始工作")
}

func (p Phone) stop() {
	fmt.Println("手机停止工作")
}

func (c Camera) start() {
	fmt.Println("电脑开始工作")
}

func (c Camera) stop() {
	fmt.Println("电脑停止工作")
}

type Computer struct {
}

func (c Computer) Working(usb USB) {
	usb.start()
	usb.stop()
}

func main() {
	phone := Phone{}
	camera := Camera{}
	computer := Computer{}
	computer.Working(phone)
	computer.Working(camera)

	var test [3]USB
	test[0] = Phone{"Apple"}
	test[1] = Camera{"佳能"}
	test[2] = Phone{"华为"}

	fmt.Println(test)
}
