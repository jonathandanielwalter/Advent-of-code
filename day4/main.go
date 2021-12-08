package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	Rows    []Row
	Columns []Column
}

type Row struct {
	Numbers []string
}

type Column struct {
	Numbers []string
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var bingoCalledNumbers []string
	var allBoards []Board

	var previousLineBlank bool = false
	var previousLineNumbers bool = false
	var currentBoard Board
	for scanner.Scan() {
		row := scanner.Text()
		if strings.Contains(row, ",") {
			bingoCalledNumbers = strings.Split(row, ",")
			previousLineBlank = false
			previousLineNumbers = true
		} else if row == "" && !previousLineNumbers {
			allBoards = append(allBoards, currentBoard)
			previousLineBlank = true
			previousLineNumbers = false
		} else if row == "" {
			previousLineBlank = true
			previousLineNumbers = false
		} else {
			if previousLineBlank {
				currentBoard = createBoard(row) //create new board
				// currentBoard = currentBoard
				previousLineBlank = false

			} else {
				currentBoard = addRowAndAddToColumns(&currentBoard, row)
				//fmt.Printf("%v\n", currentBoard)
				previousLineBlank = false
			}
		}

	}
	allBoards = append(allBoards, currentBoard)

	// fmt.Printf("%v\n", allBoards[0])
	// fmt.Printf("%v\n", allBoards[1])
	// fmt.Printf("%v\n", allBoards[2])

	winner, winningNumberStr := playBingo(allBoards, bingoCalledNumbers)

	winningNumber, _ := strconv.Atoi(winningNumberStr)

	val := calculateRemainingNumbersTotal(winner) * winningNumber
	println(val)

}

func playBingo(allBoards []Board, bingoCalledNumbers []string) (Board, string) {
	for _, number := range bingoCalledNumbers {
		for _, board := range allBoards {
			board = mark(board, number)
			// println("removing", number)
			// fmt.Printf("%v\n", board)
			for _, column := range board.Columns {
				if !hasNonEmptyValues(column.Numbers) {
					return board, number
				}
			}
			for _, row := range board.Rows {
				if !hasNonEmptyValues(row.Numbers) {
					return board, number
				}
			}
		}

		//mark(allBoards[0], number)
	}
	return Board{}, ""
}

func mark(board Board, number string) Board {
	//fmt.Printf("%v\n", board)
	for _, row := range board.Rows {
		for i := 0; i <= len(row.Numbers)-1; i++ {
			// fmt.Printf("comparing %v and %v\n", row.Numbers[i], number)
			if row.Numbers[i] == number {
				row.Numbers = remove(row.Numbers, i)
				// fmt.Printf("Removed %v from %v\n", number, row.Numbers)
			}
		}

		for _, column := range board.Columns {
			for i := len(column.Numbers) - 1; i >= 0; i-- {
				if column.Numbers[i] == number {
					column.Numbers = remove(column.Numbers, i)
				}
			}
			//println(len(column.Numbers))
		}
	}

	//fmt.Printf("%v\n", board)
	return board
}

func remove(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:]) // Shift a[i+1:] left one index.
	slice[len(slice)-1] = ""     // Erase last element (write zero value).
	return slice[:len(slice)-1]
}

func createBoard(initialRow string) Board {
	//println("creating board")
	var board Board = Board{}
	numbers := strings.Fields(initialRow)
	var row Row = Row{}
	row.Numbers = numbers
	board.Rows = append(board.Rows, row)

	for _, number := range numbers {
		var column = Column{}
		column.Numbers = append(column.Numbers, number)
		board.Columns = append(board.Columns, column)
	}

	//fmt.Printf("%v", board)
	return board
}

func addRowAndAddToColumns(board *Board, newRow string) Board {
	numbers := strings.Fields(newRow)
	var row Row = Row{}
	row.Numbers = numbers
	board.Rows = append(board.Rows, row)

	for i, number := range numbers {
		board.Columns[i].Numbers = append(board.Columns[i].Numbers, number)
	}

	return *board
	//fmt.Printf("%v\n", board)
}

func hasNonEmptyValues(values []string) bool {
	for _, str := range values {
		if str != "" {
			return true
		}
	}
	return false
}

func calculateRemainingNumbersTotal(board Board) int {
	var total int

	for _, row := range board.Rows {
		for _, numberStr := range row.Numbers {
			number, _ := strconv.Atoi(numberStr)
			total += number
		}
	}
	println(total)
	return total
}