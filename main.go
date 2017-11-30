//////////////////////////////////////////////

//Name: Brian Xu
//Class: Section 2
//Honor Statement: Brian Xu
//Date: Nov 30

//Name: Samat Davletshin
//Class: Section 2
//Honor Statement: Samat Davletshin
//Date: Nov 30

//Description: Sudoku solver in Go language

//////////////////////////////////////////////


package main

import (
	"fmt"
	"strconv"
	"time"
	"io/ioutil"
)

func main() {
	var str	string
	str = ""

	//read the file from specified address
	rawBoard:=readFile(str)

	//parse the text board into a convenient format for each digit
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

	//start the time count to see how quick the solver is
	start := time.Now()
	if solve(board) {
	} else{
		fmt.Printf("No solution found")
	}
	totalTime := time.Since(start)
	fmt.Printf("\nSolution took %s", totalTime)
}

//reading the sudoku board from the specified address
func readFile(rawBoard string) string{
fmt.Println("Enter the file location: ")
	var fileName string
	fmt.Scanln(&fileName)

	//read the file
	txt, err:= ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("File does not exist, retry.")
		//Asks for another input
		return readFile(rawBoard)
	} else {
		for i:=0; i<len(txt); i++{
			if _, err := strconv.Atoi(string(txt[i])); err == nil{
				rawBoard = rawBoard + string(txt[i])
			}
		}
	}
	//return original given puzzle
	return rawBoard
}


//sudoku solver function
func solve(board [9][9]int) bool {

	//check if there are any zero cells left
	a,b,c := findZero(board)

	if !a {
		//If no zeroes, then we're done!
		fmt.Printf("Solved board: ")
		fmt.Println()
		print(board)
		return true
	}

	//if there are zeros left, keep solving
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

//print the board in the pretty format
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

//returns true and the location of the first found zero cell if there is one
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

	//check row
	for i := 0; i < 9; i++ {
		if i != row && board[i][column] == number {
			return false
		}
	}

	//check column
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
