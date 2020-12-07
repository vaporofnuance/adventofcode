package main

import (
	"strings"
)

func CalcDay6Part1(lines []string) (result int64, err error) {
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

	return result, err
}

func CalcDay6Part2(lines []string) (result int64, err error) {
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
				result = result + int64(len(line))
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

	return result, err
}