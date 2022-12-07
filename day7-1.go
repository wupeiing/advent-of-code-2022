package main

import (
	"strings"
	"strconv"
	"bufio"
	"os"
	"fmt"
	"math"
)

type treeNode struct {
	nodeType string
	size     int
	children map[string]*treeNode
	parent   *treeNode
}

func getDirsUnder(tree *treeNode, sizeThreshold int) int {
	total := 0
	for _, node := range tree.children {
		if node.nodeType == "dir" && node.size <= sizeThreshold {
			total += node.size
		}
		total += getDirsUnder(node, sizeThreshold)
	}
	return total
}

func findSmallestToDelete(tree *treeNode, sizeRequired int) int {
	smallest := math.MaxInt
	if tree.nodeType != "dir" {
		return smallest
	}
	if tree.size >= sizeRequired {
		smallest = tree.size
	}
	for _, node := range tree.children {
		if node.nodeType == "dir" && node.size >= sizeRequired && node.size < smallest {
			smallest = node.size
		}
		subSmallest := findSmallestToDelete(node, sizeRequired)
		if subSmallest < smallest {
			smallest = subSmallest
		}
	}
	return smallest
}

func main() {
	// Parse the input and create a representation of the filesystem

	readFile, err := os.Open("day7.txt")
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)

	var rootNode = &treeNode{
		nodeType: "dir",
		size:     0,
		children: make(map[string]*treeNode),
	}
	var currentNode = rootNode

	// Read each line of the input
	for fileScanner.Scan() {
		// Split the line into parts
		line := fileScanner.Text()
		//fmt.Println(line)
		parts := strings.Split(strings.TrimLeft(line, "$ "), " ")
		switch parts[0] {
		case "cd":
			if parts[1] == "/" {
				currentNode = rootNode
			} else if parts[1] == ".." {
				currentNode = currentNode.parent
			} else {
				nextNode, ok := currentNode.children[parts[1]]
				if !ok {
					nextNode = &treeNode{
						nodeType: "dir",
						size:     0,
						children: make(map[string]*treeNode),
						parent:   currentNode,
					}
					currentNode.children[parts[1]] = nextNode
				}
				currentNode = nextNode
			}
		case "dir":
			currentNode.children[parts[1]] = &treeNode{
				nodeType: "dir",
				size:     0,
				children: make(map[string]*treeNode),
				parent:   currentNode,
			}
		case "ls":
			continue
		default:
			size, _ := strconv.Atoi(parts[0])
			fileName := parts[1]
			currentNode.children[fileName] = &treeNode{
				nodeType: "file",
				size:     size,
			}
			currentNode.size += size
			parent := currentNode.parent
			for parent != nil {
				parent.size += size
				parent = parent.parent
			}

		}
	}

	fmt.Println(getDirsUnder(rootNode, 100000))
	requiredSize := 30000000 - (70000000 - rootNode.size)
	fmt.Println(findSmallestToDelete(rootNode, requiredSize))

}