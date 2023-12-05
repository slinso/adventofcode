package main

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type Game struct {
	ID int

	Red   int
	Blue  int
	Green int
}

func (g Game) Possible(red int, blue int, green int) bool {
	return red >= g.Red && blue >= g.Blue && green >= g.Green
}

func (g Game) Power() int {
	return g.Red * g.Blue * g.Green
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		sum := 0

		scanner := bufio.NewScanner(strings.NewReader(input))
		for scanner.Scan() {
			line := strings.Split(scanner.Text(), " ")

			g := Game{}
			g.ID, _ = strconv.Atoi(line[1][:len(line[1])-1])

			for i := 2; i < len(line); i += 2 {
				v, _ := strconv.Atoi(line[i])

				switch {
				case strings.Contains(line[i+1], "red"):
					if v > g.Red {
						g.Red = v
					}
				case strings.Contains(line[i+1], "blue"):
					if v > g.Blue {
						g.Blue = v
					}
				case strings.Contains(line[i+1], "green"):
					if v > g.Green {
						g.Green = v
					}
				}
			}

			power := g.Power()
			sum += power
		}

		return sum
	}

	sum := 0

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		g := Game{}
		g.ID, _ = strconv.Atoi(line[1][:len(line[1])-1])

		for i := 2; i < len(line); i += 2 {
			v, _ := strconv.Atoi(line[i])

			switch {
			case strings.Contains(line[i+1], "red"):
				g.Red = v
			case strings.Contains(line[i+1], "blue"):
				g.Blue = v
			case strings.Contains(line[i+1], "green"):
				g.Green = v
			}

			if !g.Possible(12, 14, 13) {
				break
			}
		}

		if g.Possible(12, 14, 13) {
			sum += g.ID
		}
	}

	return sum
}
