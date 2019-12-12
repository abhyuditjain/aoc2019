/*
--- Day 4: Secure Container ---

You arrive at the Venus fuel depot only to discover it's protected by a password. The Elves had written the password on a sticky note, but someone threw it out.

However, they do remember a few key facts about the password:

    It is a six-digit number.
    The value is within the range given in your puzzle input.
    Two adjacent digits are the same (like 22 in 122345).
    Going from left to right, the digits never decrease; they only ever increase or stay the same (like 111123 or 135679).

Other than the range rule, the following are true:

    111111 meets these criteria (double 11, never decreases).
    223450 does not meet these criteria (decreasing pair of digits 50).
    123789 does not meet these criteria (no double).

How many different passwords within the range given in your puzzle input meet these criteria?

--- Part Two ---

An Elf just remembered one more important detail: the two adjacent matching digits are not part of a larger group of matching digits.

Given this additional criterion, but still ignoring the range rule, the following are now true:

    112233 meets these criteria because the digits never decrease and all repeated digits are exactly two digits long.
    123444 no longer meets the criteria (the repeated 44 is part of a larger group of 444).
    111122 meets the criteria (even though 1 is repeated more than twice, it still contains a double 22).

How many different passwords within the range given in your puzzle input meet all of the criteria?

*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	lines := readLines("input.txt")
	min, max := parseInput(lines[0])
	fmt.Println(totalWithMatchingDigits(min, max))
	fmt.Println(totalWithIsolatedMatchingDigits(min, max))
}

func totalWithMatchingDigits(min, max int) int {
	total := 0
	for current := min; current <= max; current++ {
		password := fmt.Sprintf("%06d", current)

		if hasMatchingDigits(password) && !decreasing(password) {
			total++
		}
	}
	return total
}

func totalWithIsolatedMatchingDigits(min, max int) int {
	total := 0
	for current := min; current <= max; current++ {
		password := fmt.Sprintf("%06d", current)

		if hasIsolatedMatchingDigits(password) && !decreasing(password) {
			total++
		}
	}
	return total
}

func decreasing(s string) bool {
	bytes := []byte(s)
	for i := 1; i < len(bytes); i++ {
		if bytes[i] < bytes[i-1] {
			return true
		}
	}
	return false
}

func hasMatchingDigits(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			return true
		}
	}
	return false
}

func hasIsolatedMatchingDigits(s string) bool {
	seen := make(map[rune]int, 6)

	for _, r := range s {
		seen[r]++
	}

	for _, duplicate := range seen {
		if duplicate == 2 {
			return true
		}
	}
	return false
}

func readLines(filename string) []string {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func parseInput(str string) (int, int) {
	regex := regexp.MustCompile(`^(\d{6})-(\d{6})$`)
	match := regex.FindStringSubmatch(str)
	return toInt(match[1]), toInt(match[2])
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func toInt(str string) int {
	val, err := strconv.Atoi(str)
	check(err)
	return val
}
