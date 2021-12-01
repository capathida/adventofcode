package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ParseLines(filePath string, parse func(string) (string,bool)) ([]int, error) {
  inputFile, err := os.Open(filePath)
  if err != nil {
    return nil, err
  }
  defer inputFile.Close()

  scanner := bufio.NewScanner(inputFile)
  var results []int
  for scanner.Scan() {
    if output, add := parse(scanner.Text()); add {
        i, _ := strconv.Atoi(output)
        results = append(results, i)
    }
  }
  if err := scanner.Err(); err != nil {
    return nil, err
  }
  return results, nil
}

func find_number(list []int) int {
    for i, res := range list {
        // Iterate using for loop
        for e := (i+1); e < len(list)-1; e++ {
            for s := (e+1); s < len(list)-2; s++ {
                // fmt.Println(result)
                if res + list[e] + list[s] == 2020 {
                    fmt.Println(res + list[e] + list[s])
                    return res*list[e]*list[s]
                }
            }
            // // fmt.Println(result)
            // if res + list[e] == 2020 {
            //     return res*list[e]
            // }
        }
    }
    return 0
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

  num := find_number(lines)
  fmt.Println(num)

//   for _, l := range lines {
//     fmt.Println(l)
//   }
}