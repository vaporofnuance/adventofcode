package day3

import (
	"io/ioutil"
	"os"
	"strings"
)

func RunPart1(startX int, startY int, slopeX int, slopeY int) (result int, err error) {
	f, err := os.Open("day3/data/map.txt")

	if err == nil {
		defer f.Close()

		var data []byte

		if data, err = ioutil.ReadAll(f); err == nil {
			lines := strings.Split(string(data), "\n")

			locX := startX
			locY := startY

			for locY < len(lines) {
				line := lines[locY]

				if len(line) > 0 {
					for {
						part1 := []byte(line)
						part2 := []byte(line)
						line = string(append(part1, part2...))
						if locX < len(line) {
							break
						}
					}

					if string(line[locX]) == "#" {
						result++
					}

					locX = locX + slopeX
					locY = locY + slopeY
				}
			}
		}
	}

	return result, err
}

func RunPart2() (result int64, err error) {
	var i, j, k, l, m int
	if i, err = RunPart1(0, 0, 1, 1); err == nil {
		if j, err = RunPart1(0, 0, 3, 1); err == nil {
			if k, err = RunPart1(0, 0, 5, 1); err == nil {
				if l, err = RunPart1(0, 0, 7, 1); err == nil {
					if m, err = RunPart1(0, 0, 1, 2); err == nil {
						result = int64(i) * int64(j) * int64(k) * int64(l) * int64(m)
					}
				}
			}
		}
	}

	return result, err
}