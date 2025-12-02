package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("/Users/jonathanwalter/dev/Advent-of-code/2025/day2/input.txt")
	if err != nil {
		panic(err)
	}

	input := string(content)

	ranges := strings.Split(input, ",")

	fmt.Println(part1(ranges))
	fmt.Println(part2(ranges))
}

func part1(ranges []string) int {

	invalidTotal := 0

	for _, r := range ranges {
		edges := strings.Split(r, "-")

		start, err := strconv.Atoi(edges[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(edges[1])
		if err != nil {
			panic(err)
		}

		for i := start; i <= end; i++ {
			//the pattern needs to be repeated twice, so needs to be an even size
			numStr := strconv.Itoa(i)
			runes := []rune(numStr)

			if len(runes)%2 == 0 {
				halfway := len(runes) / 2
				if string(runes[:halfway]) == string(runes[halfway:]) {
					invalidTotal = invalidTotal + i
				}
			}
		}
	}
	return invalidTotal
}

func part2(ranges []string) int {

	invalidTotal := 0

	for _, r := range ranges {
		edges := strings.Split(r, "-")

		start, err := strconv.Atoi(edges[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(edges[1])
		if err != nil {
			panic(err)
		}

		for i := start; i <= end; i++ {
			numStr := strconv.Itoa(i)
			switch {
			case checkChunksSame(splitIntoChunks(numStr, 1)) == true:
				invalidTotal += i
			case checkChunksSame(splitIntoChunks(numStr, 5)) == true:
				invalidTotal += i
			case checkChunksSame(splitIntoChunks(numStr, 4)) == true:
				invalidTotal += i
			case checkChunksSame(splitIntoChunks(numStr, 3)) == true:
				invalidTotal += i
			case checkChunksSame(splitIntoChunks(numStr, 2)) == true:
				invalidTotal += i
			}

			//switch {
			//case checkChunksSame(splitIntoChunks(numStr, 1)):
			//	invalidTotal += i
			//case len(numStr)%5 == 0 && len(numStr) != 5:
			//	if checkChunksSame(splitIntoChunks(numStr, 5)) {
			//		invalidTotal += i
			//	}
			//case len(numStr)%4 == 0 && len(numStr) != 4:
			//	if checkChunksSame(splitIntoChunks(numStr, 4)) {
			//		invalidTotal += i
			//	}
			//case len(numStr)%3 == 0 && len(numStr) != 3:
			//	if checkChunksSame(splitIntoChunks(numStr, 3)) {
			//		invalidTotal += i
			//	}
			//
			//case len(numStr)%2 == 0 && len(numStr) != 2:
			//	if checkChunksSame(splitIntoChunks(numStr, 2)) {
			//		invalidTotal += i
			//	}
			//
			//}
		}
	}
	return invalidTotal
}

func checkChunksSame(chunks []string) bool {
	if len(chunks) == 0 {
		return false
	}

	first := chunks[0]

	for i := 1; i < len(chunks); i++ {
		if chunks[i] != first {
			return false
		}
	}
	return true
}

func splitIntoChunks(s string, chunkSize int) []string {
	if chunkSize >= len(s) {
		return []string{}
	}

	var chunks []string

	runes := []rune(s)

	for i := 0; i < len(runes); i += chunkSize {
		// Define the end index of the current chunk.
		end := i + chunkSize

		// Slice the runes array from the start index (i) to the calculated end index.
		chunks = append(chunks, string(runes[i:end]))
	}

	return chunks
}
