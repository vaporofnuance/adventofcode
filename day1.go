package main

import (
	"errors"
	"strconv"
)

func CalcDay1Part1(lines []string) (result int64, err error) {
	intValues := make([]int64, 0)

	for _, line := range lines {
		intValue, err1 := strconv.Atoi(line)

		if err1 == nil {
			intValues = append(intValues, int64(intValue))
		}
	}

	for i := 0; i < len(intValues) - 1; i++ {
		for j := i + 1; j < len(intValues); j++ {
			if intValues[i] + intValues[j] == 2020 {
				return intValues[i] * intValues[j], nil
			}
		}
	}

	return result, errors.New("could not find match")
}

func CalcDay1Part2(lines []string) (result int64, err error) {
	intValues := make([]int64, 0)

	for _, line := range lines {
		intValue, err1 := strconv.Atoi(line)

		if err1 == nil {
			intValues = append(intValues, int64(intValue))
		}
	}

	for i := 0; i < len(intValues) - 2; i++ {
		for j := i + 1; j < len(intValues) - 1; j++ {
			for k := j + 1; k < len(intValues); k++ {
				if intValues[i]+intValues[j]+intValues[k] == 2020 {
					return intValues[i] * intValues[j] * intValues[k], nil
				}
			}
		}
	}

	return result, errors.New("could not find match")
}
