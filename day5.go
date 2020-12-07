package main

import (
	"errors"
	"math"
)

func calcNumber(value string, upper int32) (result int) {
	for i, char := range value {
		if char == upper {
			result += int(math.Pow(2, float64(len(value) - i - 1)))
		}
	}

	return result
}

func CalcDay5Part1(lines []string) (result int64, err error) {
	for _, line := range lines {
		row := calcNumber(line[:7], 'B')
		column := calcNumber(line[7:], 'R')

		seat := (row * 8) + column

		if seat > int(result) {
			result = int64(seat)
		}
	}

	return result, err
}

func getSeat(line string) (result int) {
	row := calcNumber(line[:7], 'B')
	column := calcNumber(line[7:], 'R')

	return (row * 8) + column
}

func CalcDay5Part2(lines []string) (result int64, err error) {
	for _, line1 := range lines {
		seat1 := getSeat(line1)

		foundSeat2 := false
		foundSeat3 := false

		for _, line2 := range lines {
			seat2 := getSeat(line2)

			if seat2 == seat1 + 1 {
				foundSeat2 = true
			} else if seat2 == seat1 + 2 {
				foundSeat3 = true
			}
		}

		if !foundSeat2 && foundSeat3 {
			return int64(seat1 + 1), nil
		}
	}

	return result, errors.New("seat not found")
}
