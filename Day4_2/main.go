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

func processLine(line string) int {
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

	wins := 0
	for _, hand_num := range hand {
		if winning_numbers[hand_num] {
			wins += 1
		}
	}

	return wins
}

func main() {
	file, err := os.Open("input_data")
	check(err)
	defer file.Close()

	// var schematic []string
	win_count := map[int]int{}
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	card_num := 1
	for scanner.Scan() {
		win_count[card_num] += 1
		wins := int(processLine(scanner.Text()))
		for i := 1; i <= wins; i++ {
			win_count[i+card_num] += (1 * win_count[card_num])
		}

		card_num += 1
		check(err)
	}

	sum := 0
	for _, card_num := range win_count {
		sum += card_num
	}

	fmt.Println("Part sum: " + strconv.FormatInt(int64(sum), 10))

}
