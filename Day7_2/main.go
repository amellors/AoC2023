package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getCardVal(card byte) int {
	if card == '2' {
		return 2
	} else if card == '3' {
		return 3
	} else if card == '4' {
		return 4
	} else if card == '5' {
		return 5
	} else if card == '6' {
		return 6
	} else if card == '7' {
		return 7
	} else if card == '8' {
		return 8
	} else if card == '9' {
		return 9
	} else if card == 'T' {
		return 10
	} else if card == 'J' {
		return 1
	} else if card == 'Q' {
		return 12
	} else if card == 'K' {
		return 13
	} else if card == 'A' {
		return 14
	} else {
		panic("invalid card!")
	}
}

func getHandType(hand string) int {
	cardCount := map[rune]int{}
	origCardCount := map[rune]int{}
	for _, card := range hand {
		origCardCount[card] += 1
	}

	// Calculate Js
	jVal := 0
	card_max := 0
	card_key_max := rune(0)
	for k, v := range origCardCount {
		if k == 'J' {
			jVal = v
		} else {
			if v > card_max {
				card_max = v
				card_key_max = k
			}
			cardCount[k] = v
		}
	}

	if jVal > 0 {
		cardCount[card_key_max] += jVal
	}

	if len(cardCount) == 5 {
		// High Card
		return 1
	}

	if len(cardCount) == 4 {
		// One Pair
		return 2
	}

	if len(cardCount) == 3 {
		// Two Pair Or 3 of a kind
		for _, num := range cardCount {
			if num == 3 {
				return 4
			}
		}
		return 3
	}

	if len(cardCount) == 2 {
		// Full house or 4 of a Kind
		for _, num := range cardCount {
			if num == 4 {
				// 4 of a kind
				return 6
			}
		}
		// full house
		return 5
	}

	// One card == Five of a kind
	return 7
}

func main() {
	file, err := os.Open("input_data")
	// file, err := os.Open("examp_input")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// optionally, resize scanner's capacity for lines over 64K, see next example

	handList := map[int]int64{}
	for scanner.Scan() {
		line := scanner.Text()

		line_strings := strings.Split(line, " ")

		// process hand
		if len(line_strings[0]) != 5 {
			panic("Invalid Hand length")
		}
		// 		hand_value :=
		// for i := 0; i < len(line_strings[0]); i += 1 {
		hand_value := getHandType(line_strings[0])<<20 +
			getCardVal(line_strings[0][0])<<16 +
			getCardVal(line_strings[0][1])<<12 +
			getCardVal(line_strings[0][2])<<8 +
			getCardVal(line_strings[0][3])<<4 +
			getCardVal(line_strings[0][4])

		bid_val, err := strconv.ParseInt(line_strings[1], 10, 64)
		handList[hand_value] = bid_val
		check(err)
	}

	keys := make([]int, 0, len(handList))

	for k := range handList {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	sum := int64(0)
	for i, k := range keys {
		sum += int64(i+1) * handList[k]
	}

	// product := int64(1)
	/*for i := range times {
		product *= calcValues(times[i], distances[i])
	}*/

	fmt.Println("Sum of winnings: ", sum)
}
