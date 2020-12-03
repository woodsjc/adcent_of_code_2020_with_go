package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func parseLine(line string) []bool {
	line = strings.TrimSpace(line)
	result := make([]bool, len(line))

	for i, c := range line {
		if c == '.' {
			result[i] = false
		} else {
			result[i] = true
		}
	}

	return result
}

func readInputFile() []string {
	file_name := "input/day_3_input.txt"
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return strings.Split(string(data), "\n")
}

func calcTrees(cur_x int, cur []bool, slope_x int) int {
	total := 0

	if cur[(cur_x+slope_x)%len(cur)] {
		total++
	}
	return total
}

func main() {
	line_array := readInputFile()
	trees_encountered := []int{0, 0, 0, 0, 0}
	position := []int{0, 0, 0, 0, 0}
	slope := []int{1, 3, 5, 7, 1}

	for i, line := range line_array {
		cur_row := parseLine(line)
		if i == 0 || len(cur_row) == 0 {
			continue
		}

		for j := 0; j < 5; j++ {
			if j == 4 && i%2 == 1 {
				continue
			}
			trees_encountered[j] += calcTrees(position[j], cur_row, slope[j])
			position[j] = (position[j] + slope[j]) % len(cur_row)
		}
	}

	super_total := int64(1)
	for j := 0; j < 5; j++ {
		fmt.Println("Trees encountered-", trees_encountered[j], " with slope-", slope[j])
		super_total *= int64(trees_encountered[j])
	}
	//Trees encountered- 198 with slope- 3

	fmt.Println("Multiplied total: ", super_total)
	//Trees encountered- 84  with slope- 1
	//Trees encountered- 198  with slope- 3
	//Trees encountered- 72  with slope- 5
	//Trees encountered- 81  with slope- 7
	//Trees encountered- 53  with slope- 1
	//Multiplied total:  5140884672
}
