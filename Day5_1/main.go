package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type map_entry struct {
	src   int64
	dest  int64
	count int64
}

func (me map_entry) in_map(src int64) bool {
	return (src >= me.src && src < (me.src+me.count))
}

func (me map_entry) convert(src int64) int64 {
	if !me.in_map(src) {
		return -1
	}

	delta := src - me.src
	return me.dest + delta
}

type ProcessMap struct {
	entries []map_entry
}

func (p ProcessMap) convert(src int64) int64 {
	for _, entry := range p.entries {
		if entry.in_map(src) {
			return entry.convert(src)
		}
	}

	return src
}

func (p ProcessMap) addMapEntry(src int64, dest int64, count int64) {
	new_entry := map_entry{src: src, dest: dest, count: count}
	p.entries = append(p.entries, new_entry)
}

func loadSeeds(line string) []int64 {
	seedList := []int64{}

	seeds := strings.Split(line, " ")
	_, seeds = seeds[0], seeds[1:]

	for _, seed_val := range seeds {
		seeds_num, err := strconv.ParseInt(seed_val, 10, 64)
		check(err)
		seedList = append(seedList, seeds_num)
	}

	return seedList
}

func chainConvert(seed int64) int64 {
	return HumiToLoc.convert(tempToHumi.convert(lightToTemp.convert(waterToLight.convert(fertToWater.convert(soilToFert.convert(seedToSoil.convert(seed)))))))
}

func processesSeeds(seedList []int64) int64 {

	lowestLoc := float64(9223372036854775807)
	for _, seed := range seedList {
		lowestLoc = math.Min(lowestLoc, float64((chainConvert(seed))))
	}
	// 	locList = append(locList, chainConvert(seed))

	return int64(lowestLoc)
}

var seedToSoil ProcessMap
var soilToFert ProcessMap
var fertToWater ProcessMap
var waterToLight ProcessMap
var lightToTemp ProcessMap
var tempToHumi ProcessMap
var HumiToLoc ProcessMap

func processMapData(lines []string) {
	mapIndex := -1
	var curMap ProcessMap
	for _, line := range lines {
		if !unicode.IsDigit(rune(line[0])) {
			mapIndex += 1
			if mapIndex == 1 {
				seedToSoil = curMap
			} else if mapIndex == 2 {
				soilToFert = curMap
			} else if mapIndex == 3 {
				fertToWater = curMap
			} else if mapIndex == 4 {
				waterToLight = curMap
			} else if mapIndex == 5 {
				lightToTemp = curMap
			} else if mapIndex == 6 {
				tempToHumi = curMap
			}

			curMap = ProcessMap{}
			continue
		}

		map_data := strings.Split(line, " ")

		if len(map_data) != 3 {
			panic("Bad Mapping data: " + line)
		}
		dest_num, err := strconv.ParseInt(map_data[0], 10, 0)
		check(err)
		src_num, err := strconv.ParseInt(map_data[1], 10, 0)
		check(err)
		count_num, err := strconv.ParseInt(map_data[2], 10, 0)
		check(err)

		new_entry := map_entry{src: src_num, dest: dest_num, count: count_num}
		curMap.entries = append(curMap.entries, new_entry)
	}

	HumiToLoc = curMap
}

func main() {
	file, err := os.Open("input_data")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// optionally, resize scanner's capacity for lines over 64K, see next example

	var seeds []int64
	var mappingData []string
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "seeds:") {
			seeds = loadSeeds(line)
			continue
		}

		mappingData = append(mappingData, line)
		check(err)
	}

	processMapData(mappingData)

	lowest := processesSeeds(seeds)

	fmt.Println("Lowest location: " + strconv.FormatInt(lowest, 10))
}
