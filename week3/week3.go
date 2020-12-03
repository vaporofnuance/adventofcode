package week3

import (
	"io/ioutil"
	"os"
	"strings"
)

func Run(startX int, startY int, slopeX int, slopeY int) (result int, err error) {
	f, err := os.Open("week3/data/map.txt")

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