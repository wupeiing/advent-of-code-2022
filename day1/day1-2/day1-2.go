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
  
	max1, max2, max3 := 0, 0, 0
	temp := 0
    for fileScanner.Scan() {
		if fileScanner.Text() == "" {
			if temp >= max1 {
				max3 = max2
				max2 = max1
				max1 = temp
			} else if temp >= max2 {
				max3 = max2
				max2 = temp
			} else if temp >= max3 {
				max3 = temp
			}
			temp = 0
		} else {
			i, _ := strconv.Atoi(fileScanner.Text())
			temp += i
		}
    }

	fmt.Println("3 Max value: ", max1 + max2 + max3)
  
    readFile.Close()
}