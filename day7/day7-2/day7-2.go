package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type directory struct {
	name string
	size int
	sub map[string]*directory
	parent *directory
}

var ans int = 0
func getSmallest(dir *directory, required int) {
	for _, subdir := range dir.sub {
		if subdir.size >= required {
			if ans == 0 {
				ans = subdir.size
			} else if subdir.size < ans{
				ans = subdir.size
			}
		}
		getSmallest(subdir, required)
	}
}

func main() {

	readFile, err := os.Open("../day7.txt")
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)

	var root = &directory {
		name: "root",
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
				name: lines[1],
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
						name: lines[2],
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

	req := 30000000 - (70000000 - root.size)
	getSmallest(root, req)
	fmt.Println(ans)
    readFile.Close()
}