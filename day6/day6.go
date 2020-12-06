package day6

import (
	"io/ioutil"
	"os"
	"strings"
)

func RunPart1() (result int, err error) {
	f, err := os.Open("day6/data/answers.txt")

	if err == nil {
		defer f.Close()

		var data []byte

		if data, err = ioutil.ReadAll(f); err == nil {
			lines := strings.Split(string(data), "\n")

			currentGroup := ""

			for _, line := range lines {
				line = strings.TrimSpace(line)
				if len(line) == 0 {
					currentGroup = ""
				} else {
					for _, char := range line {
						if !strings.Contains(currentGroup, string(char)) {
							currentGroup = currentGroup + string(char)
							result++
						}
					}
				}
			}
		}
	}

	return result, err
}

func RunPart2() (result int, err error) {
	f, err := os.Open("day6/data/answers.txt")

	if err == nil {
		defer f.Close()

		var data []byte

		if data, err = ioutil.ReadAll(f); err == nil {
			lines := strings.Split(string(data), "\n")

			currentGroup := ""
			groupStart := true

			for _, line := range lines {
				line = strings.TrimSpace(line)
				if len(line) == 0 {
					groupStart = true
				} else {
					if groupStart {
						groupStart = false
						currentGroup = line
						result = result + len(line)
					} else {
						newGroup := currentGroup
						for _, char := range currentGroup {
							if !strings.Contains(line, string(char)) {
								newGroup = strings.ReplaceAll(newGroup, string(char), "")
								result = result - 1
							}
						}
						currentGroup = newGroup
					}
				}
			}
		}
	}

	return result, err
}