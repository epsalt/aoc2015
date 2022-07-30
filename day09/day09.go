package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

type route struct {
	start, end string
	distance   int
}

func shortDFS(start, visited, end int, graph map[int]map[int]int) int {
	visited = visited ^ (1 << start)
	if visited == end {
		return 0
	}
	curr, ans := 0, math.MaxInt
	for neighbour := range graph[start] {
		if (visited & (1 << neighbour)) == 0 {
			curr = graph[start][neighbour] + shortDFS(neighbour, visited, end, graph)

			if curr < ans {
				ans = curr
			}

		}
	}
	return ans
}

func longDFS(start, visited, end int, graph map[int]map[int]int) int {
	visited = visited ^ (1 << start)
	if visited == end {
		return 0
	}
	curr, ans := 0, 0
	for neighbour := range graph[start] {
		if (visited & (1 << neighbour)) == 0 {
			curr = graph[start][neighbour] + longDFS(neighbour, visited, end, graph)

			if curr > ans {
				ans = curr
			}

		}
	}
	return ans
}

func part1(graph map[int]map[int]int, n int) int {
	var curr int
	end := (1 << n) - 1
	ans := math.MaxInt

	for i := 0; i < n; i++ {
		curr = shortDFS(i, 0, end, graph)
		if curr < ans {
			ans = curr
		}
	}
	return ans
}

func part2(graph map[int]map[int]int, n int) int {
	var curr int
	end := (1 << n) - 1
	ans := 0

	for i := 0; i < n; i++ {
		curr = longDFS(i, 0, end, graph)
		if curr > ans {
			ans = curr
		}
	}
	return ans
}

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var start, end string
	var distance int

	routes := make([]route, 0)
	input := strings.TrimSpace(string(bytes))
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%s to %s = %d", &start, &end, &distance)
		routes = append(routes, route{start, end, distance})
	}

	n := 0
	places := make(map[string]int)
	graph := make(map[int]map[int]int)

	for _, r := range routes {
		if _, ok := places[r.start]; !ok {
			places[r.start] = n
			graph[places[r.start]] = make(map[int]int)
			n++
		}

		if _, ok := places[r.end]; !ok {
			places[r.end] = n
			graph[places[r.end]] = make(map[int]int)
			n++
		}

		graph[places[r.start]][places[r.end]] = r.distance
		graph[places[r.end]][places[r.start]] = r.distance
	}

	fmt.Println("Part1:", part1(graph, n))
	fmt.Println("Part2:", part2(graph, n))
}
