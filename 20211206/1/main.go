package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseLines(filePath string, parse func(string) (string, bool)) ([]string, error) {
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

func main() {
    if(len(os.Args) != 2) {
        panic("missing argument")
    }

    lines, err := parseLines(os.Args[1], func(s string) (string, bool) {
        return s, true
    })
    if err != nil {
        fmt.Println("Error while parsing file", err)
        return
    }

    numbers := strings.Split(lines[0], ",")
    var list = [9]int{}
    for _, num := range numbers {
        n,_ := strconv.Atoi(num)
        list[n]++
    }
    //fmt.Println(list)
    for i:=0;i<256;i++ {
        var add [9]int
        add[8] += list[0]
        add[6] += list[0]
        for i:=0; i<8;i++ {
            add[i] += list[i+1]
        }
        list = add
    }
    //fmt.Println(list)
    res := 0
	for i := 0; i < 9; i++ {
		res += list[i]
	}
	fmt.Println(res)
        //     //var addition = []int{}
    //     for _, num := range numbers {
    //         n := strconv.Atoi(num)
    //         if n == 0 {
    //             numbers = append(numbers, "8")
    //         } else if n
    //     }
    // }
}