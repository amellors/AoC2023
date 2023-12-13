package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
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

func (ml mapLoc) leftOrZero() int {
	return int(math.Max(float64(ml.x-1), 0))
}

func (ml mapLoc) rightOrMax() int {
	return int(math.Min(float64(ml.x+1), float64(map_length)))
}

func (ml mapLoc) upOrZero() int {
	return int(math.Max(float64(ml.y-1), 0))
}

func (ml mapLoc) downOrMax() int {
	return int(math.Min(float64(ml.y+1), float64(map_height)))
}

type mapDirection int

const (
	LEFT mapDirection = iota
	RIGHT
	UP
	DOWN
)

var pipe_map []string
var map_length int
var map_height int

func fillPipe(startPoint, curStep mapLoc, dir mapDirection) {
	if startPoint.x == curStep.x && startPoint.y == curStep.y {
		line := pipe_map[curStep.y]
		pipe_map[curStep.y] = line[:curStep.x] + string('#') + line[curStep.x+1:]
		return
	}

	step_char := pipe_map[curStep.y][curStep.x]
	if dir == LEFT {
		if step_char == '-' {
			fillPipe(startPoint, mapLoc{x: curStep.leftOrZero(), y: curStep.y}, LEFT)
		} else if step_char == 'L' {
			fillPipe(startPoint, mapLoc{x: curStep.x, y: curStep.upOrZero()}, UP)
		} else if step_char == 'F' {
			fillPipe(startPoint, mapLoc{x: curStep.x, y: curStep.downOrMax()}, DOWN)
		} else {
			panic("Shouldn't be here by stepping left")
		}
	} else if dir == RIGHT {
		if step_char == '-' {
			fillPipe(startPoint, mapLoc{x: curStep.rightOrMax(), y: curStep.y}, RIGHT)
		} else if step_char == 'J' {
			fillPipe(startPoint, mapLoc{x: curStep.x, y: curStep.upOrZero()}, UP)
		} else if step_char == '7' {
			fillPipe(startPoint, mapLoc{x: curStep.x, y: curStep.downOrMax()}, DOWN)
		} else {
			panic("Shouldn't be here by stepping right")
		}
	} else if dir == UP {
		if step_char == '|' {
			fillPipe(startPoint, mapLoc{x: curStep.x, y: curStep.upOrZero()}, UP)
		} else if step_char == 'F' {
			fillPipe(startPoint, mapLoc{x: curStep.rightOrMax(), y: curStep.y}, RIGHT)
		} else if step_char == '7' {
			fillPipe(startPoint, mapLoc{x: curStep.leftOrZero(), y: curStep.y}, LEFT)
		} else {
			panic("Shouldn't be here by stepping up")
		}
	} else if dir == DOWN {
		if step_char == '|' {
			fillPipe(startPoint, mapLoc{x: curStep.x, y: curStep.downOrMax()}, DOWN)
		} else if step_char == 'L' {
			fillPipe(startPoint, mapLoc{x: curStep.rightOrMax(), y: curStep.y}, RIGHT)
		} else if step_char == 'J' {
			fillPipe(startPoint, mapLoc{x: curStep.leftOrZero(), y: curStep.y}, LEFT)
		} else {
			panic("Shouldn't be here by stepping down")
		}
	}

	line := pipe_map[curStep.y]
	pipe_map[curStep.y] = line[:curStep.x] + string('#') + line[curStep.x+1:]
}

func fillGround() {
	for i, line := range pipe_map {
		line = strings.ReplaceAll(line, "F", ".")
		line = strings.ReplaceAll(line, "7", ".")
		line = strings.ReplaceAll(line, "|", ".")
		line = strings.ReplaceAll(line, "-", ".")
		line = strings.ReplaceAll(line, "L", ".")
		line = strings.ReplaceAll(line, "J", ".")
		pipe_map[i] = line
	}
}

func floodFill(row, col int, oldC, newC byte) {

	// Base Cases
	if row < 0 || col < 0 || row >= map_length || col >= map_height {
		return
	}

	if pipe_map[row][col] != oldC {
		return
	}

	line := pipe_map[row]
	pipe_map[row] = line[:col] + string(newC) + line[col+1:]

	// recursion
	floodFill(row+1, col-1, oldC, newC)
	floodFill(row+1, col, oldC, newC)
	floodFill(row+1, col+1, oldC, newC)
	floodFill(row, col+1, oldC, newC)
	floodFill(row, col-1, oldC, newC)
	floodFill(row-1, col-1, oldC, newC)
	floodFill(row-1, col, oldC, newC)
	floodFill(row-1, col+1, oldC, newC)
}

func countEnclosed() int {
	count := 0
	for _, line := range pipe_map {
		for _, char := range line {
			if char == '.' {
				count += 1
			}
		}
	}
	return count
}

func printAnswer(steps int) {
	half := (steps + steps%2) / 2
	fmt.Println("Max distance: ", half)
}

func printMap() {
	for _, line := range pipe_map {
		fmt.Println(line)
	}
}

func main() {
	file, err := os.Open("input_data")
	// file, err := os.Open("examp_input")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	pipe_map = []string{}

	cur_y := 0
	start_point := mapLoc{}
	for scanner.Scan() {
		map_line := scanner.Text()
		if len(map_line) == 0 {
			continue
		}
		pipe_map = append(pipe_map, map_line)
		i := strings.Index(map_line, "S")
		if i > -1 {
			start_point.x = i
			start_point.y = cur_y
		}

		cur_y += 1
	}
	map_height = len(pipe_map)
	map_length = len(pipe_map[0])

	// fmt.Println("Start Point = (", start_point.x, ", ", start_point.y, ")")
	test_char1 := pipe_map[start_point.y][start_point.leftOrZero()]
	test_char2 := pipe_map[start_point.y][start_point.rightOrMax()]
	test_char3 := pipe_map[start_point.upOrZero()][start_point.x]
	test_char4 := pipe_map[start_point.downOrMax()][start_point.x]
	if test_char1 == '-' || test_char1 == 'L' || test_char1 == 'F' {
		fillPipe(start_point, mapLoc{x: start_point.leftOrZero(), y: start_point.y}, LEFT)
	} else if test_char2 == '-' || test_char2 == 'J' || test_char2 == '7' {
		fillPipe(start_point, mapLoc{x: start_point.rightOrMax(), y: start_point.y}, RIGHT)
	} else if test_char3 == '|' || test_char3 == '7' || test_char3 == 'F' {
		fillPipe(start_point, mapLoc{x: start_point.x, y: start_point.upOrZero()}, UP)
	} else if test_char4 == '|' || test_char4 == 'J' || test_char4 == 'L' {
		fillPipe(start_point, mapLoc{x: start_point.x, y: start_point.downOrMax()}, DOWN)
	}

	fillGround()
	floodFill(0, 0, '.', '*')
	printMap()
	fmt.Println("Enclosed spaces: ", countEnclosed())
}
