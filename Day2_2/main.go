package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func processLine(line string) int64 {
	var game_id int
	games := regexp.MustCompile(": ").Split(line, -1)
	fmt.Sscanf(games[0], "Game %d", &game_id)

	var max map[string]int
	max = make(map[string]int)
	max["red"] = 1
	max["green"] = 1
	max["blue"] = 1

	hands := regexp.MustCompile("; ").Split(games[1], -1)
	for _, hand := range hands {
		pulls := regexp.MustCompile(", ").Split(hand, -1)
		for _, pull := range pulls {
			var color string
			var count int
			fmt.Sscanf(pull, "%d %s", &count, &color)
			max[color] = int(math.Max(float64(count), float64(max[color])))
		}
	}

	return int64(max["red"]) * int64(max["green"]) * int64(max["blue"])
}

func main() {
	file, err := os.Open("input_data")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	var sum int64
	sum = 0
	for scanner.Scan() {
		sum += processLine(scanner.Text())
	}

	check(err)

	fmt.Println("sum is: " + strconv.FormatInt(sum, 10))
}
