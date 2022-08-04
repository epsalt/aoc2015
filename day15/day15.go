package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type cookie struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func part1(cookies []cookie) (ans int) {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100-i; j++ {
			for k := 0; k < 100-i-j; k++ {
				l := 100 - i - j - k

				var scores cookie
				for cookie, amount := range []int{i, j, k, l} {
					scores.capacity += cookies[cookie].capacity * amount
					scores.durability += cookies[cookie].durability * amount
					scores.flavor += cookies[cookie].flavor * amount
					scores.texture += cookies[cookie].texture * amount
				}

				curr := 1
				for _, score := range []int{scores.capacity, scores.durability, scores.flavor, scores.texture} {
					if score > 0 {
						curr *= score
					} else {
						curr = 0
					}
				}
				if ans < curr {
					ans = curr
				}
			}
		}
	}
	return ans
}

func part2(cookies []cookie) (ans int) {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100-i; j++ {
			for k := 0; k < 100-i-j; k++ {
				l := 100 - i - j - k

				var scores cookie
				for cookie, amount := range []int{i, j, k, l} {
					scores.capacity += cookies[cookie].capacity * amount
					scores.durability += cookies[cookie].durability * amount
					scores.flavor += cookies[cookie].flavor * amount
					scores.texture += cookies[cookie].texture * amount
					scores.calories += cookies[cookie].calories * amount
				}

				curr := 1
				for _, score := range []int{scores.capacity, scores.durability, scores.flavor, scores.texture} {
					if score > 0 {
						curr *= score
					} else {
						curr = 0
					}
				}
				if ans < curr && scores.calories == 500 {
					ans = curr
				}
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

	var name string
	var capacity, durability, flavor, texture, calories int

	input := strings.TrimSpace(string(bytes))
	scanner := bufio.NewScanner(strings.NewReader(input))
	cookies := make([]cookie, 0)

	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%s capacity %d, durability %d, flavor %d, texture %d, calories %d", &name, &capacity, &durability, &flavor, &texture, &calories)
		cookies = append(cookies, cookie{name, capacity, durability, flavor, texture, calories})
	}
	fmt.Println(cookies)

	fmt.Println("Part1:", part1(cookies))
	fmt.Println("Part2:", part2(cookies))
}
