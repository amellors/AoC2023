package main

import (
	"bufio"
	"fmt"
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
	games := regexp.MustCompile(": +").Split(line, -1)
	hands := regexp.MustCompile(" \\| +").Split(games[1], -1)
	// winners := [10]int{}
	winners := regexp.MustCompile(" +").Split(hands[0], -1)
	hand := regexp.MustCompile(" +").Split(hands[1], -1)

	if len(winners) != 10 || len(hand) != 25 {
		fmt.Println("Invalid Parsing!")
		return 0
	}

	winning_numbers := map[string]bool{}
	for _, winner := range winners {
		winning_numbers[winner] = true
	}

	points := 0
	for _, hand_num := range hand {
		if winning_numbers[hand_num] {
			if points == 0 {
				points = 1
			} else {
				points *= 2
			}
		}
	}

	return int64(points)
}

func main() {
	file, err := os.Open("input_data")
	check(err)
	defer file.Close()

	// var schematic []string
	sum := 0
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	for scanner.Scan() {
		sum += int(processLine(scanner.Text()))
		// schematic = append(schematic, scanner.Text())
		check(err)
	}

	fmt.Println("Part sum: " + strconv.FormatInt(int64(sum), 10))

}
