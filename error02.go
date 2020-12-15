package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(100)
	fmt.Println(num)
	guessNum := 0
	label:
	for i := 1; i <= 10; i++ {
		fmt.Print("请输入你想猜的数字: ")
		fmt.Scanln(&guessNum)
		if guessNum == num {
			switch i {
			case 1:
				fmt.Println("你真是个天才")
				break label
			case 2, 3:
				fmt.Println("你真聪明，快赶上我了")
				break label
			case 4,5,6,7,8,9:
				fmt.Println("一般般")
				break label
			case 10:
				fmt.Println("可算猜对了")
				break label
			case 11:
				fmt.Println("说点啥好呢")
				break label
			}
		} else {
			continue
		}
	}
}
