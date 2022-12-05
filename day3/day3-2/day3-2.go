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
	l_count := 0
	var aa, bb, cc []byte
    for fileScanner.Scan() {
		if l_count % 3 == 0 {
			aa = []byte(fileScanner.Text())
			l_count++
			continue
		} else if l_count % 3 == 1 {
			bb = []byte(fileScanner.Text())
			l_count++
			continue
		} else {
			cc = []byte(fileScanner.Text())
			l_count++
		}
		// fmt.Println("aa: ", string(aa))
		// fmt.Println("bb: ", string(bb))
		// fmt.Println("cc ", string(cc))
		arr := make([]int, 53)
		arrb := make([]int, 53)
		for _, b := range aa {
			if b <= 'Z' {
				arr[b - 'A' + 26]++
			} else {
				arr[b - 'a']++
			}
		}
		for _, by := range bb {
			if by <= 'Z' {
				if arr[by-'A' + 26] > 0 {
					arrb[by - 'A' + 26]++
				}
			} else {
				if arr[by-'a'] > 0 {
					arrb[by - 'a']++
				}
			}
		}
		for _, bz := range cc {
			if bz <= 'Z' {
				if arrb[bz-'A' + 26] > 0 {
					total += int(bz - 'A') + 27
					break
				}
			} else {
				if arrb[bz-'a'] > 0 {
					total += int(bz - 'a') + 1
					break
				}
			}
		}
    }
    fmt.Println(total)
    readFile.Close()
}