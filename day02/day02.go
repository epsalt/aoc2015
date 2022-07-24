package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func wrapping(dims []int) int {
	sides := []int{dims[0] * dims[1], dims[1] * dims[2], dims[2] * dims[0]}
	minSide := math.MaxInt
	area := 0

	for _, side := range sides {
		if side < minSide {
			minSide = side
		}
		area += (side * 2)
	}
	return area + minSide
}

func ribbon(dims []int) int {
	minPermimeter := math.MaxInt
	perimeter := 0

	for i := range dims {
		for j := i + 1; j < len(dims); j++ {
			perimeter = 2*dims[i] + 2*dims[j]

			if perimeter < minPermimeter {
				minPermimeter = perimeter
			}
		}
	}
	return minPermimeter
}

func bow(dims []int) int {
	return dims[0] * dims[1] * dims[2]
}

func dims(input string) [][]int {
	s := make([][]int, 0, 0)

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		strings := strings.Split(scanner.Text(), "x")
		dims := make([]int, len(strings))

		for i, s := range strings {
			dims[i], _ = strconv.Atoi(s)
		}
		s = append(s, dims)
	}

	return s
}

func part1(input string) int {
	ans := 0
	lines := dims(input)

	for i := range lines {
		ans += wrapping(lines[i])
	}
	return ans
}

func part2(input string) int {
	ans := 0
	lines := dims(input)

	for i := range lines {
		ans += ribbon(lines[i]) + bow(lines[i])
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
