package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
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

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int64, integers ...int64) int64 {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func isNotDone(curNodes []*Node, endNodes []*Node) bool {

	for i := 0; i < len(curNodes); i += 1 {
		if curNodes[i] != endNodes[i] {
			return true
		}
	}

	return false
}

func isNotDone_v2(cycleCounts map[int]int64) bool {
	for i := range cycleCounts {
		if cycleCounts[i] == 0 {
			return true
		}
	}

	return false
}

func walkMap(startNodes []*Node, path string, endNodes []*Node) []int64 {
	stepCount := int64(0)
	stepLen := int64(len(path))

	currentNodes := startNodes
	cycleCounts := map[int]int64{}
	cycleCounts[0] = 0
	cycleCounts[1] = 0
	cycleCounts[2] = 0
	cycleCounts[3] = 0
	cycleCounts[4] = 0
	cycleCounts[5] = 0

	for isNotDone_v2(cycleCounts) {
		for i, currentNode := range currentNodes {
			if cycleCounts[i] != 0 {
				continue
			}
			if path[stepCount%stepLen] == 'L' {
				currentNodes[i] = currentNode.left
			} else if path[stepCount%stepLen] == 'R' {
				currentNodes[i] = currentNode.right
			} else {
				panic("Invalid Direction!")
			}

			if currentNodes[i].name[2] == 'Z' {
				cycleCounts[i] = stepCount + 1
			}
		}

		stepCount += 1
	}

	var counts = []int64{}
	for _, v := range cycleCounts {
		counts = append(counts, v)
	}

	return counts
}

func main() {
	file, err := os.Open("input_data")
	// file, err := os.Open("examp_input")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	path := scanner.Text()

	nodeList := map[string]*Node{}
	var rgx = regexp.MustCompile(`(\w\w\w) = \((\w\w\w), (\w\w\w)\)`)

	startNodesNames := []string{}
	endNodesNames := []string{}

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
		if node_name[2] == 'A' {
			startNodesNames = append(startNodesNames, node_name)
		} else if node_name[2] == 'Z' {
			endNodesNames = append(endNodesNames, node_name)
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
	sort.Strings(startNodesNames)
	sort.Strings(endNodesNames)

	startNodes := []*Node{}
	for _, name := range startNodesNames {
		startNodes = append(startNodes, nodeList[name])
	}
	endNodes := []*Node{}
	for _, name := range endNodesNames {
		endNodes = append(endNodes, nodeList[name])
	}

	cycles := walkMap(startNodes, path, endNodes)

	fmt.Println("Steps taken: ", LCM(cycles[0], cycles[1], cycles[2], cycles[3], cycles[4], cycles[5]))
}
