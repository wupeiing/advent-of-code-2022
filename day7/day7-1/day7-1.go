package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type directory struct {
	size int
	sub map[string]*directory
	parent *directory
}

func getSmallDirs(dir *directory, maxsize int) int {
	total := 0
	for _, subdir := range dir.sub {
		if subdir.size <= maxsize {
			total += subdir.size
		}
		total += getSmallDirs(subdir, maxsize)
	}
	return total
}

func main() {

	readFile, err := os.Open("../day7.txt")
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)

	var root = &directory {
		size:     0,
		sub: make(map[string]*directory),
	}
	var cur = root

    for fileScanner.Scan() {
		line := fileScanner.Text()
		lines := strings.Split(line, " ")
		cmd := lines[0]
		if lines[0] == "$" {
			cmd = lines[1]
		}

		if cmd == "ls" {
			continue
		} else if cmd == "dir" {
			cur.sub[lines[1]] = &directory {
				size:     0,
				sub: make(map[string]*directory),
				parent:   cur,
			}
		} else if cmd == "cd" {
			if lines[2] == ".." {
				cur = cur.parent
			} else if lines[2] == "/" {
				cur = root
			} else {
				nextl, ok := cur.sub[lines[2]]
				if !ok {
					nextl = &directory {
						size: 0,
						sub: make(map[string]*directory),
						parent: cur,
					}
					cur.sub[lines[2]] = nextl
				}
				cur = nextl
			}
		} else {
			// For not cd / dir / ls command, like file
			size, _ := strconv.Atoi(lines[0])
			cur.size += size
			par := cur.parent
			for par != nil {
				par.size += size
				par = par.parent
			}
		}
    }

	fmt.Println(getSmallDirs(root, 100000))

    readFile.Close()
}