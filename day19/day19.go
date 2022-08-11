package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func part1(molecule string, replacements map[string][]string) int {
	combos := make(map[string]bool)

	for before := range replacements {
		r := regexp.MustCompile(before)

		for _, match := range r.FindAllStringIndex(molecule, -1) {
			for _, after := range replacements[before] {
				new := molecule[:match[0]] + after + molecule[match[1]:]
				combos[new] = true
			}
		}
	}
	return len(combos)
}

func part2(molecule string, replacements map[string]string) (ans int) {
outer:
	for molecule != "e" {
		for after := range replacements {
			if strings.Count(after, "Rn") > 1 && strings.Contains(molecule, after) {
				molecule = strings.Replace(molecule, after, replacements[after], 1)
				ans++
				continue outer
			}
		}
		for after := range replacements {
			if strings.Contains(molecule, after) {
				molecule = strings.Replace(molecule, after, replacements[after], 1)
				ans++
				continue outer
			}
		}
	}
	return ans
}

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.TrimSpace(string(bytes))
	scanner := bufio.NewScanner(strings.NewReader(input))

	replacements := make(map[string][]string)
	reverseReplacements := make(map[string]string)
	var molecule string

	for scanner.Scan() {
		s := scanner.Text()
		if strings.Contains(s, "=>") {
			start, end, _ := strings.Cut(s, " => ")

			if _, ok := replacements[start]; !ok {
				replacements[start] = make([]string, 0)
			}

			replacements[start] = append(replacements[start], end)
			reverseReplacements[end] = start
		} else if len(s) > 0 {
			molecule = s
		}
	}

	fmt.Println("Part1:", part1(molecule, replacements))
	fmt.Println("Part2:", part2(molecule, reverseReplacements))
}
