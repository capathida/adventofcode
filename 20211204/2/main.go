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

func winner(boards [][]box, m map[int]bool) []int {
	// rows
	//fmt.Println("---------------------")
	var list = []int{}
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
				if m[i] != true {
					list = append(list, i)
				}
			}
		}
	}
	for i, b := range boards {
		//fmt.Println(i, b)
		for w := 0; w < 5; w++ {
			found := true
			for e := 0; e < 5; e++ {
				bb := b[e*5+w]
				//fmt.Println("e", e, w, e*5+w, b, bb, found)
				if bb.marked == 0 {
					found = false
					//fmt.Println("g", e, w, e+w*5, b, found)
				}
				//fmt.Println("f", e, w, e+w*5, b, found)
			}
			if found == true {
				//fmt.Println("ff", found, b, i, w)
				if m[i] != true {
					//fmt.Println("ff", found, b, i, w)
					list = append(list, i)
				}
			}
		}
	}
	return list
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
	fmt.Println(numbers, "len", len(numbers))

	//for len(lines) d

	var boards = [][]box{}
	size := (len(lines) - 1) / 6
	fmt.Println("Number of boards", size)
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
	res := 0
	foundLast := 0
	var m map[int]bool
	m = make(map[int]bool)
	//fmt.Println(boards)
	for q := 0; q < len(numbers); q++ {
		l := numbers[q]
		//		for _, b := range boards {
		for w := 0; w < len(boards); w++ {
			b := boards[w]
			v, _ := strconv.Atoi(l)
			for e := 0; e < len(b); e++ {
				//			for _, bb := range b {
				bb := b[e]
				if bb.number == v {
					bb.marked = 1
					//fmt.Println("mark", bb, w, v, bb.marked, bb.number)
				}
				//fmt.Println("hej12", bb, w)
				b[e] = bb
			}
			boards[w] = b
			fmt.Println(b)
			//fmt.Println("hej", b, w)
		}
		//fmt.Println("hej123", l)
		/* 		for i, b := range boards {
			fmt.Println("bb", i, b)
		} */

		list := winner(boards, m)
		fmt.Println(m, len(m), l, q)
		for _, found := range list {
			m[found] = true
			// fmt.Println("hesjhkshdf", found, m)
			foundLast = found
			if len(m) == len(boards) {
				fmt.Println("last number", l, m)
				res, _ = strconv.Atoi(l)
				break
			}
		}
		//fmt.Println("hesjhkshdf", found)
	}
	//if q == len(numbers)
	fmt.Println("matched board", foundLast)
	fmt.Println(boards[foundLast])
	sum := 0
	for _, b := range boards[foundLast] {
		if b.marked == 0 {
			sum += b.number
		}
	}
	fmt.Println(sum)

	//fmt.Println("hesjhkshdf", found)
	fmt.Println("Result: ", sum*res)
}
