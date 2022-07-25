package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

type location struct {
	x, y int
}

func (loc *location) move(c rune) {
	switch c {
	case '^':
		loc.y++
	case 'v':
		loc.y--
	case '>':
		loc.x++
	case '<':
		loc.x--
	}
}

func part1(input string) int {
	set := make(map[location]bool)
	loc := location{0, 0}
	set[loc] = true

	for _, dir := range input {
		loc.move(dir)
		set[loc] = true
	}
	return len(set)
}

func part2(input string) int {
	set := make(map[location]bool)
	santa, robot := location{0, 0}, location{0, 0}
	set[santa] = true

	for i, dir := range input {
		if i%2 == 0 {
			santa.move(dir)
			set[santa] = true
		} else {
			robot.move(dir)
			set[robot] = true
		}
	}
	return len(set)
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
