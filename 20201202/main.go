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
type passwordpolicy struct {
  lowest int
  highest  int
  chr string
  password string
}

func ispasswordok(pw passwordpolicy) bool {
  found := strings.Count(pw.password, pw.chr)
  if found <= pw.highest && found >= pw.lowest {
    return true
  }
  return false
}

func get_password_policy(line string) passwordpolicy {

     // Splitting the given strings
    // Using Split() function
    res1 := strings.Split(line, "-")
    res2 := strings.Split(res1[1], " ")
    res3 := strings.Split(res2[1], ":")
    res4 := res2[2]
    // res4 := strings.Split(line "GeeksforGeeks, geeks")
    low, _ := strconv.Atoi(res1[0])
    high, _ := strconv.Atoi(res2[0])

    p := passwordpolicy{
      lowest: low,
      highest: high,
      chr: res3[0],
      password: res4}
    fmt.Println(p.password)
  return p
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
  total := 0
  for _, line := range lines {
    p := get_password_policy(line)
    if ispasswordok(p) {
      total++
    }
  }
  fmt.Println(total)
//   for _, l := range lines {
//     fmt.Println(l)
//   }
}