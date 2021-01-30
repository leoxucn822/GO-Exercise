//golang中是没有构造函数的,使用工厂模式来解决这个问题
//举例说明如何定义的结构体是小写别的包想调用那肯定不行，所以只能用工厂模式来解决

package main

import (
	"fmt"
	"goproject/model"
)

func main() {
	var stu = model.NewStudent("Leo Xu", 60)
	fmt.Println(stu)
	fmt.Println(*stu)
}