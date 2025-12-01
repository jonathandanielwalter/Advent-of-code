package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("/Users/jonathanwalter/dev/Advent-of-code/2025/day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}
	fmt.Println("part 1 : ", timesDialAt0(input))
	fmt.Println("part 2 : ", timesDialClicks(input))
}

func timesDialClicks(inputs []string) int {
	point := 50
	count := 0

	var clicks int
	for _, input := range inputs {
		runes := []rune(input)

		direction := string(runes[0])

		value, err := strconv.Atoi(string(runes[1:]))

		if err != nil {
			panic(err)
		}

		point, clicks = turnDialForClicks(point, value, direction)

		count = count + clicks
	}
	return count
}

func timesDialAt0(inputs []string) int {
	point := 50
	count := 0
	for _, input := range inputs {
		runes := []rune(input)

		direction := string(runes[0])

		value, err := strconv.Atoi(string(runes[1:]))

		if err != nil {
			panic(err)
		}

		point = turnDial(point, value, direction)
		if point == 0 {
			count++
		}
	}
	return count
}

func turnDial(pos, count int, dir string) int {
	for i := 0; i < count; i++ {
		switch dir {
		case "R": // go up
			pos++
			if pos == 100 {
				pos = 0
			}
		case "L":
			pos--
			if pos == -1 {
				pos = 99
			}
		}
	}

	return pos
}

func turnDialForClicks(pos, count int, dir string) (int, int) {
	clicks := 0

	for i := 0; i < count; i++ {
		switch dir {
		case "R": // go up
			pos++
			if pos == 100 {
				pos = 0
			}
			if pos == 0 {
				clicks++
			}

		case "L":
			pos--
			if pos == -1 {
				pos = 99
			}

			if pos == 0 {
				clicks++
			}
		}
	}

	return pos, clicks
}
