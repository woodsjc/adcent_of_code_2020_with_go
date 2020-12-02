package main

import (
	"io/ioutil"
	"os"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func calcMultiple(num_list []int, target int, ignore int) int64 {
	val := make(map[int]bool)

	for i:=0; i<len(num_list); i++ {
		if ignore == i {
			continue
		}
		pair := target - num_list[i]
		
		if _, in_map := val[pair]; in_map {
			return int64(num_list[i]) * int64(pair)
		}
		
		val[num_list[i]] = true
	}

	return 0
}

func calcTripleMultiple(num_list []int, target int) int64 {
	for i:=0; i<len(num_list); i++ {
		pair_total := target - num_list[i]

		if tmp := calcMultiple(num_list, pair_total, i); tmp != 0 {
			return tmp * int64(num_list[i])
		}
	}

	return 0
}

func main () {
	target := 2020
	file_name := "input/day_1_input.txt"
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	to_parse := strings.Split(string(data), "\n")

	num_array := make([]int, len(to_parse))
	for i, token := range to_parse {
		num, _ := strconv.Atoi(strings.TrimSpace(token))
		num_array[i] = num
	}

	fmt.Println("Two entries summing to ",
		target, " when multiplied: ", calcMultiple(num_array, target, -1))
	fmt.Println("Three entries summing to ",
		target, " when multiplied: ", calcTripleMultiple(num_array, target))

	 //Two entries summing to 2020 when multiplied:  355875
	 //Three entries summing to 2020 when multiplied:  140379120
}

