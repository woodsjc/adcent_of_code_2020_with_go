package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func parseGroup(passport string) int {
	answers := make(map[rune]bool)

	for _, l := range strings.Split(passport, "\n") {
		l = strings.TrimSpace(l)
		for _, c := range l {
			answers[c] = true
		}
	}

	return len(answers)
}

func parseGroupPart2(passport string) int {
	answers := make(map[rune]bool)

	for i, l := range strings.Split(passport, "\n") {
		l = strings.TrimSpace(l)
		if len(l) <= 0 {
			continue
		}

		//log.Print("line", i, "-", l)
		if i == 0 {
			for _, c := range l {
				answers[c] = true
			}
			continue
		}

		for k, _ := range answers {
			if strings.Contains(l, string(k)) {
				continue
			} else {
				//log.Print("deleted", k, "from", answers)
				delete(answers, k)
			}
		}
	}

	return len(answers)
}

func readInputFile() []string {
	file_name := "input/day_6_input.txt"
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
	return strings.Split(string(data), "\n\n") //passports separated by double line
}

func main() {
	pass_array := readInputFile()
	valid := 0
	valid2 := 0

	for _, p := range pass_array {
		valid += parseGroup(p)
		valid2 += parseGroupPart2(p)
	}

	fmt.Println("Sum of group responses:", valid)
	//Sum of group responses: 6542

	fmt.Println("Part 2 - Number of group questions answered:", valid2)
	//Part 2 - Number of group questions answered: 3299
}
