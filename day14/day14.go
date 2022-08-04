package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type reindeer struct {
	name     string
	speed    int
	duration int
	rest     int
}

func part1(reindeers []reindeer, t int) (ans int) {
	var interval, remainder, dist int

	for _, r := range reindeers {
		interval = r.duration + r.rest
		dist = (t / interval) * r.duration * r.speed
		remainder = t % interval

		if remainder <= r.duration {
			dist += remainder * r.speed
		} else {
			dist += r.duration * r.speed
		}

		if dist > ans {
			ans = dist
		}
	}
	return ans
}

func part2(reindeers []reindeer, end int) (ans int) {
	var interval, remainder, dist, curr int
	scores := make(map[reindeer]int)
	dists := make(map[reindeer]int)

	for t := 1; t < end+1; t++ {
		for _, r := range reindeers {
			interval = r.duration + r.rest
			dist = (t / interval) * r.duration * r.speed
			remainder = t % interval

			if remainder <= r.duration {
				dist += remainder * r.speed
			} else {
				dist += r.duration * r.speed
			}

			if dist > curr {
				curr = dist
			}
			dists[r] = dist
		}

		for _, r := range reindeers {
			if dists[r] == curr {
				scores[r]++

				if scores[r] > ans {
					ans = scores[r]
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
	var speed, duration, rest int

	input := strings.TrimSpace(string(bytes))
	scanner := bufio.NewScanner(strings.NewReader(input))
	reindeers := make([]reindeer, 0)

	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &name, &speed, &duration, &rest)
		reindeers = append(reindeers, reindeer{name, speed, duration, rest})
	}

	fmt.Println("Part1:", part1(reindeers, 2503))
	fmt.Println("Part2:", part2(reindeers, 2503))
}
