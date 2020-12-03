package main

import (
	"fmt"
	"github.com/vaporofnuance/adventofcode/day1"
	"github.com/vaporofnuance/adventofcode/day2"
	"github.com/vaporofnuance/adventofcode/day3"
)

func main() {
	fmt.Print("day 1 part 1= ")
	result, err := day1.RunPart1()
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}

	fmt.Print("day 1 part 2 = ")
	result, err = day1.RunPart2()
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}

	fmt.Print("day 2 part 1 = ")
	result, err = day2.Run(day2.NumberPolicy)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}

	fmt.Print("day 2 part 2 = ")
	result, err = day2.Run(day2.PositionPolicy)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}

	fmt.Print("day 3 part 1 = ")
	result, err = day3.RunPart1(0, 0, 3, 1)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}

	var result64 int64
	fmt.Print("day 3 part 2 = ")
	result64, err = day3.RunPart2()
	if err == nil {
		fmt.Println(result64)
	} else {
		fmt.Println(err)
	}
}
