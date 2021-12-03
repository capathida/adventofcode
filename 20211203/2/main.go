package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
  // gamma := ""
  // epsilon := ""

  rest := []string{}
  for i := 0; i < 12; i++ { // change to 12 in input file or read length
    g := 0
    e := 0
    // fmt.Println(rest, lines)
    for _, line1 := range lines {
      c := line1[i]
      if c == '1' {
        g++
      } else {
        e++ 
      }
      // fmt.Println(g, e)
    }

    // fmt.Println(g, e)
    for _, line := range lines {
      if g > e && line[i] == '1' {
        rest = append(rest, line)
        // fmt.Println(i, line[i])
        // 11110, 10110, 10111, 10101, 11100, 10000, and 11001.
        // 10110, 10111, 10101, and 10000.
      } else {
        if g < e && line[i] == '0' {
          // fmt.Println(i, line[i], rest)
          rest = append(rest, line)
        } else {
          if g == e && line[i] == '1' {
            rest = append(rest, line)
          }
        }
      }
      if len(lines) == 1 {
        fmt.Println(rest, lines)
        gg, _ := strconv.ParseInt(lines[0], 2, 64) 
        fmt.Println(gg)
      }
    }
    // fmt.Println(i, rest)
    lines = rest
    rest = nil
    if len(lines) == 1 {
      fmt.Println(rest, lines)
      gg, _ := strconv.ParseInt(lines[0], 2, 64) 
      fmt.Println(gg)
    }
          //fmt.Println(rest, lines)
  }

  lines, err = ParseLines(os.Args[1], func(s string)(string,bool){ 
    return s, true
  })
  if err != nil {
    fmt.Println("Error while parsing file", err)
    return
  }

  rest = nil
  for i := 0; i < 12; i++ { // change to 12 in input file or read length
    g := 0
    e := 0
    // fmt.Println(rest, lines)
    for _, line1 := range lines {
      c := line1[i]
      if c == '1' {
        g++
      } else {
        e++ 
      }
      // fmt.Println(g, e)
    }

    // fmt.Println(g, e)
    for _, line := range lines {
      if g < e && line[i] == '1' {
        rest = append(rest, line)
        // fmt.Println(i, line[i])
        // 11110, 10110, 10111, 10101, 11100, 10000, and 11001.
        // 10110, 10111, 10101, and 10000.
      } else {
        if g > e && line[i] == '0' {
          // fmt.Println(i, line[i], rest)
          rest = append(rest, line)
        } else {
          if g == e && line[i] == '0' {
            rest = append(rest, line)
          }
        }
      }
      if len(lines) == 1 {
        fmt.Println(rest, lines)
        gg, _ := strconv.ParseInt(lines[0], 2, 64) 
        fmt.Println(gg)
      }
    }
    lines = rest
    rest = nil
    if len(lines) == 1 {
      fmt.Println(rest, lines)
      gg, _ := strconv.ParseInt(lines[0], 2, 64) 
      fmt.Println(gg)
    }
  }
}