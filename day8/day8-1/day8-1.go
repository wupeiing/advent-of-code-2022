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
	matxb := make([][]int, 0)
    for fileScanner.Scan() {
		line := fileScanner.Text()
		row := make([]int, 0)
		for i, _ := range line {
			row = append(row, int(line[i]-'0'))
		}
		rowb := make([]int, len(row))
		matx = append(matx, row)
		matxb = append(matxb, rowb)
    }

	base := 99 + 99 + 97 + 97

	for i:=1;i<len(matx)-1;i++ {
		max := matx[i][0]
		for j:=1;j<len(matx[i])-1;j++ {
			if matx[i][j] > max {
				max = matx[i][j]
				if matxb[i][j] == 0 {
					base = base + 1
					matxb[i][j] = 1
				}
			}
		}
	}

	for i:=1;i<len(matx)-1;i++ {
		max := matx[i][len(matx[0])-1]
		for j:=len(matx[i])-2;j>=1;j-- {
			if matx[i][j] > max {
				max = matx[i][j]
				if matxb[i][j] == 0 {
					base = base + 1
					matxb[i][j] = 1
				}
			}
		}
	}

	for i:=1;i<len(matx)-1;i++ {
		max := matx[0][i]
		for j:=1;j<len(matx[i])-1;j++ {
			if matx[j][i] > max {
				max = matx[j][i]
				if matxb[j][i] == 0 {
					base = base + 1
					matxb[j][i] = 1
				}
			}
		}
	}

	for i:=1;i<len(matx)-1;i++ {
		max := matx[len(matx[0])-1][i]
		for j:=len(matx[i])-2;j>=1;j-- {
			if matx[j][i] > max {
				max = matx[j][i]
				if matxb[j][i] == 0 {
					base = base + 1
					matxb[j][i] = 1
				}
			}
		}
	}

	fmt.Println(base)
    readFile.Close()
}