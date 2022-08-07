package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type xy struct {
	x, y int
}

func step(grid [][]int) [][]int {
	next := make([][]int, 0)
	for i, row := range grid {
		nextRow := make([]int, 0)
		for j, val := range row {
			count := 0
			for _, n := range []xy{{i + 1, j}, {i, j + 1}, {i - 1, j}, {i, j - 1}, {i + 1, j + 1}, {i - 1, j - 1}, {i + 1, j - 1}, {i - 1, j + 1}} {
				if n.x >= 0 && n.x < len(grid) && n.y >= 0 && n.y < len(grid[0]) && grid[n.x][n.y] == 1 {
					count++
				}
			}
			nextVal := 0
			switch {
			case val == 1 && (count == 2 || count == 3):
				nextVal = 1
			case val == 0 && count == 3:
				nextVal = 1
			}
			nextRow = append(nextRow, nextVal)
		}
		next = append(next, nextRow)
	}
	return next
}

func part1(grid [][]int, n int) (ans int) {
	for i := 0; i < n; i++ {
		grid = step(grid)
	}
	for _, row := range grid {
		for _, val := range row {
			ans += val
		}
	}
	return ans
}

func part2(grid [][]int, n int) (ans int) {
	r, c := len(grid)-1, len(grid[0])-1

	for i := 0; i < n; i++ {
		grid[0][0], grid[0][c], grid[r][0], grid[r][c] = 1, 1, 1, 1
		grid = step(grid)
	}

	grid[0][0], grid[0][c], grid[r][0], grid[r][c] = 1, 1, 1, 1
	for _, row := range grid {
		for _, val := range row {
			ans += val
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

	grid := make([][]int, 0)
	for scanner.Scan() {
		row := make([]int, 0)
		for _, rune := range scanner.Text() {
			switch rune {
			case '#':
				row = append(row, 1)
			case '.':
				row = append(row, 0)
			}
		}
		grid = append(grid, row)
	}

	fmt.Println("Part1:", part1(grid, 100))
	fmt.Println("Part2:", part2(grid, 100))
}
