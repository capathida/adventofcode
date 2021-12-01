package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// While it appears you validated the passwords correctly, they don't seem to be what the Official Toboggan Corporate Authentication System is expecting.

// The shopkeeper suddenly realizes that he just accidentally explained the password policy rules from his old job at the sled rental place down the street! The Official Toboggan Corporate Policy actually works a little differently.

// Each policy actually describes two positions in the password, where 1 means the first character, 2 means the second character, and so on. (Be careful; Toboggan Corporate Policies have no concept of "index zero"!) Exactly one of these positions must contain the given letter. Other occurrences of the letter are irrelevant for the purposes of policy enforcement.

// Given the same example list from above:

// 1-3 a: abcde is valid: position 1 contains a and position 3 does not.
// 1-3 b: cdefg is invalid: neither position 1 nor position 3 contains b.
// 2-9 c: ccccccccc is invalid: both position 2 and position 9 contain c.
// How many passwords are valid according to the new interpretation of the policies?

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