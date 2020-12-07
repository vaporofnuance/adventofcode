package main

import (
)

func calcDay3ByCoords(lines []string, startX int, startY int, slopeX int, slopeY int) (result int64, err error) {
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

	return result, err
}

func CalcDay3Part1(lines []string) (result int64, err error) {
	return calcDay3ByCoords(lines, 0, 0, 3, 1)
}

func CalcDay3Part2(lines []string) (result int64, err error) {
	var i, j, k, l, m int64
	if i, err = calcDay3ByCoords(lines,0, 0, 1, 1); err == nil {
		if j, err = calcDay3ByCoords(lines, 0, 0, 3, 1); err == nil {
			if k, err = calcDay3ByCoords(lines, 0, 0, 5, 1); err == nil {
				if l, err = calcDay3ByCoords(lines, 0, 0, 7, 1); err == nil {
					if m, err = calcDay3ByCoords(lines, 0, 0, 1, 2); err == nil {
						result = i * j * k * l * m
					}
				}
			}
		}
	}

	return result, err
}
