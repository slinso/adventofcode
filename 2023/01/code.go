package main

import (
	"bufio"
	"strings"
	"unicode"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
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
			first := 0
			f1, fv1 := findFirstNumber(scanner.Text())
			f2, fv2 := findFirstWord(scanner.Text())

			switch {
			case f1 == -1 && f2 == -1:
				panic("nothing found")
			case f1 == -1:
				first = fv2
			case f2 == -1:
				first = fv1
			case f1 < f2:
				first = fv1
			default:
				first = fv2
			}

			last := 0
			l1, lv1 := findLastNumber(scanner.Text())
			l2, lv2 := findLastWord(scanner.Text())

			switch {
			case l1 == -1 && l2 == -1:
				panic("nothing found")
			case l1 == -1:
				last = lv2
			case l2 == -1:
				last = lv1
			case l1 > l2:
				last = lv1
			default:
				last = lv2
			}

			sum += (first * 10) + last
		}

		return sum
	}

	sum := 0

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		_, first := findFirstNumber(scanner.Text())
		_, last := findLastNumber(scanner.Text())

		sum += (first * 10) + last
	}

	return sum
}

func findFirstNumber(s string) (int, int) {
	for i, r := range s {
		if unicode.IsDigit(r) {
			return i, int(r) - '0'
		}
	}

	return -1, -1
}

func findLastNumber(s string) (int, int) {
	n := len(s)
	for n > 0 {
		n--
		if unicode.IsDigit(rune(s[n])) {
			return n, int(s[n]) - '0'
		}
	}

	return -1, -1
}

var words = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func findFirstWord(s string) (int, int) {
	found := -1
	value := -1
	for v, w := range words {
		i := strings.Index(s, w)

		if i >= 0 && (i < found || found == -1) {
			found = i
			value = v + 1
		}
	}

	if found >= 0 {
		return found, value
	}

	return -1, -1
}

func findLastWord(s string) (int, int) {
	found := -1
	value := -1
	for v, w := range words {
		i := strings.LastIndex(s, w)

		if i > found || found == -1 {
			found = i
			value = v + 1
		}
	}

	if found >= 0 {
		return found, value
	}

	return -1, -1
}
