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
	return int(math.Min(float64(ml.x+1), float64(map_length*3)))
}

func (ml mapLoc) upOrZero() int {
	return int(math.Max(float64(ml.y-1), 0))
}

func (ml mapLoc) downOrMax() int {
	return int(math.Min(float64(ml.y+1), float64(map_height*3)))
}

func convertPoint(x, y int) (int, int) {
	return (x * 3) + 1, (y * 3) + 1
}

type mapDirection int

const (
	LEFT mapDirection = iota
	RIGHT
	UP
	DOWN
)

var pipe_map []string
var big_pipe_map [][]byte
var map_length int
var map_height int

func fillPipe(startPoint, curStep mapLoc, dir mapDirection) {
	if startPoint.x == curStep.x && startPoint.y == curStep.y {
		big_pipe_map[curStep.y][curStep.x] = '#'
		return
	}

	for true {
		step_char := big_pipe_map[curStep.y][curStep.x]
		big_pipe_map[curStep.y][curStep.x] = '#'
		if startPoint.x == curStep.x && startPoint.y == curStep.y {
			break
		}

		if dir == LEFT {
			if step_char == '-' {
				curStep = mapLoc{x: curStep.leftOrZero(), y: curStep.y}
				dir = LEFT
			} else if step_char == 'L' {
				curStep = mapLoc{x: curStep.x, y: curStep.upOrZero()}
				dir = UP
			} else if step_char == 'F' {
				curStep = mapLoc{x: curStep.x, y: curStep.downOrMax()}
				dir = DOWN
			} else {
				panic("Shouldn't be here by stepping left")
			}
		} else if dir == RIGHT {
			if step_char == '-' {
				curStep = mapLoc{x: curStep.rightOrMax(), y: curStep.y}
				dir = RIGHT
			} else if step_char == 'J' {
				curStep = mapLoc{x: curStep.x, y: curStep.upOrZero()}
				dir = UP
			} else if step_char == '7' {
				curStep = mapLoc{x: curStep.x, y: curStep.downOrMax()}
				dir = DOWN
			} else {
				panic("Shouldn't be here by stepping right")
			}
		} else if dir == UP {
			if step_char == '|' {
				curStep = mapLoc{x: curStep.x, y: curStep.upOrZero()}
				dir = UP
			} else if step_char == 'F' {
				curStep = mapLoc{x: curStep.rightOrMax(), y: curStep.y}
				dir = RIGHT
			} else if step_char == '7' {
				curStep = mapLoc{x: curStep.leftOrZero(), y: curStep.y}
				dir = LEFT
			} else {
				panic("Shouldn't be here by stepping up")
			}
		} else if dir == DOWN {
			if step_char == '|' {
				curStep = mapLoc{x: curStep.x, y: curStep.downOrMax()}
				dir = DOWN
			} else if step_char == 'L' {
				curStep = mapLoc{x: curStep.rightOrMax(), y: curStep.y}
				dir = RIGHT
			} else if step_char == 'J' {
				curStep = mapLoc{x: curStep.leftOrZero(), y: curStep.y}
				dir = LEFT
			} else {
				panic("Shouldn't be here by stepping down")
			}
		}
	}
}

func fillGround() {
	for i := range big_pipe_map {
		for j := range big_pipe_map[i] {
			if big_pipe_map[i][j] != '#' {
				big_pipe_map[i][j] = '.'
			}
		}
	}
}

func floodFill(row, col int, oldC, newC byte) {

	// Base Cases
	if row < 0 || col < 0 || row >= (map_height*3) || col >= (map_length*3) {
		return
	}

	if big_pipe_map[row][col] != oldC {
		return
	}

	big_pipe_map[row][col] = newC

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
	for y := range pipe_map {
		for x := range pipe_map[y] {
			new_x, new_y := convertPoint(x, y)
			if big_pipe_map[new_y][new_x] == '.' {
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
	for _, line := range big_pipe_map {
		fmt.Println(line)
	}
}

func convertMap() {
	convert_map := map[byte][]string{
		'.': {"...",
			"...",
			"..."},
		'|': {".|.",
			".|.",
			".|."},
		'-': {"...",
			"---",
			"..."},
		'7': {"...",
			"-7.",
			".|."},
		'F': {"...",
			".F-",
			".|."},
		'J': {".|.",
			"-J.",
			"..."},
		'L': {".|.",
			".L-",
			"..."},
		'S': {"...",
			"---",
			"..."}}

	big_height := len(pipe_map) * 3
	big_length := len(pipe_map[0]) * 3

	big_pipe_map = make([][]byte, big_height)
	for i := range big_pipe_map {
		big_pipe_map[i] = make([]byte, big_length)
	}

	for y := range pipe_map {
		for x := range pipe_map[y] {
			new_x, new_y := convertPoint(x, y)
			populateBigMap(new_x, new_y, convert_map[pipe_map[y][x]])
		}
	}
}

func populateBigMap(x, y int, copy_data []string) {
	start_x := x - 1
	start_y := y - 1
	for j := 0; j <= 2; j += 1 {
		for i := 0; i <= 2; i += 1 {
			big_pipe_map[start_y+j][start_x+i] = copy_data[j][i]
		}
	}
}

func main() {
	file, err := os.Open("input_data")
	// file, err := os.Open("examp_input_2")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// pipe_map = []string{}

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

	convertMap()

	// fmt.Println("Start Point = (", start_point.x, ", ", start_point.y, ")")
	test_char1 := pipe_map[start_point.y][start_point.leftOrZero()]
	test_char2 := pipe_map[start_point.y][start_point.rightOrMax()]
	test_char3 := pipe_map[start_point.upOrZero()][start_point.x]
	test_char4 := pipe_map[start_point.downOrMax()][start_point.x]

	big_start_x, big_start_y := convertPoint(start_point.x, start_point.y)
	big_start := mapLoc{x: big_start_x, y: big_start_y}
	if test_char1 == '-' || test_char1 == 'L' || test_char1 == 'F' {
		fillPipe(big_start, mapLoc{x: big_start.leftOrZero(), y: big_start.y}, LEFT)
	} else if test_char2 == '-' || test_char2 == 'J' || test_char2 == '7' {
		fillPipe(big_start, mapLoc{x: big_start.rightOrMax(), y: big_start.y}, RIGHT)
	} else if test_char3 == '|' || test_char3 == '7' || test_char3 == 'F' {
		fillPipe(big_start, mapLoc{x: big_start.x, y: big_start.upOrZero()}, UP)
	} else if test_char4 == '|' || test_char4 == 'J' || test_char4 == 'L' {
		fillPipe(big_start, mapLoc{x: big_start.x, y: big_start.downOrMax()}, DOWN)
	}

	fillGround()
	floodFill(0, 0, '.', '*')
	// printMap()
	fmt.Println("Enclosed spaces: ", countEnclosed())
}
