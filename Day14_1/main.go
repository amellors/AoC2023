package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type mapLoc struct {
	x int
	y int
}

type mapDirection int

const (
	LEFT mapDirection = iota
	RIGHT
	UP
	DOWN
)

type pattern_map []string

func printMap(pmap pattern_map) {
	for _, line := range pmap {
		fmt.Println(line)
	}
}

func countRockWeight(pmap pattern_map, col int) int {
	totalWeight := 0
	maxRockWeight := len(pmap)
	currRockWeight := maxRockWeight
	for row := range pmap {
		if pmap[row][col] == 'O' {
			totalWeight += currRockWeight
			currRockWeight -= 1
		} else if pmap[row][col] == '#' {
			currRockWeight = maxRockWeight - row - 1
		}
	}

	return totalWeight
}

func main() {
	file, err := os.Open("input_data")
	// file, err := os.Open("examp_input")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	cur_map := pattern_map{}
	for scanner.Scan() {

		map_line := scanner.Text()

		cur_map = append(cur_map, map_line)
	}

	sum := 0
	for i := range cur_map[0] {
		colWeight := countRockWeight(cur_map, i)
		sum += colWeight
	}
	fmt.Println("Total Rock Weight: ", sum)
}
