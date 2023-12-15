package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	orderedmap "github.com/iancoleman/orderedmap"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var box_map map[int]*orderedmap.OrderedMap

func calcHash(instruction []byte) byte {
	hash := byte(0)
	for i := range instruction {
		hash = byte((int(hash+instruction[i]) * 17) % int(256))
	}
	return hash
}

func processInstruction(instruction []byte, boxes map[int]*orderedmap.OrderedMap) byte {

	instr_str := string(instruction)
	add_index := strings.Index(instr_str, "=")
	sub_index := strings.Index(instr_str, "-")
	if add_index != -1 {
		short_instr := instr_str[:add_index]
		lens_val, err := strconv.ParseInt(instr_str[add_index+1:], 10, 0)
		check(err)
		fmt.Println("Adding", short_instr, "=", lens_val)
		addLens(short_instr, int(lens_val), boxes)
	} else if sub_index != -1 {
		short_instr := instr_str[:sub_index]
		fmt.Println("Removing", short_instr, "=")
		rmLens(short_instr, boxes)
	}

	return calcHash(instruction)
}

func addLens(short_instr string, lens_val int, boxes map[int]*orderedmap.OrderedMap) {
	box_hash := int(calcHash([]byte(short_instr)))
	box, ok := boxes[box_hash]
	if !ok {
		boxes[box_hash] = orderedmap.New()
		box = boxes[box_hash]
	}

	box.Set(short_instr, lens_val)
}

func rmLens(short_instr string, boxes map[int]*orderedmap.OrderedMap) {
	box_hash := int(calcHash([]byte(short_instr)))
	box, ok := boxes[box_hash]
	if !ok {
		boxes[box_hash] = orderedmap.New()
		box = boxes[box_hash]
	}

	box.Delete(short_instr)
}

func calcLens(box_map map[int]*orderedmap.OrderedMap) int {
	sum := 0
	for index, box := range box_map {
		box_sum := 0
		for i, key := range box.Keys() {
			value, _ := box.Get(key)
			box_sum += (index + 1) * (i + 1) * value.(int)
		}
		sum += box_sum
	}

	return sum
}

func main() {
	file, err := os.Open("input_data")
	//file, err := os.Open("examp_input")
	check(err)
	defer file.Close()
	reader := bufio.NewReader(file)

	total_hash := 0
	box_map := map[int]*orderedmap.OrderedMap{}
	for done := false; !done; {
		instruction, err := reader.ReadBytes(',')
		if err != nil {
			if errors.Is(err, io.EOF) { // prefered way by GoLang doc
				done = true
			} else {
				check(err)
			}
		}
		if len(instruction) > 0 && instruction[len(instruction)-1] == ',' {
			instruction = instruction[:len(instruction)-1]
		}

		instr_hash := processInstruction(instruction, box_map)
		total_hash += int(instr_hash)

	}
	lens_calc := calcLens(box_map)

	fmt.Println("Total HASH =", total_hash)
	fmt.Println("Lens Sum =", lens_calc)
}
