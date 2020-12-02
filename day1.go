package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readInputFile() []int {
	file, err := os.Open("day1.in")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var numbers []int

	for scanner.Scan() {
		line := scanner.Text()
		n, _ := strconv.Atoi(line)
		numbers = append(numbers, n)
	}

	return numbers
}

// n, find two numbers whose sum = 2020 - n

func twoSum(numbers []int, target int) int {
	seen := make(map[int]int)

	for _, n := range numbers {
		// Have we seen the required number?
		if m, ok := seen[target-n]; ok {
			return n * m
		}

		seen[n] = n
	}

	return -1
}

func threeSum(numbers []int, target int) int {
	for _, n := range numbers {
		m := twoSum(numbers, target-n)
		if m != -1 {
			return n * m
		}
	}

	return -1
}

func main() {
	numbers := readInputFile()

	fmt.Println(twoSum(numbers, 2020))
	fmt.Println(threeSum(numbers, 2020))
}
