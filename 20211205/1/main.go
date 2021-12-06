package main

import (
	"bufio"
	"fmt"
	"os"
)

func ParseLines(filePath string, parse func(string) (string, bool)) ([]string, error) {
	inputFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	var results []string
	for scanner.Scan() {
		if output, add := parse(scanner.Text()); add {
			results = append(results, output)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return results, nil
}

//1-3 a: abcde
type box struct {
	number int
	marked int
}

func mark(bb []box, i int) []box {
	//fmt.Println(bb, i)
	for _, b := range bb {
		if b.number == i {
			b.marked = 1
			fmt.Println("mark", b)
		}
	}
	return bb
}

func winner(boards [][]box) int {
	// rows
	fmt.Println("---------------------")
	for i, b := range boards {
		//fmt.Println(i, b)
		for w := 0; w < 5; w++ {
			found := true
			for e := 0; e < 5; e++ {
				bb := b[e+w*5]
				//fmt.Println("e", e, w, e+w*5, b, bb, found)
				if bb.marked == 0 {
					found = false
					//fmt.Println("g", e, w, e+w*5, b, found)
				}
				//fmt.Println("f", e, w, e+w*5, b, found)
			}
			if found == true {
				//fmt.Println("ff", found, b, i, w)
				return i
			}
		}
	}
	return -1
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: line_parser ")
		return
	}

	lines, err := ParseLines(os.Args[1], func(s string) (string, bool) {
		return s, true
	})
	if err != nil {
		fmt.Println("Error while parsing file", err)
		return
	}

	//numbers := strings.Split(lines[0], ",")
	//fmt.Println(numbers)

	//for len(lines) d

	type B = [][]int
	b := make([][]int, 10)
	for r := range b {
		b[r] = make([]int, 10)
	}
	board := B(b)

	for _, r := range lines {
		//var board = []box{}
		//for i := 2; i < 7; i++ { // change to 12 in input file or read length
		//fmt.Println("res", j, i, lines[i+j*6])
		var arr = make([]int, 4)

		fmt.Sscanf(r, "%d,%d -> %d,%d\n", &arr[0], &arr[1], &arr[2], &arr[3])
		//fmt.Println(arr, i)

		//add arr to board
		if arr[0] == arr[2] {
			if arr[1] <= arr[3] {
				for y := arr[1]; y <= arr[3]; y++ {
					fmt.Println("1num", arr[0], arr[1], arr[2], arr[3])
					board[y][arr[0]]++
				}
			} else {
				for y := arr[3]; y <= arr[1]; y++ {
					//fmt.Println("n1um", arr[0], arr[1], arr[2], arr[3])
					board[y][arr[0]]++
				}
			}
		} else if arr[1] == arr[3] {
			if arr[0] <= arr[2] {
				for x := arr[0]; x <= arr[2]; x++ {
					fmt.Println("num1", arr[0], arr[1], arr[2], arr[3], x)
					board[arr[1]][x]++
				}
			} else {
				for x := arr[2]; x <= arr[0]; x++ {
					fmt.Println("nu1m", arr[0], arr[1], arr[2], arr[3], x)
					board[arr[1]][x]++
				}
			}
		}
		sum := 0
		for _, bb := range board {
			for _, cc := range bb {
				if cc > 1 {
					sum++
				}
			}
		}
		fmt.Println(sum)
		/* 			for _, l := range arr {
				b := box{
					number: l,
					marked: 0,
				}
				board = append(board, b)
				fmt.Println("box", board)
			}
		}
		//fmt.Println("board", board)
		boards = append(boards, board) */
	}
	for _, t := range board {
		fmt.Println(t)
	}
}
