package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type aunt map[string]int

func part1(aunts []aunt, goal aunt) (ans int) {
outer:
	for i, a := range aunts {
		for prop := range a {
			if goal[prop] != a[prop] {
				continue outer
			}
		}
		return i + 1
	}
	panic("expected return in loop")
}

func part2(aunts []aunt, goal aunt) (ans int) {
outer:
	for i, a := range aunts {
		for prop := range a {
			switch prop {
			case "cats", "trees":
				if a[prop] <= goal[prop] {
					continue outer
				}
			case "pomeranians", "goldfish":
				if a[prop] >= goal[prop] {
					continue outer
				}
			default:
				if goal[prop] != a[prop] {
					continue outer
				}
			}
		}
		return i + 1
	}
	panic("expected return in loop")
}

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.TrimSpace(string(bytes))
	scanner := bufio.NewScanner(strings.NewReader(input))
	aunts := make([]aunt, 0)

	for scanner.Scan() {
		var t1, t2, t3 string
		var n, n1, n2, n3 int

		fmt.Sscanf(scanner.Text(), "Sue %d: %s %d, %s %d, %s %d", &n, &t1, &n1, &t2, &n2, &t3, &n3)
		t1, t2, t3 = t1[:len(t1)-1], t2[:len(t2)-1], t3[:len(t3)-1]
		aunts = append(aunts, aunt{t1: n1, t2: n2, t3: n3})
	}

	goal := aunt{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}

	fmt.Println("Part1:", part1(aunts, goal))
	fmt.Println("Part2:", part2(aunts, goal))
}
