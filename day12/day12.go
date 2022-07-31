package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func count(val any, f func(any) int) int {
	switch x := val.(type) {
	case float64:
		return int(x)
	case map[string]any, []any:
		return f(x)
	}
	return 0
}

func part1(json any) int {
	ans := 0
	switch x := json.(type) {
	case []any:
		for _, v := range x {
			ans += count(v, part1)
		}
	case map[string]any:
		for _, v := range x {
			ans += count(v, part1)
		}
	}
	return ans
}

func part2(json any) int {
	ans := 0
	switch x := json.(type) {
	case []any:
		for _, v := range x {
			ans += count(v, part2)
		}
	case map[string]any:
		for _, v := range x {
			if s, ok := v.(string); ok {
				if s == "red" {
					return 0
				}
			}
			ans += count(v, part2)
		}
	}
	return ans
}

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var input any
	json.Unmarshal(bytes, &input)

	fmt.Println("Part1:", part1(input))
	fmt.Println("Part2:", part2(input))
}
