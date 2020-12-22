package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type BagTable struct {
	Name []string
	Num  []int
}

var all_counts = make(map[string]int)

func contains(arr []string, sub string) bool {
	for _, a := range arr {
		if a == sub {
			return true
		}
	}
	return false
}

func readInputFile() []string {
	file_name := "input/day_7_input.txt"
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

func parseBags(input []string) map[string]BagTable {
	//match groups [0]-full match | [1]-num | [2]-name
	container_re := regexp.MustCompile(`^\w+\s\w+`)
	re := regexp.MustCompile(`(\d{1,5})\s(\w+\s\w+)\sbag`)
	bags := make(map[string]BagTable)

	for _, line := range input {
		bag_name := string(container_re.Find([]byte(line)))
		if len(bag_name) <= 0 {
			log.Println("bag_name no length on line:", line)
			continue
		}

		m := re.FindAllStringSubmatch(line, -1)
		if len(m) <= 0 {
			log.Println("Unable to find substring in line:", line)
			continue
		}
		entry_bag := make([]string, len(m))
		entry_num := make([]int, len(m))

		for i, bag := range m {
			tmp, err := strconv.Atoi(bag[1])
			if err != nil {
				continue
			}

			entry_bag[i] = bag[2]
			entry_num[i] = tmp
		}
		bags[bag_name] = BagTable{Name: entry_bag, Num: entry_num}
	}

	return bags
}

func calcPart1(bags map[string]BagTable) int {
	//total := 0
	target := "shiny gold"
	map_to_target := make(map[string]bool)

	for i := len(map_to_target); ; i = len(map_to_target) {
		for bag_name, bt := range bags {
			bag_array := bt.Name
			if i == 0 && contains(bag_array, target) {
				map_to_target[bag_name] = true
			} else if i != 0 {
				for _, sub_bag := range bag_array {
					if map_to_target[sub_bag] {
						map_to_target[bag_name] = true
						break
					}
				}
			}
		}
		if i == len(map_to_target) {
			break
		}
	}

	return len(map_to_target)
}

func sumBags(bags map[string]BagTable, sum map[string]int) map[string]int {
bag_loop:
	for key, bag := range bags {
		if _, ok := sum[key]; ok {
			continue
		}
		total := 0
		for i := 0; i < len(bag.Name); i++ {
			if _, ok := sum[bag.Name[i]]; !ok {
				if _, ok := bags[bag.Name[i]]; ok {
					continue bag_loop
				} else {
					sum[bag.Name[i]] = 0
				}
			}
			total += sum[bag.Name[i]]*bag.Num[i] + bag.Num[i]
		}
		sum[key] = total
	}

	return sum
}

func calcPart2(bags map[string]BagTable, target string) int {
	sum := sumBags(bags, make(map[string]int))
	for i, count := 0, 0; ; i, count = len(sum), count+1 {
		if i, ok := sum[target]; ok {
			log.Println("Found target on ", count, " iteration.")
			return i
		}
		sum = sumBags(bags, sum)
		if i > 0 && i == len(sum) {
			break
		}
		if count%1000 == 0 {
			log.Println("Count: ", count)
		}
	}

	return sum[target]
}

func main() {
	input := readInputFile()
	bags := parseBags(input)
	total_part1 := calcPart1(bags)

	fmt.Println("number of bags that can hold shiny gold bags:", total_part1)
	//Number of bags that can hold shiny gold bags: 128

	total_part2 := calcPart2(bags, "shiny gold")
	fmt.Println("number of bags contained in a shiny gold bag:", total_part2)
	//
}
