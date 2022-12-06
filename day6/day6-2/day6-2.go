package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	readFile, err := os.Open("../day6.txt")
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)
  
	secs := ""
    for fileScanner.Scan() {
		secs = fileScanner.Text()
    }
	var sol [27]int
	for i, _ := range secs {
		ans := false
		sol = [27]int{}
		for j := 0; j<14; j++ {
			t := secs[i+j] - 'a'
			if sol[t] == 0 {
				sol[t]++
			} else {
				ans = false
				break
			}
			ans = true
		}
		if ans == true {
			fmt.Println(i+14)
			break
		}
	}
    readFile.Close()
}