package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readInputFile() []string {
	file, err := os.Open("day3.in")
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

func getNextPos(x, y, right, down, width int) (int, int) {
	return (x + right) % width, y + down
}

func countTrees(right, down int, grid []string) int {
	numTrees := 0
	width := len(grid[0])

	x, y := 0, 0
	for {
		x, y = getNextPos(x, y, right, down, width)
		if y >= len(grid) {
			break
		}

		if grid[y][x] == '#' {
			numTrees++
		}
	}

	return numTrees
}

func countTreesMulti(grid []string) int {
	options := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	prod := 1

	for _, opt := range options {
		count := countTrees(opt[0], opt[1], grid)
		prod *= count
	}

	return prod
}

func main() {
	data := readInputFile()

	fmt.Println(countTrees(3, 1, data))
	fmt.Println(countTreesMulti(data))
}
