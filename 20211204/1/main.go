package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	//fmt.Println("---------------------")
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

	numbers := strings.Split(lines[0], ",")
	fmt.Println(numbers)

	//for len(lines) d

	res := 0
	var boards = [][]box{}
	size := (len(lines) - 1) / 6
	fmt.Println(size)
	for j := 0; j < size; j++ {
		var board = []box{}
		for i := 2; i < 7; i++ { // change to 12 in input file or read length
			//fmt.Println("res", j, i, lines[i+j*6])
			var arr = make([]int, 5)

			fmt.Sscanf(lines[i+j*6], "%d %d %d %d %d\n", &arr[0], &arr[1], &arr[2], &arr[3], &arr[4])
			for _, l := range arr {
				b := box{
					number: l,
					marked: 0,
				}
				board = append(board, b)
				//fmt.Println("box", board)
			}
		}
		//fmt.Println("board", board)
		boards = append(boards, board)
	}

	//runGame(&boards, numbers)
	//for _, l := range numbers {
	found := -1
	for q := 0; q < len(numbers); q++ {
		l := numbers[q]
		//fmt.Println("bla", len(numbers), l)
		//		for _, b := range boards {
		for w := 0; w < len(boards); w++ {
			b := boards[w]
			v, _ := strconv.Atoi(l)
			for e := 0; e < len(b); e++ {
				//			for _, bb := range b {
				bb := b[e]
				if bb.number == v {
					bb.marked = 1
					//fmt.Println("mark", bb, v, bb.marked, bb.number)
				}
				//fmt.Println("hej12", bb, w)
				b[e] = bb
			}
			boards[w] = b
			//fmt.Println("hej", b, w)
		}
		//fmt.Println("hej123", l)
		/* 		for i, b := range boards {
			fmt.Println("bb", i, b)
		} */
		found = winner(boards)
		if found != -1 {
			fmt.Println("last number", l)
			res, _ = strconv.Atoi(l)
			break
		}
		//fmt.Println("hesjhkshdf", found)
	}
	fmt.Println("matched board", found)
	fmt.Println(boards[found])
	sum := 0
	for _, b := range boards[found] {
		if b.marked == 0 {
			sum += b.number
		}
	}
	fmt.Println(sum)
	fmt.Println("Result: ", sum*res)

}
