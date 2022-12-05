package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	readFile, err := os.Open("../day3.txt")
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)
  
	total := 0
    for fileScanner.Scan() {
		byte_arr := []byte(fileScanner.Text())
		ba := byte_arr[:len(byte_arr)/2]
		bb := byte_arr[len(byte_arr)/2:]
		arr := make([]int, 53)
		for _, b := range ba {
			fmt.Println(string(b))
			if b <= 'Z' {
				arr[b - 'A' + 26]++
			} else {
				arr[b - 'a']++
			}
		}
		for _, by := range bb {
			if by <= 'Z' {
				if arr[by-'A' + 26] > 0 {
					total += int(by - 'A') + 27
					break
				}
			} else {
				if arr[by-'a'] > 0 {
					total += int(by - 'a') + 1
					break
				}
			}
		}
    }
    fmt.Println(total)
    readFile.Close()
}