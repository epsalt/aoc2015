package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func lookSay(input string) string {
	var ans strings.Builder
	runes := []rune(input)
	last := runes[0]
	run := 1

	for i := 1; i < len(runes); i++ {
		if runes[i] == last {
			run++
		} else {
			ans.WriteString(fmt.Sprintf("%d%c", run, runes[i-1]))
			run = 1
		}
		last = runes[i]
	}
	ans.WriteString(fmt.Sprintf("%d%c", run, runes[len(runes)-1]))
	return ans.String()
}

func part1(input string) int {
	s := input
	for i := 0; i < 40; i++ {
		s = lookSay(s)
	}
	return len(s)
}

func part2(input string) int {
	s := input
	for i := 0; i < 50; i++ {
		s = lookSay(s)
	}
	return len(s)
}

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.TrimSpace(string(bytes))
	fmt.Println("Part1:", part1(input))
	fmt.Println("Part2:", part2(input))
}
