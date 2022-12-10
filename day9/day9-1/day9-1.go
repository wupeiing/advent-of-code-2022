package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {

	readFile, err := os.Open("../day9.txt")
	//readFile, err := os.Open("../day9sample.txt")
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)
	mp := make(map[string]int, 0)
	hx, hy := 0, 0
	tx, ty := 0, 0
	h := fmt.Sprintf("%d , %d", tx, ty)
	mp[h] +=1
    for fileScanner.Scan() {
		line := fileScanner.Text()
		strs := strings.Split(line, " ")
		steps, _ := strconv.Atoi(strs[1])
		//fmt.Println(str1[0], "    ", str1[1])
		if strs[0] == "U" {
			for i := 0; i<steps; i++ {
				hy += 1
				if abs(hx - tx) > 1 || abs(hy - ty) > 1 {
					tx = hx
					ty = hy - 1
					h = fmt.Sprintf("%d , %d", tx, ty)
					mp[h] +=1
				}
			}
		}
		if strs[0] == "D" {
			for i := 0; i<steps; i++ {
				hy -= 1
				if abs(hx - tx) > 1 || abs(hy - ty) > 1 {
					tx = hx
					ty = hy + 1
					h = fmt.Sprintf("%d , %d", tx, ty)
					mp[h] +=1
				}
			}
		}
		if strs[0] == "L" {
			for i := 0; i<steps; i++ {
				hx -= 1
				if abs(hx - tx) > 1 || abs(hy - ty) > 1 {
					ty = hy
					tx = hx + 1
					h = fmt.Sprintf("%d , %d", tx, ty)
					mp[h] +=1
				}
			}
		}
		if strs[0] == "R" {
			for i := 0; i<steps; i++ {
				hx += 1
				if abs(hx - tx) > 1 || abs(hy - ty) > 1 {
					ty = hy
					tx = hx - 1
					h = fmt.Sprintf("%d , %d", tx, ty)
					mp[h] +=1
				}
			}
		}
    }
	fmt.Println(len(mp))
    readFile.Close()
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}