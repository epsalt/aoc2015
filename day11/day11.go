package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

func decode(i int) string {
	var r rune
	var ans strings.Builder
	i += 1

	for i > 0 {
		r = rune((i-1)%26 + 97)
		i = (i - 1) / 26
		ans.WriteString(fmt.Sprintf("%c", r))
	}
	return reverse(ans.String())
}

func encode(s string) int {
	var ans int
	for i, r := range reverse(s) {
		ans += (int(r) - 97 + 1) * int(math.Pow(26, float64(i)))
	}
	return ans - 1
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func straight(s string) bool {
	var substr string
	for i := 2; i < 26; i++ {
		substr = fmt.Sprintf("%c%c%c", i+95, i+96, i+97)
		if strings.Contains(s, substr) {
			return true
		}
	}
	return false
}

func ambiguous(s string) bool {
	for _, r := range []rune{'i', 'o', 'l'} {
		if strings.ContainsRune(s, r) {
			return true
		}
	}
	return false
}

func dupe(s string) bool {
	var substr string
	hits := 0

	for i := 0; i < 26; i++ {
		substr = fmt.Sprintf("%c%c", i+97, i+97)
		hits += strings.Count(s, substr)

		if hits > 1 {
			return true
		}
	}
	return false
}

func part1(input string) string {
	i := encode(input)
	for ambiguous(input) || !straight(input) || !dupe(input) {
		i++
		input = decode(i)
	}
	return input
}

func part2(input string) string {
	input = decode(encode(part1(input)) + 1)
	return part1(input)
}

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.TrimSpace(string(bytes))
	fmt.Println("Part1:", part1(input))
	fmt.Println("Part2:", part2(input))
}
