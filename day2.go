package main

import (
	"strconv"
	"strings"
)

func day2NumberPolicy(password string, char uint8, min int, max int) bool {
	var count int

	for i, _ := range password {
		if password[i] == char {
			count++
		}
	}

	return count >= min && count <= max
}

func day2PositionPolicy(password string, char uint8, first int, second int) bool {
	char1 := password[first - 1]
	char2 := password[second - 1]

	return (char == char1 && char != char2) || (char != char1 && char == char2)
}

func calcDay2ByPolicy(lines []string, policy func(string, uint8, int, int) bool) (result int64, err error) {
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

	return result, err
}

func CalcDay2Part1(lines []string) (result int64, err error) {
	return calcDay2ByPolicy(lines, day2NumberPolicy)
}

func CalcDay2Part2(lines []string) (result int64, err error) {
	return calcDay2ByPolicy(lines, day2PositionPolicy)
}
