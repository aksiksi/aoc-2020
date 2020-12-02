package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type passwordConstraint struct {
	c   byte
	min int
	max int
}

type passwordEntry struct {
	constraint passwordConstraint
	password   string
}

func readInputFile() []passwordEntry {
	file, err := os.Open("day2.in")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var data []passwordEntry

	for scanner.Scan() {
		line := strings.Replace(scanner.Text(), ":", "", 1)
		sp := strings.Split(line, " ")

		cRange := strings.Split(sp[0], "-")
		min, _ := strconv.Atoi(cRange[0])
		max, _ := strconv.Atoi(cRange[1])

		c := sp[1][0]
		password := sp[2]

		pw := passwordEntry{
			password: password,
			constraint: passwordConstraint{
				c:   c,
				min: min,
				max: max,
			},
		}

		data = append(data, pw)
	}

	return data
}

func findNumValidPasswords(data []passwordEntry) int {
	numValid := 0

	for _, entry := range data {
		constraint := entry.constraint
		password := entry.password
		count := 0

		for i := range password {
			if password[i] == constraint.c {
				count++
			}
		}

		if count >= constraint.min && count <= constraint.max {
			numValid++
		}
	}

	return numValid
}

func findNumValidPasswords2(data []passwordEntry) int {
	numValid := 0

	for _, entry := range data {
		constraint := entry.constraint
		password := entry.password

		// For part 2 of the problem, the min and max represent indices
		// We need at most one of the characters at these positions to
		// match to declare the password as "valid"
		leftMatch := password[constraint.min-1] == constraint.c
		rightMatch := password[constraint.max-1] == constraint.c

		if (leftMatch && !rightMatch) || (!leftMatch && rightMatch) {
			numValid++
		}
	}

	return numValid
}

func main() {
	data := readInputFile()

	fmt.Println(findNumValidPasswords(data))
	fmt.Println(findNumValidPasswords2(data))
}
