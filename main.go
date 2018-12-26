package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

var board [9][9]int
var recursion int64

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please specify filename!")
		return
	}

	jsonfile := os.Args[1]

	//fmt.Println(jsonfile)
	result := loadInputFile(jsonfile)
	initBoard(result)
	printBoard()
	solveBoard()
	printBoard()

	//cp := "\x86\x8C\x89\xD1\x8C"
	//str := cp1252ToUTF8(cp)
	//fmt.Printf("%q\n", cp)

}

func solveBoard() bool {
	recursion++

	found, y, x := findNextEmptyItem()
	if !found {
		fmt.Println(recursion)
		return true
	}

	//printBoard()

	for v := 1; v <= 9; v++ {

		//fmt.Println("  - ", y, x, v)
		if isValid(y, x, v) {
			board[y][x] = v
			//fmt.Print("Press 'Enter' to continue...")
			//bufio.NewReader(os.Stdin).ReadBytes('\n')
			//fmt.Println(y, x)

			if solveBoard() {
				return true
			}
		}
	}

	//fmt.Println("Dead end !!!")
	board[y][x] = 0
	return false

}

func isValid(y int, x int, value int) bool {
	// check x
	var h, v, b []int

	for i := 0; i < 9; i++ {
		if board[y][i] > 0 {
			h = append(h, board[y][i])
		}

		if board[i][x] > 0 {
			v = append(v, board[i][x])
		}
	}

	var xf, xt, yf, yt int
	if x <= 2 {
		xf = 0
		xt = 2
	} else if x > 2 && x <= 5 {
		xf = 3
		xt = 5
	} else if x > 5 {
		xf = 6
		xt = 8
	}

	if y <= 2 {
		yf = 0
		yt = 2
	} else if y > 2 && y <= 5 {
		yf = 3
		yt = 5
	} else if y > 5 {
		yf = 6
		yt = 8
	}

	for by := yf; by <= yt; by++ {
		for bx := xf; bx <= xt; bx++ {
			if board[by][bx] > 0 {
				b = append(b, board[by][bx])
			}
		}
	}

	//fmt.Println(value, v, h, b, isValidValue(v, value), isValidValue(h, value), isValidValue(b, value))

	//fmt.Println(value, v, h, b, isValidValue(v, value) && isValidValue(h, value) && isValidValue(b, value))

	return isValidValue(v, value) && isValidValue(h, value) && isValidValue(b, value)

}

func isValidValue(currentValues []int, value int) bool {

	ret := true
	for _, v := range currentValues {
		if v == value {
			ret = false
		}
	}

	return ret
}

func findNextEmptyItem() (bool, int, int) {

	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if board[y][x] == 0 {
				return true, y, x
			}
		}
	}

	return false, 0, 0
}

func printBoard() {
	fmt.Println("")
	for y := 0; y < 9; y++ {
		line := ""
		for x := 0; x < 9; x++ {
			if board[y][x] > 0 {
				line += " " + strconv.Itoa(board[y][x]) + " "
			} else {
				line += " . "
			}
			if x == 2 || x == 5 {
				line += "|"
			}
		}
		fmt.Println(line)
		if y == 2 || y == 5 {
			fmt.Println("----------------------------")
		}
	}
	fmt.Println("")

}

func loadInputFile(filename string) [][]int {

	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened", filename)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)

	//fmt.Println(byteValue)

	if err != nil {
		fmt.Println(err)
	}

	var result [][]int

	json.Unmarshal([]byte(byteValue), &result)

	return result

}

func initBoard(items [][]int) {
	//for y := 0; y < 9; y++ {
	//	for x := 0; y < 9; y++ {
	//		board[y][x] = 0
	//	}
	//}

	for i := 0; i < len(items); i++ {
		item := items[i]

		board[item[0]][item[1]] = item[2]
	}
}
