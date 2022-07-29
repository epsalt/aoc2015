package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func vowels(input string) int {
	ans := 0
	for _, rune := range input {
		switch rune {
		case 'a', 'e', 'i', 'o', 'u':
			ans++
		}
	}
	return ans
}

func hasDupe(input string) bool {
	for i := 1; i < len(input); i++ {
		if input[i] == input[i-1] {
			return true
		}
	}
	return false
}

func hasBad(input string) bool {
	for _, s := range []string{"ab", "cd", "pq", "xy"} {
		if strings.Contains(input, s) {
			return true
		}
	}
	return false
}

func twoPair(input string) bool {
	var curr, after string

	for i := 1; i < len(input); i++ {
		curr = string(input[i-1 : i+1])
		after = string(input[i+1:])

		if strings.Contains(after, curr) {
			return true
		}
	}
	return false
}

func separated(input string) bool {
	for i := 2; i < len(input); i++ {
		if (input[i] == input[i-2]) && (input[i] != input[i-1]) {
			return true
		}
	}
	return false
}

func part1(input string) int {
	ans := 0
	var s string
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		s = scanner.Text()
		if vowels(s) >= 3 && hasDupe(s) && !hasBad(s) {
			ans++
		}
	}
	return ans
}

func part2(input string) int {
	ans := 0
	var s string
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		s = scanner.Text()
		if twoPair(s) && separated(s) {
			ans++
		}
	}
	return ans
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
