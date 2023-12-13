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

func nextStep(startPoint, curStep mapLoc, dir mapDirection) int {
	if startPoint.x == curStep.x && startPoint.y == curStep.y {
		return 0
	}

	step_char := pipe_map[curStep.y][curStep.x]
	if dir == LEFT {
		if step_char == '-' {
			return 1 + nextStep(startPoint, mapLoc{x: curStep.leftOrZero(), y: curStep.y}, LEFT)
		} else if step_char == 'L' {
			return 1 + nextStep(startPoint, mapLoc{x: curStep.x, y: curStep.upOrZero()}, UP)
		} else if step_char == 'F' {
			return 1 + nextStep(startPoint, mapLoc{x: curStep.x, y: curStep.downOrMax()}, DOWN)
		} else {
			panic("Shouldn't be here by stepping left")
		}
	} else if dir == RIGHT {
		if step_char == '-' {
			return 1 + nextStep(startPoint, mapLoc{x: curStep.rightOrMax(), y: curStep.y}, RIGHT)
		} else if step_char == 'J' {
			return 1 + nextStep(startPoint, mapLoc{x: curStep.x, y: curStep.upOrZero()}, UP)
		} else if step_char == '7' {
			return 1 + nextStep(startPoint, mapLoc{x: curStep.x, y: curStep.downOrMax()}, DOWN)
		} else {
			panic("Shouldn't be here by stepping right")
		}
	} else if dir == UP {
		if step_char == '|' {
			return 1 + nextStep(startPoint, mapLoc{x: curStep.x, y: curStep.upOrZero()}, UP)
		} else if step_char == 'F' {
			return 1 + nextStep(startPoint, mapLoc{x: curStep.rightOrMax(), y: curStep.y}, RIGHT)
		} else if step_char == '7' {
			return 1 + nextStep(startPoint, mapLoc{x: curStep.leftOrZero(), y: curStep.y}, LEFT)
		} else {
			panic("Shouldn't be here by stepping up")
		}
	} else if dir == DOWN {
		if step_char == '|' {
			return 1 + nextStep(startPoint, mapLoc{x: curStep.x, y: curStep.downOrMax()}, DOWN)
		} else if step_char == 'L' {
			return 1 + nextStep(startPoint, mapLoc{x: curStep.rightOrMax(), y: curStep.y}, RIGHT)
		} else if step_char == 'J' {
			return 1 + nextStep(startPoint, mapLoc{x: curStep.leftOrZero(), y: curStep.y}, LEFT)
		} else {
			panic("Shouldn't be here by stepping down")
		}

	}
	return 1
}

func printAnswer(steps int) {
	half := (steps + steps%2) / 2
	fmt.Println("Max distance: ", half)
}

func main() {
	file, err := os.Open("input_data")
	//file, err := os.Open("examp_input")
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

	fmt.Println("Start Point = (", start_point.x, ", ", start_point.y, ")")
	test_char := pipe_map[start_point.y][start_point.leftOrZero()]
	if test_char == '-' || test_char == 'L' || test_char == 'F' {
		steps := nextStep(start_point, mapLoc{x: start_point.leftOrZero(), y: start_point.y}, LEFT)
		printAnswer(steps)
		return
	}

	test_char = pipe_map[start_point.y][start_point.rightOrMax()]
	if test_char == '-' || test_char == 'J' || test_char == '7' {
		steps := nextStep(start_point, mapLoc{x: start_point.rightOrMax(), y: start_point.y}, RIGHT)
		printAnswer(steps)
		return
	}

	test_char = pipe_map[start_point.upOrZero()][start_point.x]
	if test_char == '|' || test_char == '7' || test_char == 'F' {
		steps := nextStep(start_point, mapLoc{x: start_point.x, y: start_point.upOrZero()}, UP)
		printAnswer(steps)
		return
	}

	test_char = pipe_map[start_point.downOrMax()][start_point.x]
	if test_char == '|' || test_char == 'J' || test_char == 'L' {
		steps := nextStep(start_point, mapLoc{x: start_point.x, y: start_point.downOrMax()}, DOWN)
		printAnswer(steps)
		return
	}
}
