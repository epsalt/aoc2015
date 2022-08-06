package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/bits"
	"strconv"
	"strings"
)

func part1(containers []int, n, used int) (ans int) {
	if n == 0 {
		return 1
	}

	for i, c := range containers {
		if (used & (1 << i)) == 0 {
			used |= (1 << i)
			if c <= n {
				ans += part1(containers, n-c, used)
			}
			used |= (1 << i)
		}
	}
	return ans
}

func part2(containers []int, n int, used, left uint, found *[]uint) (ans int) {
	if n == 0 {
		*found = append(*found, used)
	}

	for i, c := range containers {
		if (left & (1 << i)) == 0 {
			left |= (1 << i)
			if c <= n {
				used |= (1 << i)
				part2(containers, n-c, used, left, found)
				used &^= (1 << i)
			}
		}
	}

	minContainers := math.MaxInt
	results := make(map[int]int)

	for _, f := range *found {
		curr := bits.OnesCount(f)
		if curr < minContainers {
			minContainers = curr
		}
		results[curr]++
	}
	return results[minContainers]
}

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.TrimSpace(string(bytes))
	scanner := bufio.NewScanner(strings.NewReader(input))
	arr := make([]int, 0)
	found := make([]uint, 0)

	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		arr = append(arr, n)
	}

	fmt.Println("Part1:", part1(arr, 150, 0))
	fmt.Println("Part2:", part2(arr, 150, 0, 0, &found))
}
