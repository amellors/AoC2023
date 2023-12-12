package main

import (
	"bufio"
	"fmt"
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

type MatchType int64

const (
	PARTIAL_MATCH MatchType = iota
	NO_MATCH
	FULL_MATCH
)

func (me map_entry) range_in_map(start int64, count int64) MatchType {
	endVal := me.src + me.count - 1
	range_end := start + count - 1
	if start >= me.src && start <= endVal && range_end <= endVal {
		return FULL_MATCH
	}

	if start >= me.src && start <= endVal {
		return PARTIAL_MATCH
	}

	if range_end >= me.src && range_end <= endVal {
		return PARTIAL_MATCH
	}

	return NO_MATCH
}

func (me map_entry) convert(src int64) int64 {
	if !me.in_map(src) {
		return -1
	}

	delta := src - me.src
	return me.dest + delta
}

func (me map_entry) convertRange(start int64, count int64) (int64, int64) {
	delta := start - me.src
	var leftover int64
	if (delta + count) > me.count {
		leftover = delta + count - me.count
	}
	return (me.dest + delta), leftover
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

func (p ProcessMap) convertRange(start int64, count int64) map[int64]int64 {
	convertedRanges := map[int64]int64{}

	remainder := count
	start_val := start

	for remainder > 0 {
		new_remainder := remainder
		for _, entry := range p.entries {
			if entry.range_in_map(start_val, remainder) == FULL_MATCH {
				convert := entry.convert(start_val)
				convertedRanges[convert] = count
				new_remainder = 0
				break
			} else if entry.range_in_map(start_val, remainder) == PARTIAL_MATCH {
				if entry.in_map(start) {
					convert, ret_remainder := entry.convertRange(start_val, remainder)
					convertedRanges[convert] = remainder - ret_remainder
					new_remainder = ret_remainder
					start_val = start_val + (remainder - ret_remainder)
					break
				} else {
					// start is not in range, but end is...
					for i := int64(1); i <= remainder; i++ {
						if entry.in_map(start_val + i) {
							convert, ret_remainder := entry.convertRange(start_val+i, remainder-i)
							convertedRanges[convert] = remainder - i - ret_remainder
							new_remainder = i + ret_remainder
							break
						}
					}
				}
			}
		}

		if new_remainder == remainder {
			convertedRanges[start_val] = remainder
			new_remainder = 0
		}
		remainder = new_remainder
	}

	return convertedRanges
}

func (p ProcessMap) addMapEntry(src int64, dest int64, count int64) {
	new_entry := map_entry{src: src, dest: dest, count: count}
	p.entries = append(p.entries, new_entry)
}

func loadSeeds(line string) map[int64]int64 {
	seedList := map[int64]int64{}

	seeds := strings.Split(line, " ")
	_, seeds = seeds[0], seeds[1:]

	for i := 0; i < len(seeds); i += 2 {
		seeds_start, err := strconv.ParseInt(seeds[i], 10, 64)
		check(err)
		seeds_count, err := strconv.ParseInt(seeds[i+1], 10, 64)
		check(err)
		seedList[seeds_start] = seeds_count
	}

	return seedList
}

func chainConvert(seed int64) int64 {
	return HumiToLoc.convert(tempToHumi.convert(lightToTemp.convert(waterToLight.convert(fertToWater.convert(soilToFert.convert(seedToSoil.convert(seed)))))))
}

func chainConvertRange(seeds map[int64]int64) map[int64]int64 {
	convertedList := map[int64]int64{}

	soils := map[int64]int64{}
	for start, count := range seeds {
		for sstart, scount := range seedToSoil.convertRange(start, count) {
			soils[sstart] = scount
		}
	}

	ferts := map[int64]int64{}
	for start, count := range soils {
		for sstart, scount := range soilToFert.convertRange(start, count) {
			ferts[sstart] = scount
		}
	}

	waters := map[int64]int64{}
	for start, count := range ferts {
		for sstart, scount := range fertToWater.convertRange(start, count) {
			waters[sstart] = scount
		}
	}

	lights := map[int64]int64{}
	for start, count := range ferts {
		for sstart, scount := range waterToLight.convertRange(start, count) {
			lights[sstart] = scount
		}
	}

	temps := map[int64]int64{}
	for start, count := range ferts {
		for sstart, scount := range lightToTemp.convertRange(start, count) {
			temps[sstart] = scount
		}
	}

	humis := map[int64]int64{}
	for start, count := range ferts {
		for sstart, scount := range tempToHumi.convertRange(start, count) {
			humis[sstart] = scount
		}
	}

	for start, count := range ferts {
		for sstart, scount := range HumiToLoc.convertRange(start, count) {
			convertedList[sstart] = scount
		}
	}

	return convertedList
}

func processesSeeds(seedList map[int64]int64) int64 {

	lowestLoc := float64(9223372036854775807)
	chainConvertRange(seedList)
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
	// file, err := os.Open("input_data")
	file, err := os.Open("examp_input")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// optionally, resize scanner's capacity for lines over 64K, see next example

	var seeds map[int64]int64
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
