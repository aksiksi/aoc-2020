package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type passport struct {
	byr bool // 1920-2002
	iyr bool // 2010-2020
	eyr bool // 2020-2030
	hgt bool // (150-193)cm OR (59-76)in
	hcl bool // #[\w\d] * 6
	ecl bool // one of: amb blu brn gry grn hzl oth.
	pid bool // 9 digits
}

func checkFieldExists(entry string, p *passport) {
	t := strings.Split(entry, ":")[0]

	switch t {
	case "byr":
		p.byr = true
	case "iyr":
		p.iyr = true
	case "eyr":
		p.eyr = true
	case "hgt":
		p.hgt = true
	case "hcl":
		p.hcl = true
	case "ecl":
		p.ecl = true
	case "pid":
		p.pid = true
	case "cid":
		break
	default:
		panic("not found")
	}
}

func checkFieldValid(entry string, p *passport) {
	s := strings.Split(entry, ":")
	t, v := s[0], s[1]

	switch t {
	case "byr":
		i, err := strconv.Atoi(v)
		if err != nil {
			return
		}

		p.byr = i >= 1920 && i <= 2002
	case "iyr":
		i, err := strconv.Atoi(v)
		if err != nil {
			return
		}

		p.iyr = i >= 2010 && i <= 2020
	case "eyr":
		i, err := strconv.Atoi(v)
		if err != nil {
			return
		}

		p.eyr = i >= 2020 && i <= 2030
	case "hgt":
		re := regexp.MustCompile(`^(\d+)(\w\w)$`)
		m := re.FindStringSubmatch(v)
		if len(m) != 3 {
			return
		}

		hgt, unit := m[1], m[2]
		i, _ := strconv.Atoi(hgt)

		switch unit {
		case "cm":
			p.hgt = i >= 150 && i <= 193
		case "in":
			p.hgt = i >= 59 && i <= 76
		}
	case "hcl":
		re := regexp.MustCompile(`^#[0-9a-f]+$`)
		p.hcl = re.MatchString(v)
	case "ecl":
		switch v {
		case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
			p.ecl = true
		}
	case "pid":
		re := regexp.MustCompile(`^\d+$`)
		p.pid = len(v) == 9 && re.MatchString(v)
	case "cid":
		break
	default:
		panic("not found")
	}
}

func solution(part int) []passport {
	file, err := os.Open("day4.in")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var data []passport

	curr := passport{}
	lineNo := 0

	for scanner.Scan() {
		line := scanner.Text()
		lineNo++

		if strings.Trim(line, " ") == "" {
			data = append(data, curr)
			curr = passport{}
		} else {
			for _, entry := range strings.Split(line, " ") {
				if part == 1 {
					checkFieldExists(entry, &curr)
				} else {
					checkFieldValid(entry, &curr)
				}
			}
		}
	}

	return data
}

func main() {
	part1 := solution(1)
	numValid := 0
	for _, p := range part1 {
		if p.eyr && p.hgt && p.pid && p.ecl && p.byr && p.hcl && p.iyr {
			numValid++
		}
	}
	fmt.Println(numValid)

	part2 := solution(2)
	numValid = 0
	for _, p := range part2 {
		if p.eyr && p.hgt && p.pid && p.ecl && p.byr && p.hcl && p.iyr {
			numValid++
		}
	}
	fmt.Println(numValid)
}
