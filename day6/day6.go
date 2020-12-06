package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readInput() []string {
	file, err := os.Open("day6.in")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var data []string

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return data
}

func getNumAnswers(input []string, part2 bool) int {
	questionSet := make(map[string]int)
	groupSize := 0
	numPart1 := 0
	numPart2 := 0

	for _, line := range input {
		if strings.Trim(line, " ") == "" {
			for _, v := range questionSet {
				if v == groupSize {
					// This question was answered by everyone, so record it
					numPart2++
				}
			}

			numPart1 += len(questionSet)

			groupSize = 0
			questionSet = make(map[string]int)
			continue
		}

		for _, c := range line {
			questionSet[string(c)]++
		}

		groupSize++
	}

	if part2 {
		return numPart2
	}

	return numPart1
}

func part1(input []string) int {
	return getNumAnswers(input, false)
}

func part2(input []string) int {
	return getNumAnswers(input, true)
}

func main() {
	data := readInput()
	fmt.Println(part1(data))
	fmt.Println(part2(data))
}
