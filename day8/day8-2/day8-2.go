package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	readFile, err := os.Open("../day8.txt")
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)
	matx := make([][]int, 0)
    for fileScanner.Scan() {
		line := fileScanner.Text()
		row := make([]int, 0)
		for i, _ := range line {
			row = append(row, int(line[i]-'0'))
		}
		matx = append(matx, row)
    }

	max := 0
	l := len(matx)
	w := len(matx[0])
	for y := 1; y < l-1; y++ {
		for x := 1; x < w-1; x++ {
			left := 0
			for i := x - 1; i >= 0; i-- {
				left++
				if matx[y][i] >= matx[y][x] {
					break
				}
			}

			right := 0
			for i := x + 1; i < w; i++ {
				right++
				if matx[y][i] >= matx[y][x] {
					break
				}
			}

			top := 0
			for i := y - 1; i >= 0; i-- {
				top++
				if matx[i][x] >= matx[y][x] {
					break
				}
			}

			bottom := 0
			for i := y + 1; i < l; i++ {
				bottom++
				if matx[i][x] >= matx[y][x] {
					break
				}
			}
			s := left * right * top * bottom
			if s > max {
				max = s
			}
		}
	}

	fmt.Println(max)
    readFile.Close()
}