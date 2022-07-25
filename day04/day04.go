package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func hash(s string) string {
	data := []byte(s)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func mine(input string, goal string) int {
	i := 0
	for {
		h := hash(fmt.Sprintf("%s%d", input, i))
		if strings.HasPrefix(h, goal) {
			return i
		}
		i++
	}
}

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.TrimSpace(string(bytes))
	fmt.Println("Part1:", mine(input, "00000"))
	fmt.Println("Part2:", mine(input, "000000"))
}
