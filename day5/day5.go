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

// Convert the seat string to an integer ID
// Each seat string is basically a concatenation of two
// binary numbers:
//   1. Row: bits 9-3
//   2. Col: bits 2-0
func getID(seat string) int {
	n := 0
	pos := len(seat) - 1

	// Convert seat string to bits
	for _, c := range seat {
		if c == 'B' || c == 'R' {
			n |= (1 << pos)
		}

		pos--
	}

	// No manipulation is needed because the row is _already_
	// multiplied by 8 (shifted by 3 bits)
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
