package main

import (
	"io/ioutil"
	"os"
	"fmt"
	"log"
	"strconv"
	"strings"
	"regexp"
)


func parseLine(line string) bool {
	re := regexp.MustCompile(`^(?P<start>\d+)-(?P<end>\d+) (?P<str>\w+): (?P<pass>\w+)$`)
	m :=	re.FindStringSubmatch(strings.TrimSpace(line))
	if len(m) < 5 {
		log.Print("Failed to match: ", m)
		return false
	}

	//log.Print("m[0]:", m, " m[1]:", m[1], " m[2]:", m[2], " m[3]:", m[3])
	sub_re := regexp.MustCompile(m[3])
	count := len(sub_re.FindAllStringIndex(m[4], -1))
	start, err := strconv.Atoi(m[1])
	if err != nil {
		return false
	}

	end, err := strconv.Atoi(m[2])
	if err != nil {
		return false
	}
	
	if count < start || count > end {
		return false
	}

	return true
}

func parseLinePart2 (line string) bool {
	re := regexp.MustCompile(`^(?P<start>\d+)-(?P<end>\d+) (?P<str>\w+): (?P<pass>\w+)$`)
	m :=	re.FindStringSubmatch(strings.TrimSpace(line))
	if len(m) < 5 {
		log.Print("Failed to match: ", m)
		return false
	}

	start, err := strconv.Atoi(m[1])
	if err != nil {
		return false
	}

	end, err := strconv.Atoi(m[2])
	if err != nil {
		return false
	}
	
	is_first := string(m[4])[start-1] != m[3][0]
	is_second := string(m[4])[end-1] != m[3][0]
	if (is_first || is_second) && !(is_first && is_second)  {
		return true 
	}

	return false
}

func readInputFile() []string {
	file_name := "input/day_2_input.txt"
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

func main () {
	line_array := readInputFile()
	correct := 0
	correct_part2 := 0

	for _, line := range line_array {
		if parseLine(line) {
			correct++
		}	

		if parseLinePart2(line) {
			correct_part2++
		} else {
			log.Print("Incorrect line part 2: ", line, "\n")
		}
	}

	fmt.Println("Correct lines: ", correct)
	fmt.Println("Correct lines part 2: ", correct_part2)

	//Correct lines:  548
	//Correct lines part 2:  502
}
