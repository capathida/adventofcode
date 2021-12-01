package main

import (
	"bufio"
	"fmt"
	"os"
)

func parseLines(filePath string, parse func(string) (string,bool)) ([]string, error) {
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

  lines, err := parseLines(os.Args[1], func(s string)(string,bool) { 
    return s, true
  })
  if err != nil {
    fmt.Println("Error while parsing file", err)
    return
  }

  
  total := 0
  for i := range lines {
    if i < 1 {
      continue;
    }
    // fmt.Println(sum1, sum2)
    if lines[i] > lines[i-1] {
      total++
    }
  }

  fmt.Println(total)
}