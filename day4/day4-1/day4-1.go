package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {

	readFile, err := os.Open("../day4.txt")
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)
  
	total := 0
    for fileScanner.Scan() {
		secs := strings.Split(fileScanner.Text(), ",")
		seca := secs[0]
		secb := secs[1]
		secase := strings.Split(seca, "-")
		secas, _ := strconv.Atoi(secase[0])
		secae, _ := strconv.Atoi(secase[1])
		secbse := strings.Split(secb, "-")
		secbs, _ := strconv.Atoi(secbse[0])
		secbe, _ := strconv.Atoi(secbse[1])
		if secbs <= secas && secbe >= secae {
			total += 1
		} else if secas <= secbs && secae >= secbe {
			total += 1
		}
    }
    fmt.Println(total)
    readFile.Close()
}