package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("C:/Users/jonat/playground/Advent-of-code/2025/day1/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

}
