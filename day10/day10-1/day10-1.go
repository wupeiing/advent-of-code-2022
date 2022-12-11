package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {

	readFile, err := os.Open("../day10.txt")
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)
	cycle := 0
	total := 0
	count := 1
    for fileScanner.Scan() {
		line := fileScanner.Text()
		inst := strings.Split(line, " ")
		if len(inst) == 1 {
			cycle += 1
			if cycle == 20 || (cycle - 20)%40 == 0 {
				total += count * cycle
			}
		} else {
			op, _ := strconv.Atoi(inst[1])
			for i:=0;i<2;i++ {
				cycle += 1
				if cycle == 20 || (cycle - 20)%40 == 0 {
					total += count * cycle
				}
			}
			count += op
		}
    }
	fmt.Println(total)
    readFile.Close()
}