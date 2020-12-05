package day5

import (
	"io/ioutil"
	"math"
	"os"
	"strings"
)

func CalcNumber(value string, upper int32) (result int) {
	for i, char := range value {
		if char == upper {
			result += int(math.Pow(2, float64(len(value) - i - 1)))
		}
	}

	return result
}

func Run() (result int, err error) {
	f, err := os.Open("day5/data/seats.txt")

	if err == nil {
		defer f.Close()
		var data []byte

		if data, err = ioutil.ReadAll(f); err == nil {
			lines := strings.Split(string(data), "\n")

			for _, line := range lines {
				row := CalcNumber(line[:7], 'B')
				column := CalcNumber(line[7:], 'R')

				seat := (row * 8) + column

				if seat > result {
					result = seat
				}
			}
		}
	}

	return result, err
}

func GetSeat(line string) (result int) {
	row := CalcNumber(line[:7], 'B')
	column := CalcNumber(line[7:], 'R')

	return (row * 8) + column
}

func RunPart2() (result int, err error) {
	f, err := os.Open("day5/data/seats.txt")

	if err == nil {
		defer f.Close()
		var data []byte

		if data, err = ioutil.ReadAll(f); err == nil {
			lines := strings.Split(string(data), "\n")

			for _, line1 := range lines {
				seat1 := GetSeat(line1)

				foundSeat2 := false
				foundSeat3 := false

				for _, line2 := range lines {
					seat2 := GetSeat(line2)

					if seat2 == seat1 + 1 {
						foundSeat2 = true
					} else if seat2 == seat1 + 2 {
						foundSeat3 = true
					}
				}

				if !foundSeat2 && foundSeat3 {
					return seat1 + 1, nil
				}
			}
		}
	}

	return result, err
}