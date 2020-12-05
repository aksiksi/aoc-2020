package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func readInput() []string {
	file, err := os.Open("day5.in")
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

func getID(seat string) int {
	n := 0

	rowPart := seat[:7]
	colPart := seat[7:]
	rowPos := 6
	colPos := 2

	row := 0
	col := 0

	for _, c := range rowPart {
		if c == 'B' {
			row |= (1 << rowPos)
		}

		rowPos--
	}

	for _, c := range colPart {
		if c == 'R' {
			col |= (1 << colPos)
		}

		colPos--
	}

	n = row*8 + col

	return n
}

func findMaxID(seats []string) int {
	maxID := 0

	for _, seat := range seats {
		n := getID(seat)
		if n > maxID {
			maxID = n
		}
	}

	return maxID
}

func findMySeat(seats []string) int {
	mySeatID := -1

	seatIDs := make([]int, len(seats))

	for i, seat := range seats {
		seatIDs[i] = getID(seat)
	}

	// Sort the seat IDs
	sort.Ints(seatIDs)

	// Look for a gap between seat IDs
	for i := 0; i < len(seatIDs)-1; i++ {
		if seatIDs[i+1] != seatIDs[i]+1 {
			mySeatID = seatIDs[i] + 1
			break
		}
	}

	return mySeatID
}

func main() {
	data := readInput()
	fmt.Println(findMaxID(data))
	fmt.Println(findMySeat(data))
}
