package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strings"
)

func parseBoarding(boarding string) int {
	row := 0
	col := 0
	boarding = strings.TrimSpace(boarding)

	for i := 0; i <= 6; i++ {
		if boarding[i] == "B"[0] {
			row += int(math.Pow(2, float64(6-i)))
			//log.Print("Updated row:", row, " boarding[", i, "]:", boarding[i])
		}
	}

	for i := 7; i <= 9; i++ {
		if boarding[i] == "R"[0] {
			col += int(math.Pow(2, float64(9-i)))
		}
	}
	//log.Print("Boarding:", boarding, " Row:", row, " Col:", col, " Return:", row*8+col)
	return row*8 + col
}

func readInputFile() []string {
	file_name := "input/day_5_input.txt"
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

func main() {
	pass_array := readInputFile()
	valid := 0
	valid2 := 0
	all_seats := make([]bool, 127*8+7)

	for i, _ := range all_seats {
		all_seats[i] = true
	}

	for _, p := range pass_array {
		if len(p) <= 0 {
			continue
		}
		tmp := parseBoarding(p)
		if tmp > valid {
			valid = tmp
		}
		all_seats[tmp] = false
	}

	for i, is_seat := range all_seats {
		if is_seat && i > 0 && !all_seats[i-1] && !all_seats[i+1] {
			valid2 = i
			break
		}
	}

	fmt.Println("Highest seat id:", valid)
	//Highest seat id: 826
	fmt.Println("Part 2 - Missing seat is:", valid2)
	//Part 2 - Missing seat is: 678
}
