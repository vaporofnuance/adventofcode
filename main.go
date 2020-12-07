package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func RunCalc(day int, part int, filename string, calcFunc func ([]string) (int64, error)) {
	fmt.Printf("day %d part %d = ", day, part)
	f, err := os.Open(filename)

	if err == nil {
		var data []byte

		if data, err = ioutil.ReadAll(f); err == nil {
			lines := strings.Split(string(data), "\n")

			var result int64

			result, err = calcFunc(lines)

			if err == nil {
				fmt.Println(result)
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}

func main() {
	RunCalc(2, 1, "data/day1.txt", CalcDay1Part1)
	RunCalc(2, 2, "data/day1.txt", CalcDay1Part2)

	RunCalc(2, 1, "data/day2.txt", CalcDay2Part1)
	RunCalc(2, 2, "data/day2.txt", CalcDay2Part2)

	RunCalc(3, 1, "data/day3.txt", CalcDay3Part1)
	RunCalc(3, 2, "data/day3.txt", CalcDay3Part2)

	RunCalc(4, 1, "data/day4.txt", CalcDay4Part1)
	RunCalc(4, 2, "data/day4.txt", CalcDay4Part2)

	RunCalc(5, 1, "data/day5.txt", CalcDay5Part1)
	RunCalc(5, 2, "data/day5.txt", CalcDay5Part2)

	RunCalc(6, 1, "data/day6.txt", CalcDay6Part1)
	RunCalc(6, 2, "data/day6.txt", CalcDay6Part2)

	RunCalc(7, 1, "data/day7.txt", CalcDay7Part1)
	RunCalc(7, 2, "data/day7.txt", CalcDay7Part2)
}
