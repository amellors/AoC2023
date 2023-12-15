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
	UP mapDirection = iota
	LEFT
	DOWN
	RIGHT
)

type pattern_map []string

func printMap(pmap pattern_map) {
	for _, line := range pmap {
		fmt.Println(line)
	}
}

func countRockWeight(pmap pattern_map) int {
	maxRockWeight := len(pmap)
	totalWeight := 0
	for row := range pmap {
		for col := range pmap[row] {
			if pmap[row][col] == 'O' {
				totalWeight += maxRockWeight - row
			}
		}
	}

	return totalWeight
}

func rollRocks(pmap pattern_map, dir mapDirection) pattern_map {
	maxHeight := len(pmap)
	newMap := make(pattern_map, maxHeight)
	newMap_byte := make([][]byte, maxHeight)

	for i := range newMap_byte {
		newMap_byte[i] = make([]byte, maxHeight)
	}

	currRockHeight := 0
	if dir == UP {
		for col := range pmap[0] {
			currRockHeight = 0
			for row := range pmap {
				if pmap[row][col] == 'O' {
					newMap_byte[currRockHeight][col] = 'O'
					currRockHeight += 1
				} else if pmap[row][col] == '#' {
					for i := currRockHeight; i < row; i += 1 {
						newMap_byte[i][col] = '.'
					}
					newMap_byte[row][col] = '#'
					currRockHeight = row + 1
				}
			}
			for i := currRockHeight; i < maxHeight; i += 1 {
				newMap_byte[i][col] = '.'
			}
		}
		for r := range newMap_byte {
			newMap[r] = string(newMap_byte[r])
		}
	} else if dir == LEFT {
		for row := range pmap[0] {
			currRockHeight = 0
			for col := range pmap {
				if pmap[row][col] == 'O' {
					newMap_byte[row][currRockHeight] = 'O'
					currRockHeight += 1
				} else if pmap[row][col] == '#' {
					for i := currRockHeight; i < col; i += 1 {
						newMap_byte[row][i] = '.'
					}
					newMap_byte[row][col] = '#'
					currRockHeight = col + 1
				}
			}
			for i := currRockHeight; i < maxHeight; i += 1 {
				newMap_byte[row][i] = '.'
			}
		}
		for r := range newMap_byte {
			newMap[r] = string(newMap_byte[r])
		}
	} else if dir == RIGHT {
		for row := len(pmap[0]) - 1; row >= 0; row -= 1 {
			currRockHeight = len(pmap[0]) - 1
			for col := len(pmap) - 1; col >= 0; col -= 1 {
				if pmap[row][col] == 'O' {
					newMap_byte[row][currRockHeight] = 'O'
					currRockHeight -= 1
				} else if pmap[row][col] == '#' {
					for i := currRockHeight; i > col; i -= 1 {
						newMap_byte[row][i] = '.'
					}
					newMap_byte[row][col] = '#'
					currRockHeight = col - 1
				}
			}
			for i := currRockHeight; i >= 0; i -= 1 {
				newMap_byte[row][i] = '.'
			}
		}
		for r := range newMap_byte {
			newMap[r] = string(newMap_byte[r])
		}
	} else if dir == DOWN {
		for col := range pmap {
			currRockHeight = len(pmap) - 1
			for row := len(pmap[0]) - 1; row >= 0; row -= 1 {
				if pmap[row][col] == 'O' {
					newMap_byte[currRockHeight][col] = 'O'
					currRockHeight -= 1
				} else if pmap[row][col] == '#' {
					for i := currRockHeight; i > row; i -= 1 {
						newMap_byte[i][col] = '.'
					}
					newMap_byte[row][col] = '#'
					currRockHeight = row - 1
				}
			}
			for i := currRockHeight; i >= 0; i -= 1 {
				newMap_byte[i][col] = '.'
			}
		}
		for r := range newMap_byte {
			newMap[r] = string(newMap_byte[r])
		}
	}

	return newMap
}

func shiftAndCount(cur_map pattern_map) (int, pattern_map) {

	roll_map := cur_map
	for dir := UP; dir <= RIGHT; dir += 1 {
		roll_map = rollRocks(roll_map, dir)
	}

	totalWeight := countRockWeight(roll_map)

	return totalWeight, roll_map
}

func main() {
	file, err := os.Open("input_data")
	//file, err := os.Open("examp_input")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	cur_map := pattern_map{}
	for scanner.Scan() {

		map_line := scanner.Text()

		cur_map = append(cur_map, map_line)
	}

	temp_map := cur_map
	var totalWeight int
	for i := 1; i <= 1000; i += 1 {
		totalWeight, temp_map = shiftAndCount(temp_map)
		// printMap(new_map)
		//	temp_map = new_map
		//if i%10000 == 0 {
		fmt.Println("itr(", i, ")-North Rock Weight: ", totalWeight)
		//}
	}

	fmt.Println("North Rock Weight: ", totalWeight)

	fmt.Println("loop pos: ", (1000000000-177)%14)
}
