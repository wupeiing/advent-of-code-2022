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
  
	total_count := 0
	total_no_olap := 0
    for fileScanner.Scan() {
		total_count = total_count + 1
		secs := strings.Split(fileScanner.Text(), ",")
		seca := secs[0]
		secb := secs[1]
		secase := strings.Split(seca, "-")
		secas, _ := strconv.Atoi(secase[0])
		secae, _ := strconv.Atoi(secase[1])
		secbse := strings.Split(secb, "-")
		secbs, _ := strconv.Atoi(secbse[0])
		secbe, _ := strconv.Atoi(secbse[1])
		if secbs > secae || secas > secbe {
			fmt.Println(secs)
			total_no_olap += 1
		}
    }
    fmt.Println(total_count - total_no_olap)
    readFile.Close()
}