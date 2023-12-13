package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Node struct {
	name  string
	left  *Node
	right *Node
}

func walkMap(root *Node, path string, endNode *Node) int64 {
	stepCount := int64(0)
	stepLen := int64(len(path))

	currentNode := root

	for currentNode != endNode {
		if path[stepCount%stepLen] == 'L' {
			currentNode = currentNode.left
		} else if path[stepCount%stepLen] == 'R' {
			currentNode = currentNode.right
		} else {
			panic("Invalid Direction!")
		}

		stepCount += 1
	}

	return stepCount
}

func main() {
	file, err := os.Open("input_data")
	//file, err := os.Open("examp_input_2")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	path := scanner.Text()

	nodeList := map[string]*Node{}
	var rgx = regexp.MustCompile(`(\w\w\w) = \((\w\w\w), (\w\w\w)\)`)

	for scanner.Scan() {
		map_line := scanner.Text()

		if len(map_line) == 0 {
			continue
		}

		rs := rgx.FindStringSubmatch(map_line)
		node_name := rs[1]
		node_left := rs[2]
		node_right := rs[3]

		val, node_found := nodeList[node_name]
		if !node_found {
			val = &Node{name: node_name}
			nodeList[node_name] = val
		}

		left_val, node_found := nodeList[node_left]
		if !node_found {
			left_val = &Node{name: node_left}
			nodeList[node_left] = left_val
		}
		val.left = left_val

		right_val, node_found := nodeList[node_right]
		if !node_found {
			right_val = &Node{name: node_right}
			nodeList[node_right] = right_val
		}
		val.right = right_val
	}

	steps := walkMap(nodeList["AAA"], path, nodeList["ZZZ"])

	fmt.Println("Steps taken: ", steps)
}
