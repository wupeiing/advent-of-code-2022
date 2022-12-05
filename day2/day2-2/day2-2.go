package main

import (
	"bufio"
	"fmt"
	"os"
)

var results = map[string]int{
	"A X": 3,
	"A Y": 4,
	"A Z": 8,
	"B X": 1,
	"B Y": 5,
	"B Z": 9,
	"C X": 2,
	"C Y": 6,
	"C Z": 7,
  }

func main() {

	readFile, err := os.Open("../day2.txt")
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)
  
	total := 0
    for fileScanner.Scan() {
		if fileScanner.Text() != "" {
			total += results[fileScanner.Text()]
		}
    }

	fmt.Println("Total value: ", total)
  
    readFile.Close()
}