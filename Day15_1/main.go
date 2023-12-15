package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func calcHash(instruction []byte) byte {
	hash := byte(0)

	for i := range instruction {
		hash = byte((int(hash+instruction[i]) * 17) % int(256))
	}

	return hash
}

func main() {
	file, err := os.Open("input_data")
	//file, err := os.Open("examp_input")
	check(err)
	defer file.Close()
	reader := bufio.NewReader(file)

	total_hash := 0
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

		instr_hash := calcHash(instruction)
		total_hash += int(instr_hash)
		fmt.Println(string(instruction), "HASH = ", instr_hash)
	}

	fmt.Println("Total HASH = ", total_hash)
}
