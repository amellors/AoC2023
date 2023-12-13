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

func calcNext(data []int64) int64 {
	data_deltas := []int64{}
	all_zeros := true

	for i := range data {
		if i == len(data)-1 {
			continue
		}
		delta := data[i+1] - data[i]
		if delta != 0 {
			all_zeros = false
		}

		data_deltas = append(data_deltas, delta)
	}

	var ret_sum int64
	if all_zeros {
		return data[0]
	} else {
		ret_sum = calcNext(data_deltas)
	}

	return ret_sum + data[len(data)-1]
}

func calcPrev(data []int64) int64 {
	data_deltas := []int64{}
	all_zeros := true

	for i := range data {
		if i == len(data)-1 {
			continue
		}
		delta := data[i+1] - data[i]
		if delta != 0 {
			all_zeros = false
		}

		data_deltas = append(data_deltas, delta)
	}

	var ret_sum int64
	if all_zeros {
		return data[0]
	} else {
		ret_sum = calcPrev(data_deltas)
	}

	retval := data[0] - ret_sum
	return retval
}

func main() {
	file, err := os.Open("input_data")
	//file, err := os.Open("examp_input")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	dataLists := [][]int64{}

	for scanner.Scan() {
		map_line := scanner.Text()
		if len(map_line) == 0 {
			continue
		}
		hist_data := []int64{}

		num_strs := strings.Split(map_line, " ")
		for _, num_str := range num_strs {
			val, err := strconv.ParseInt(num_str, 10, 0)
			check(err)
			hist_data = append(hist_data, val)
		}

		dataLists = append(dataLists, hist_data)
	}

	sum := int64(0)
	for _, data := range dataLists {
		sum += calcNext(data)
	}

	fmt.Println("Part 1 sum: ", sum)

	sum = 0
	for _, data := range dataLists {
		sum += calcPrev(data)
	}

	fmt.Println("Part 2 sum: ", sum)

}
