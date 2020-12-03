package day1

import (
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func RunPart1() (result int, err error) {
	f, err := os.Open("day1/data/data.txt")

	if err == nil {
		defer f.Close()

		var data []byte

		data, err = ioutil.ReadAll(f)

		if err == nil {
			lines := strings.Split(string(data), "\n")

			intValues := make([]int, 0)

			for _, line := range lines {
				intValue, err1 := strconv.Atoi(line)

				if err1 == nil {
					intValues = append(intValues, intValue)
				}
			}

			for i := 0; i < len(intValues) - 1; i++ {
				for j := i + 1; j < len(intValues); j++ {
					if intValues[i] + intValues[j] == 2020 {
						return intValues[i] * intValues[j], nil
					}
				}
			}

			err = errors.New("could not find match")
		}
	}

	return result, err
}

func RunPart2() (result int, err error) {
	f, err := os.Open("day1/data/data.txt")

	if err == nil {
		defer f.Close()

		var data []byte

		data, err = ioutil.ReadAll(f)

		if err == nil {
			lines := strings.Split(string(data), "\n")

			intValues := make([]int, 0)

			for _, line := range lines {
				intValue, err1 := strconv.Atoi(line)

				if err1 == nil {
					intValues = append(intValues, intValue)
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

			err = errors.New("could not find match")
		}
	}

	return result, err
}