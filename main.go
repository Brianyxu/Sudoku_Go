package main

import (
	"fmt"
	"strconv"
	"time"
	"io/ioutil"
)

func main() {
	rawBoard:=readFile()
	var board [9][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			digit, err := strconv.Atoi(string(rawBoard[i*9+j]))
			if err != nil {
				fmt.Printf("error in conversion")
			}
			board[i][j] = digit
		}
	}
	fmt.Printf("Original board: ")
	fmt.Println()
	print(board)
	start := time.Now()
	if solve(board) {
	} else{
	fmt.Printf("No solution found")
	}
	totalTime := time.Since(start)
	fmt.Printf("\nSolution took %s", totalTime)
}


func readFile() string{
fmt.Println("Enter the file location: ")
	var rawBoard string
	var fileName string
	fmt.Scanln(&fileName)
	txt, err:= ioutil.ReadFile(fileName)
	if err != nil {
	fmt.Println("File does not exist, retry.")
	//Asks for another input
	readFile()
	}else{
	for i:=0;i<len(txt);i++{
	if _, err := strconv.Atoi(string(txt[i])); err == nil{
	rawBoard=rawBoard+string(txt[i])
	}
	}
	}
	return rawBoard
}



func solve(board [9][9]int) bool {
	a,b,c := findZero(board)
	if !a {
		//If no zeroes, then we're done!
		fmt.Printf("Solved board: ")
		fmt.Println()
		print(board)
		return true
	}
	for number := 1; number < 10; number++ {
		if valid(board, b, c, number) {
			board[b][c] = number
			//If successful, returns true.
			if solve(board) {
				return true
			}
			//Change it back to 0 if it fails.
			board[b][c] = 0
		}
	}
	return false
}

func print(board [9][9]int) {
	for i := 0; i < 9; i++ {
            for j := 0; j < 9; j++ {
                fmt.Printf(strconv.Itoa(board[i][j]))
		fmt.Printf(" ")
                if j % 9 == 8 && (i == 2 || i == 5) {
                    fmt.Println()
			fmt.Printf("------+-------+------")
			fmt.Println()
                } else if j % 9 == 8 {
                    fmt.Println()
                } else if j % 3 == 2 {
			fmt.Print("| ")
                }
            }
        }
}

func findZero(board [9][9]int) (bool, int, int) {
	//Checks row
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			if board[row][column] == 0 {
				return true, row, column
			}
		}
	}
	return false, 0, 0
}

//Checks column
func valid(board [9][9]int, row int, column int, number int) bool {
	for i := 0; i < 9; i++ {
		if i != row && board[i][column] == number {
			return false
		}
	}
	for j := 0; j < 9; j++ {
		if j != column && board[row][j] == number {
			return false
		}
	}

	//Checks square
	var a int
	var b int
	a = row / 3
	b = column / 3
	var a1 int
	var b1 int
	for a1 = 3 * a; a1 < 3*a+3; a1++ {
		for b1 = 3 * b; b1 < 3*b+3; b1++ {
			if number == board[a1][b1] {
				return false
			}
		}
	}
	return true
}
