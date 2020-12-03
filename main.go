package main

import (
	"fmt"
	"github.com/vaporofnuance/adventofcode/week1"
	"github.com/vaporofnuance/adventofcode/week2"
	"github.com/vaporofnuance/adventofcode/week3"
	"github.com/vaporofnuance/adventofcode/week4"
)

func main() {
	fmt.Print("week 1 = ")
	result, err := week1.RunWeek1()
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}

	fmt.Print("week 2 = ")
	result, err = week2.RunWeek2()
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}

	fmt.Print("week 3 = ")
	result, err = week3.Run(0, 0, 3, 1)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}

	var result64 int64
	fmt.Print("week 4 = ")
	result64, err = week4.Run()
	if err == nil {
		fmt.Println(result64)
	} else {
		fmt.Println(err)
	}
}
