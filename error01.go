package main

import (
	"errors"
	"fmt"
)

func printdate(year int, month int, day int) (err error) {
	if month >= 1 && month <= 12 {
		return nil
	} else {
		return errors.New("月份只能在1到12之间")
	}
}

func testdate(year int, month int, day int) {
	err := printdate(year, month, day)
	if err == nil {
		fmt.Printf("%v年%v月%v日", year, month, day)
	} else {
		fmt.Println(printdate(year, month, day))
	}
}

func main() {
	var year, month, day int
	fmt.Scanln(&year, &month, &day)

	testdate(year, month, day)
}
