package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
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

var galaxy_map []string
var empty_rows []int
var empty_cols []int

var map_length int
var map_height int

func findGalaxies() []mapLoc {
	galaxy_list := []mapLoc{}
	for row, line := range galaxy_map {
		for col, char := range line {
			if char == '#' {
				galaxy_list = append(galaxy_list, mapLoc{x: col, y: row})
			}
		}
	}

	return galaxy_list
}

func calcPairs(galaxy_list []mapLoc) map[mapLoc]int {
	galaxy_pairs := map[mapLoc]int{}

	// galaxy_pairs[mapLoc{x: 1, y: 6}] = 0

	for i := 0; i < len(galaxy_list); i++ {
		for j := i + 1; j < len(galaxy_list); j++ {
			pair := mapLoc{x: i, y: j}
			galaxy_pairs[pair] = 0
		}
	}

	return galaxy_pairs
}

func bresenham_count(x1, y1, x2, y2 int) int {
	count := 0
	dx := x2 - x1
	dy := y2 - y1
	y := y1
	eps := 0

	for x := x1; x <= x2; x += 1 {
		count += 1
		eps += dy
		if (eps << 1) >= dx {
			count += 1
			y++
			eps -= dx
		}
	}

	return count
}

func naive_count(loc1, loc2 mapLoc) int {
	base_dist := int(math.Abs(float64(loc1.x-loc2.x)) + math.Abs(float64(loc1.y-loc2.y)))

	return base_dist + add_empties(loc1, loc2)
}

func add_empties(loc1, loc2 mapLoc) int {
	empty_multiplier := 1000000
	empties := 0
	if loc1.x > loc2.x {
		for i := loc2.x; i < loc1.x; i++ {
			if slices.Contains(empty_cols, i) {
				empties += (1 * empty_multiplier) - 1
			}
		}
	} else if loc1.x < loc2.x {
		for i := loc1.x; i < loc2.x; i++ {
			if slices.Contains(empty_cols, i) {
				empties += (1 * empty_multiplier) - 1
			}
		}
	}

	if loc1.y > loc2.y {
		for i := loc2.y; i < loc1.y; i++ {
			if slices.Contains(empty_rows, i) {
				empties += (1 * empty_multiplier) - 1
			}
		}
	} else if loc1.y < loc2.y {
		for i := loc1.y; i < loc2.y; i++ {
			if slices.Contains(empty_rows, i) {
				empties += (1 * empty_multiplier) - 1
			}
		}
	}

	return empties
}

func calcDistances(galaxy_pairs map[mapLoc]int, galaxy_list []mapLoc) int {
	sum := 0
	for k := range galaxy_pairs {
		gal1 := galaxy_list[k.x]
		gal2 := galaxy_list[k.y]

		galaxy_pairs[k] = naive_count(gal1, gal2)
		sum += galaxy_pairs[k]
	}

	return sum
}

func printMap() {
	for _, line := range galaxy_map {
		fmt.Println(line)
	}
}

func main() {
	file, err := os.Open("input_data")
	// file, err := os.Open("examp_input")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	galaxy_map = []string{}

	cur_y := 0
	for scanner.Scan() {
		map_line := scanner.Text()
		if len(map_line) == 0 {
			continue
		}
		galaxy_map = append(galaxy_map, map_line)
		i := strings.Index(map_line, "#")
		if i == -1 {
			empty_rows = append(empty_rows, cur_y)
		}

		cur_y += 1
	}

	for i := 0; i < len(galaxy_map[0]); i += 1 {
		has_galaxy := false
		for j := 0; j < len(galaxy_map); j += 1 {
			if galaxy_map[j][i] == '#' {
				has_galaxy = true
				break
			}
		}

		if !has_galaxy {
			empty_cols = append(empty_cols, i)
		}
	}

	map_height = len(galaxy_map)
	map_length = len(galaxy_map[0])
	printMap()

	galaxy_list := findGalaxies()
	galaxy_pairs := calcPairs(galaxy_list)
	sum := 0
	sum += calcDistances(galaxy_pairs, galaxy_list)

	fmt.Println("Distance Sum: ", sum)
}
