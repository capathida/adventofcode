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

    if command == "forward" {
        pos.horisontal += x1
    } else {
        if command == "up" {
            pos.depth -= x1
        } else {
            pos.depth += x1
        }
    }

    // var binary string  
    // fmt.Print("Enter Binary Number:")  
    // fmt.Scanln(&binary)  
    // output, err := strconv.ParseInt(binary, 2, 64) 

    // fmt.Println(pos)

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
  gamma := ""
  epsilon := ""
  for i := 0; i < 12; i++ { // change to 12 in input file or read length
    g := 0
    e := 0
    for _, line := range lines {
      c := line[i]
      if c == '1' {
        g++
      } else {
        e++ 
      }
      // fmt.Println(g, e)
    }
    if g > e {
      gamma += "1"
      epsilon += "0"
    } else {
      gamma += "0"
      epsilon += "1"
    }
  }
  fmt.Println(gamma, epsilon)
  gg, err := strconv.ParseInt(gamma, 2, 64) 
  ee , err := strconv.ParseInt(epsilon, 2, 64) 
  fmt.Println(gg, ee)
  fmt.Println(gg * ee)
//   for _, l := range lines {
//     fmt.Println(l)
//   }
}