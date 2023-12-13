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

/*
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
} */

type mapDirection int

const (
	LEFT mapDirection = iota
	RIGHT
	UP
	DOWN
)

type pattern_map []string

var all_maps []pattern_map

func printMap(pmap pattern_map) {
	for _, line := range pmap {
		fmt.Println(line)
	}
}

func countMirrors_orig(pmap pattern_map) (int, int) {
	x := 0
	y := 0

	for col := 1; col < len(pmap[0]); col += 1 {
		if isMirroredCol(pmap, col) {
			x += col
		}
	}

	for row := 1; row < len(pmap); row += 1 {
		if isMirroredRow(pmap, row) {
			y += row
		}
	}

	return x, y
}

func countMirrors_orig_with_skips(pmap pattern_map, skip_x, skip_y int) (int, int) {
	x := 0
	y := 0

	for col := 1; col < len(pmap[0]); col += 1 {
		if isMirroredCol(pmap, col) && col != skip_x {
			x = col
			break
		}
	}

	for row := 1; row < len(pmap); row += 1 {
		if isMirroredRow(pmap, row) && row != skip_y {
			y = row
			break
		}
	}

	return x, y
}

func countMirrors(pmap pattern_map) int {
	// count := 0

	orig_x, orig_y := countMirrors_orig(pmap)

	for row := 0; row < len(pmap); row += 1 {
		for col := 0; col < len(pmap[0]); col += 1 {
			// var new_map pattern_map
			new_map := make([]string, len(pmap))
			copy(new_map, pmap)
			rep_byte := '.'
			line := new_map[row]
			if line[col] == byte(rep_byte) {
				rep_byte = '#'
			}
			new_map[row] = line[:col] + string(rep_byte) + line[col+1:]

			new_x, new_y := countMirrors_orig_with_skips(new_map, orig_x, orig_y)

			if new_x != 0 && new_x != orig_x {
				return new_x
			} else if new_y != 0 && new_y != orig_y {
				return new_y * 100
			}
		}
	}

	// panic("DIdn't find new line")
	return 0
}

func isMirroredCol(pmap pattern_map, col int) bool {
	for sym_line := 0; ; sym_line += 1 {
		sym_left := col - sym_line - 1
		sym_right := col + sym_line
		if sym_left < 0 || sym_right >= len(pmap[0]) {
			break
		}
		for row := 0; row < len(pmap); row += 1 {
			if pmap[row][sym_left] != pmap[row][sym_right] {
				return false
			}
		}
	}

	return true
}

func isMirroredRow(pmap pattern_map, row int) bool {
	for sym_line := 0; ; sym_line += 1 {
		sym_up := row - sym_line - 1
		sym_down := row + sym_line
		if sym_up < 0 || sym_down >= len(pmap) {
			break
		}
		for col := 0; col < len(pmap[0]); col += 1 {
			if pmap[sym_up][col] != pmap[sym_down][col] {
				return false
			}
		}
	}
	return true
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
		if len(map_line) == 0 {
			// new map
			all_maps = append(all_maps, cur_map)
			cur_map = pattern_map{}
			continue
		}

		cur_map = append(cur_map, map_line)
	}
	all_maps = append(all_maps, cur_map)

	sum := 0
	for _, cur_map := range all_maps {
		// printMap(cur_map)
		mirror_count := countMirrors(cur_map)
		// fmt.Println("Mirror count: ", mirror_count)
		sum += mirror_count
		// fmt.Println()
	}
	fmt.Println("Total Mirror Count: ", sum)
}
