package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	rand.Seed(time.Now().Unix())
	fmt.Println("the rand number is", rand.Intn(100))
}