package day2

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func NumberPolicy(password string, char uint8, min int, max int) bool {
	var count int

	for i, _ := range password {
		if password[i] == char {
			count++
		}
	}

	return count >= min && count <= max
}

func PositionPolicy(password string, char uint8, first int, second int) bool {
	char1 := password[first - 1]
	char2 := password[second - 1]

	return (char == char1 && char != char2) || (char != char1 && char == char2)
}

func Run(policy func(string, uint8, int, int) bool) (result int, err error) {
	if f, err := os.Open("day2/data/passwords.txt"); err == nil {
		defer f.Close()

		var data []byte

		if data, err = ioutil.ReadAll(f); err == nil {
			lines := strings.Split(string(data), "\n")

			for _, line := range lines {
				parts := strings.Split(line, ":")

				if len(parts) == 2 {
					regPart1 := strings.Split(parts[0], " ")[0]
					regPart2 := strings.Split(parts[0], " ")[1]

					char := regPart2[0]

					if min, err := strconv.Atoi(strings.Split(regPart1, "-")[0]); err == nil {
						if max, err := strconv.Atoi(strings.Split(regPart1, "-")[1]); err == nil {
							if policy(strings.TrimSpace(parts[1]), char, min, max) {
								result++
							}
						}
					}
				}
			}
		}
	}

	return result, err
}