package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func symbolAdject(schematic []string, line_num int, char_num int, count int) bool {

	for i := math.Max(0, float64(line_num-1)); i < math.Min(float64(len(schematic)), float64(line_num+2)); i++ {
		for j := math.Max(0, float64(char_num-1)); j < math.Min(float64(len(schematic[int(i)])), float64(char_num+count+1)); j++ {
			if schematic[int(i)][int(j)] == '.' {
				continue
			} else if unicode.IsDigit(rune(schematic[int(i)][int(j)])) {
				continue
			} else {
				return true
			}
		}
	}

	return false
}

func processSchematic(schematic []string) int64 {
	sum := int64(0)

	for line_num, line := range schematic {
		num_val := 0
		num_len := 0
		for char_num, char := range line {
			if unicode.IsDigit(char) {
				num_val *= 10
				num_val += int(char - '0')
				num_len += 1

				if char_num == (len(line) - 1) {
					// Number is at the end...
					// we need to look for symbol now
					if symbolAdject(schematic, line_num, char_num-num_len+1, num_len) {
						sum += int64(num_val)
					}
					num_val = 0
					num_len = 0
				}
				continue
			} else if num_val > 0 {
				if symbolAdject(schematic, line_num, char_num-num_len, num_len) {
					sum += int64(num_val)
				}

				num_val = 0
				num_len = 0
			}
		}
	}

	return sum
}

func main() {
	file, err := os.Open("input_data")
	check(err)
	defer file.Close()

	var schematic []string

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	for scanner.Scan() {
		schematic = append(schematic, scanner.Text())
		check(err)
	}

	fmt.Println("Part sum: " + strconv.FormatInt(processSchematic(schematic), 10))

}
