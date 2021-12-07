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
    //fmt.Println(numbers)
    var list = []int{}
    for _, num := range numbers {
        n,_ := strconv.Atoi(num)
        list = append(list, n)
    }
    //fmt.Println(list)
    biggest:=0
    for _, v := range list {
        if v > biggest {
            biggest = v
        }
    }
    //fmt.Println(biggest)
    const MaxUint = ^uint(0) 
    minTotalFuel := int(MaxUint >> 1) 
    pos := 0
    for i:=0;i<biggest;i++ {
        totalFuel :=0
        for _, x := range list {
            if x > i {
                y := x-i
                totalFuel += y * (y + 1) / 2
                //fmt.Println("lg",i, x,totalFuel)
            } else if x < i {
                y := (i-x)
                totalFuel += y * (y + 1) / 2
                //fmt.Println("sm", i,x,totalFuel)
            }
        }
        if (minTotalFuel > totalFuel) {
            minTotalFuel = totalFuel
            pos = i
            //fmt.Println("loop", i,totalFuel, minTotalFuel)
        }
    }

    fmt.Println(pos, minTotalFuel)
}