package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func val(s string, r map[string]uint16, seen map[string]bool) (uint16, bool) {
	var i16 uint16
	i, err := strconv.Atoi(s)

	if err != nil {
		if !seen[s] {
			return 0, false
		}
		i16, _ = r[s]
		return i16, true
	}
	return uint16(i), true
}

func eval(s string, r map[string]uint16, seen map[string]bool) bool {
	var x, y string
	var i, j uint16
	var ok bool
	lhs, rhs, _ := strings.Cut(s, " -> ")

	switch {
	case strings.Contains(lhs, "AND"):
		fmt.Sscanf(lhs, "%s AND %s", &x, &y)

		if i, ok = val(x, r, seen); !ok {
			return false
		}
		if j, ok = val(y, r, seen); !ok {
			return false
		}

		r[rhs] = i & j

	case strings.Contains(lhs, "OR"):
		fmt.Sscanf(lhs, "%s OR %s", &x, &y)

		if i, ok = val(x, r, seen); !ok {
			return false
		}
		if j, ok = val(y, r, seen); !ok {
			return false
		}

		r[rhs] = i | j

	case strings.Contains(lhs, "NOT"):
		fmt.Sscanf(lhs, "NOT %s", &x)

		if i, ok = val(x, r, seen); !ok {
			return false
		}
		r[rhs] = ^i

	case strings.Contains(lhs, "LSHIFT"):
		fmt.Sscanf(lhs, "%s LSHIFT %s", &x, &y)
		if i, ok = val(x, r, seen); !ok {
			return false
		}
		if j, ok = val(y, r, seen); !ok {
			return false
		}
		r[rhs] = i << j

	case strings.Contains(lhs, "RSHIFT"):
		fmt.Sscanf(lhs, "%s RSHIFT %s", &x, &y)
		if i, ok = val(x, r, seen); !ok {
			return false
		}
		if j, ok = val(y, r, seen); !ok {
			return false
		}
		r[rhs] = i >> j

	default:
		if i, ok = val(lhs, r, seen); !ok {
			return false
		}
		r[rhs] = i
	}

	seen[rhs] = true
	return true
}

func part1(ops []string) uint16 {
	var ready bool
	var op string
	registers := make(map[string]uint16)
	seen := make(map[string]bool)

	for len(ops) > 0 {
		op, ops = ops[0], ops[1:]
		ready = eval(op, registers, seen)
		if !ready {
			ops = append(ops, op)
		}
	}
	return registers["a"]
}

func part2(ops []string) uint16 {
	var ready bool
	var op string
	registers := make(map[string]uint16)
	seen := make(map[string]bool)

	registers["b"] = part1(ops)
	seen["b"] = true

	for len(ops) > 0 {
		op, ops = ops[0], ops[1:]
		if strings.HasSuffix(op, "-> b") {
			continue
		}
		ready = eval(op, registers, seen)
		if !ready {
			ops = append(ops, op)
		}
	}
	return registers["a"]
}

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	ops := make([]string, 0)
	input := strings.TrimSpace(string(bytes))
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		ops = append(ops, scanner.Text())
	}

	fmt.Println("Part1:", part1(ops))
	fmt.Println("Part2:", part2(ops))
}
