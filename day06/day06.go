package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type cmd int

const (
	on cmd = iota
	off
	toggle
)

type location struct {
	x, y int
}

type instruction struct {
	cmd
	start location
	end   location
}

func onoff(light int, c cmd) int {
	switch c {
	case on:
		light = 1
	case off:
		light = 0
	case toggle:
		if light == 0 {
			light = 1
		} else {
			light = 0
		}
	}
	return light
}

func brightness(light int, c cmd) int {
	switch c {
	case on:
		light++
	case off:
		if light > 0 {
			light--
		} else {
			light = 0
		}
	case toggle:
		light += 2
	}
	return light
}

func part1(instructions []instruction) int {
	ans := 0
	grid := [1000][1000]int{}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			for _, c := range instructions {
				if c.start.x <= i && i <= c.end.x && c.start.y <= j && j <= c.end.y {
					grid[i][j] = onoff(grid[i][j], c.cmd)
				}
			}
			ans += grid[i][j]
		}
	}
	return ans
}

func part2(instructions []instruction) int {
	ans := 0
	grid := [1000][1000]int{}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			for _, c := range instructions {
				if c.start.x <= i && i <= c.end.x && c.start.y <= j && j <= c.end.y {
					grid[i][j] = brightness(grid[i][j], c.cmd)
				}
			}
			ans += grid[i][j]
		}
	}
	return ans
}

func parse(input string) []instruction {
	var s string
	var nums []string
	var x1, y1, x2, y2 int
	var curr instruction

	scanner := bufio.NewScanner(strings.NewReader(input))
	r, _ := regexp.Compile("([0-9]+)")
	instructions := make([]instruction, 0)

	for scanner.Scan() {
		s = scanner.Text()
		nums = r.FindAllString(s, 4)

		x1, _ = strconv.Atoi(nums[0])
		y1, _ = strconv.Atoi(nums[1])
		x2, _ = strconv.Atoi(nums[2])
		y2, _ = strconv.Atoi(nums[3])

		curr = instruction{}
		curr.start = location{x1, y1}
		curr.end = location{x2, y2}

		switch {
		case strings.Contains(s, "toggle"):
			curr.cmd = toggle
		case strings.Contains(s, "on"):
			curr.cmd = on
		case strings.Contains(s, "off"):
			curr.cmd = off
		}
		instructions = append(instructions, curr)
	}
	return instructions
}

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := string(bytes)
	instructions := parse(input)

	fmt.Println("Part1:", part1(instructions))
	fmt.Println("Part2:", part2(instructions))
}
