package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func calcValues(time int64, distance int64) int64 {
	wins := int64(0)
	for i := int64(0); i <= time; i++ {
		d := time*i - i*i
		if d > distance {
			wins += 1
		}
	}

	return wins
}

func main() {
	file, err := os.Open("input_data")
	// file, err := os.Open("examp_input")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// optionally, resize scanner's capacity for lines over 64K, see next example

	times := []int64{}
	distances := []int64{}
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "Time:") {
			time_strings := strings.Split(line, " ")
			_, time_strings = time_strings[0], time_strings[1:]

			for i := 0; i < len(time_strings); i += 1 {
				if len(time_strings[i]) == 0 {
					continue
				}
				time, err := strconv.ParseInt(time_strings[i], 10, 64)
				check(err)
				times = append(times, time)
			}
			continue
		}

		if strings.HasPrefix(line, "Distance:") {
			distance_strings := strings.Split(line, " ")
			_, distance_strings = distance_strings[0], distance_strings[1:]

			for i := 0; i < len(distance_strings); i += 1 {
				if len(distance_strings[i]) == 0 {
					continue
				}
				distance, err := strconv.ParseInt(distance_strings[i], 10, 64)
				check(err)
				distances = append(distances, distance)
			}
			continue
		}
	}

	product := int64(1)
	for i := range times {
		product *= calcValues(times[i], distances[i])
	}

	fmt.Println("Margin of errors: " + strconv.FormatInt(product, 10))
}
