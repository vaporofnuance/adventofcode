package main

import (
	"strconv"
	"strings"
)

type Bag struct {
	Name string
	ContainedBags []ContainedBag
}

type ContainedBag struct {
	Count int
	Name string
}

func LoadBags(lines []string) (result []Bag, err error) {
	result = make([]Bag, 0)

	for _, line := range lines {
		index := strings.Index(line, " bags contain ")
		bagName := line[:index]

		bagStrings := strings.Split(line[index + 14:], ", ")

		newBag := Bag{
			Name: bagName,
			ContainedBags: make([]ContainedBag, 0),
		}

		for _, bagString := range bagStrings {
			bagString = strings.TrimSpace(bagString)
			bagStringParts := strings.Split(bagString, " ")

			if bagStringParts[0] != "no" {
				if count, err := strconv.Atoi(bagStringParts[0]); err == nil {
					newBag.ContainedBags = append(newBag.ContainedBags, ContainedBag{
						Count: count,
						Name:  strings.Join(bagStringParts[1:len(bagStringParts)-1], " "),
					})
				} else {
					return result, err
				}
			}
		}

		result = append(result, newBag)
	}

	return result, err
}

func findBagsContaining(bags []Bag, bagName string, onlyDirectly bool) (result []string) {
	for _, bag := range bags {
		for _, containedBag := range bag.ContainedBags {
			if containedBag.Name == bagName {
				if !arrayContains(result, bag.Name) {
					result = append(result, bag.Name)
				}

				if !onlyDirectly {
					thisContainingBags := findBagsContaining(bags, bag.Name, onlyDirectly)

					for _, anotherBag := range thisContainingBags {
						if !arrayContains(result, anotherBag) {
							result = append(result, anotherBag)
						}
					}
				}
			}
		}
	}

	return result
}

func arrayContains(haystack []string, needle string) bool {
	for _, straw := range haystack {
		if straw == needle {
			return true
		}
	}

	return false
}

func calcContainedBags(bags []Bag, bagName string) (result int) {
	for _, bag := range bags {
		if bag.Name == bagName {
			for _, containedBag := range bag.ContainedBags {
				result = result + (containedBag.Count * (calcContainedBags(bags, containedBag.Name)+1))
			}
		}
	}

	return result
}

func CalcDay7Part1(lines []string) (result int64, err error) {

	bags, err := LoadBags(lines)

	if err == nil {
		containingBags := findBagsContaining(bags, "shiny gold", false)

		result = int64(len(containingBags))
	}

	return result, err
}

func CalcDay7Part2(lines []string) (result int64, err error) {

	bags, err := LoadBags(lines)

	if err == nil {
		result = int64(calcContainedBags(bags, "shiny gold"))
	}

	return result, err
}
