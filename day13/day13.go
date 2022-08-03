package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

func dfs(node string, start string, graph map[string]map[string]int, seated map[string]bool) (ans int) {
	happiness := make([]int, 0)
	seated[node] = true
	curr := 0

	for neighbour := range graph[node] {
		if !seated[neighbour] {
			curr = graph[node][neighbour] + graph[neighbour][node] + dfs(neighbour, start, graph, seated)
			happiness = append(happiness, curr)
		}
	}
	seated[node] = false

	if len(happiness) == 0 {
		return graph[node][start] + graph[start][node]
	}

	ans = math.MinInt
	for _, h := range happiness {
		if h > ans {
			ans = h
		}
	}
	return ans
}

func part1(graph map[string]map[string]int, start string) int {
	seated := make(map[string]bool)
	seated[start] = true

	return dfs(start, start, graph, seated)
}

func part2(graph map[string]map[string]int, start string) int {
	seated := make(map[string]bool)
	seated[start] = true

	for person := range graph {
		graph[person]["Evan"] = 0
	}

	m := make(map[string]int)
	for person := range graph {
		if person != "Evan" {
			m[person] = 0
		}
	}
	graph["Evan"] = m

	return dfs(start, start, graph, seated)
}

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var p1, p2 string
	var sign string
	var n int

	input := strings.TrimSpace(string(bytes))
	scanner := bufio.NewScanner(strings.NewReader(input))
	graph := make(map[string]map[string]int)

	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%s would %s %d happiness units by sitting next to %s.", &p1, &sign, &n, &p2)
		p2 = strings.Replace(p2, ".", "", -1)

		if _, ok := graph[p1]; !ok {
			graph[p1] = make(map[string]int)
		}

		switch sign {
		case "gain":
			graph[p1][p2] = n
		case "lose":
			graph[p1][p2] = -n
		}
	}

	fmt.Println("Part1:", part1(graph, p1))
	fmt.Println("Part2:", part2(graph, p1))
}
