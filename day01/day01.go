package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func part1(input string) int {
	up := strings.Count(input, "(")
	down := strings.Count(input, ")")
	ans := up - down

	return ans
}

func part2(input string) int {
	floor := 0

	for pos, char := range input {
		if char == '(' {
			floor++
		} else {
			floor--
		}

		if floor == -1 {
			return pos + 1
		}

	}

	return -1
}

func main() {
	bytes, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	input := string(bytes)

	fmt.Println("Part1:", part1(input))
	fmt.Println("Part2:", part2(input))
}
