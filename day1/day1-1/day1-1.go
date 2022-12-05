package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	readFile, err := os.Open("../day1.txt")
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)
  
	max := 0
	temp := 0
    for fileScanner.Scan() {
		if fileScanner.Text() == "" {
			if temp > max {
				max = temp
			}
			temp = 0
		} else {
			i, _ := strconv.Atoi(fileScanner.Text())
			temp += i
		}
    }

	fmt.Println("Max value: ", max)
  
    readFile.Close()
}