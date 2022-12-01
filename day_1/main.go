package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func ParseInput() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	max_found := 0
	subtotal := 0
	elf_loads := []int{}
	scanner := bufio.NewScanner(file)
	line := ""
	for scanner.Scan() {
		line = scanner.Text() // I don't think i'm getting the very last elf.
		if line == "" {
			if subtotal > max_found {
				max_found = subtotal
			}
			elf_loads = append(elf_loads, subtotal)
			subtotal = 0
			continue
		}
		value, _ := strconv.Atoi(line)
		subtotal += value
	}
	sort.Ints(elf_loads)
	top_3 := elf_loads[len(elf_loads)-3:]
	fmt.Println(max_found)
	fmt.Println(top_3[0] + top_3[1] + top_3[2])
}

func main() {
	ParseInput()
}
