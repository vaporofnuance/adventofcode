package main

import (
	"regexp"
	"strconv"
	"strings"
)

func hasKey(passport map[string]string, key string) bool {
	_, ok := passport[key]

	return ok
}

func hasRequiredFields(passport map[string]string) bool {
	return hasKey(passport, "byr") &&
		hasKey(passport, "iyr") &&
		hasKey(passport, "eyr") &&
		hasKey(passport, "hgt") &&
		hasKey(passport, "hcl") &&
		hasKey(passport, "ecl") &&
		hasKey(passport, "pid")
}

func validField(key string, value string) bool {
	switch key {
	case "byr":
		byr, err := strconv.Atoi(value)
		return err == nil && len(value) == 4 && byr >= 1920 && byr <= 2002
	case "iyr":
		iyr, err := strconv.Atoi(value)
		return err == nil && len(value) == 4 && iyr >= 2010 && iyr <= 2020
	case "eyr":
		eyr, err := strconv.Atoi(value)
		return err == nil && len(value) == 4 && eyr >= 2020 && eyr <= 2030
	case "pid":
		matched, err := regexp.Match("[0-9]{9}", []byte(value))
		return err == nil && matched && len(value) == 9
	case "hcl":
		matched, err := regexp.Match("#[0-9a-f]{6}", []byte(value))
		return err == nil && matched && len(value) == 7
	case "ecl":
		matched, err := regexp.Match("amb|blu|brn|gry|grn|hzl|oth", []byte(value))
		return err == nil && matched && len(value) == 3
	case "hgt":
		if matched, err := regexp.Match("[0-9]{3}cm", []byte(value)); err == nil && len(value) == 5 && matched {
			val, err := strconv.Atoi(value[:3])
			return err == nil && val >= 150 && val <= 193
		}
		if matched, err := regexp.Match("[0-9]{2}in", []byte(value)); err == nil && len(value) == 4 && matched {
			val, err := strconv.Atoi(value[:2])
			return err == nil && val >= 59 && val <= 76
		}
		return false
	case "cid":
		return true
	default:
		return false
	}
}
func hasValidData(passport map[string]string) (result bool) {
	result = hasRequiredFields(passport)

	for k, v := range passport {
		result = result && validField(k, v)
	}

	return result
}

func calcDay4ByPolicy(lines []string, policy func (map[string]string) bool) (result int64, err error) {
	passports := make([]map[string]string, 0)

	currentPassport := make(map[string]string, 0)
	passports = append(passports, currentPassport)

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			currentPassport = make(map[string]string, 0)
			passports = append(passports, currentPassport)
		} else {
			for _, field := range strings.Split(line, " ") {
				key := strings.Split(field, ":")[0]
				value := strings.Split(field, ":")[1]

				currentPassport[key] = value
			}
		}
	}

	for _, passport := range passports {
		if policy(passport) {
			result++
		}
	}

	return result, err
}

func CalcDay4Part1(lines []string) (result int64, err error) {
	return calcDay4ByPolicy(lines, hasRequiredFields)
}

func CalcDay4Part2(lines []string) (result int64, err error) {
	return calcDay4ByPolicy(lines, hasValidData)
}