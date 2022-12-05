package main

import (
	"bufio"
	"fmt"
	"os"
)

var results = map[string]int{
	"A X": 4,
	"A Y": 8,
	"A Z": 3,
	"B X": 1,
	"B Y": 5,
	"B Z": 9,
	"C X": 7,
	"C Y": 2,
	"C Z": 6,
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