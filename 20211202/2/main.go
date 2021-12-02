package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseLines(filePath string, parse func(string) (string,bool)) ([]string, error) {
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
type position struct {
  horisontal int
  depth  int
  aim int
}

func changePosition(line string, pos position) position {

     // Splitting the given strings
    // Using Split() function
    change := strings.Split(line, " ")
    command := change[0] //strings.Split(res1[1], " ")
    x := change[1] // := strings.Split(res2[1], ":")
    // res4 := strings.Split(line "GeeksforGeeks, geeks")
    // low, _ := strconv.Atoi(res1[0])
    x1, _ := strconv.Atoi(x)

  if command == "up" {
    // fmt.Println("up")
      pos.aim -= x1
  }

  if command == "down" {
    // fmt.Println("down")
    pos.aim += x1
  }

  if command == "forward" {
    pos.horisontal += x1
    pos.depth += x1 * pos.aim
  } 

  fmt.Println(pos)

    return pos
} 

func main() {
  if len(os.Args) != 2 {
    fmt.Println("Usage: line_parser ")
    return
  }

  lines, err := ParseLines(os.Args[1], func(s string)(string,bool){ 
    return s, true
  })
  if err != nil {
    fmt.Println("Error while parsing file", err)
    return
  }
  pos := position{
    horisontal: 0,
    depth: 0,
    aim: 0,
    }
  for _, line := range lines {
    pos = changePosition(line, pos)
  }
  fmt.Println(pos)
  fmt.Println(pos.depth * pos.horisontal)
//   for _, l := range lines {
//     fmt.Println(l)
//   }
}