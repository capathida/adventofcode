package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func sum(str1 string, str2 string, str3 string) int {
  i, _ := strconv.Atoi(str1)
  j, _ := strconv.Atoi(str2)
  k, _ := strconv.Atoi(str3)
  // fmt.Println(i, j , k)
  return i + j + k
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
  for i, _ := range lines {
    if i < 3 {
      continue;
    }
    // fmt.Println(lines[i-3], lines[i-2], lines[i])
    sum1 := sum(lines[i-3], lines[i-2], lines[i-1])
    sum2 := sum(lines[i-2], lines[i-1], lines[i])
    // fmt.Println(sum1, sum2)
    if sum1 < sum2  {
      total++
    }
  }
  fmt.Println(total)
}