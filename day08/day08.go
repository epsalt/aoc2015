package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func decode(s string) int {
	var r *regexp.Regexp
	code := len(s)

	// Remove outer quotes so we don't try to escape them with an escaped slash
	s = s[1 : len(s)-1]
	memory := len(s)

	for _, pattern := range []string{`\\"`, `\\\\`, `\\x[0-9a-f]{2}`} {
		r = regexp.MustCompile(pattern)
		for _, match := range r.FindAllString(s, -1) {
			memory -= len(match) - 1
		}
	}
	return code - memory
}

func part1(input []string) int {
	ans := 0
	for _, s := range input {
		ans += decode(s)
	}
	return ans
}

func part2(input []string) int {
	ans := 0
	for _, s := range input {
		ans += strings.Count(s, `\`) + strings.Count(s, `"`) + 2
	}
	return ans
}

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	strs := make([]string, 0)
	input := strings.TrimSpace(string(bytes))
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		strs = append(strs, scanner.Text())
	}

	fmt.Println("Part1:", part1(strs))
	fmt.Println("Part2:", part2(strs))
}
