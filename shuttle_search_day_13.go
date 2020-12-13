package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type bus_schedule struct {
	start int
	buses map[int]int
}

func parseSchedule(schedule []string) bus_schedule {
	if len(schedule) < 2 {
		fmt.Println("Schedule not long enough check input file.")
		os.Exit(1)
	}
	start, err := strconv.Atoi(schedule[0])
	if err != nil {
		log.Fatal("Unable to get bus.start from schedule[0]:", schedule[0])
		os.Exit(1)
	}
	bus := bus_schedule{start: start, buses: make(map[int]int)}

	for offset, bus_id := range strings.Split(schedule[1], ",") {
		if bus_id != "x" {
			i, err := strconv.Atoi(bus_id)
			if err != nil {
				log.Print("Unable to convert bus_id:", bus_id)
				continue
			}
			bus.buses[i] = offset
		}
	}

	return bus
}

func readInputFile() []string {
	file_name := "input/day_13_input.txt"
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

func calcPart1(bus bus_schedule) int {
	time := math.MaxInt32 //int((^uint(0)) >> 1)
	min_bus := 0

	for bus_id, _ := range bus.buses {
		offset := bus.start % bus_id
		if offset == 0 {
			return bus_id
		}
		wait_time := bus_id - offset
		if wait_time < time {
			time = wait_time
			min_bus = bus_id
		}
	}

	return min_bus * time
}

func calcPart2(bus bus_schedule) int64 {
	increment := int64(0)
	offset := int64(0)

	for b, o := range bus.buses {
		if increment == 0 {
			increment = int64(b)
			//offset = increment - int64(o)
			log.Print("Starting at ", b, " with offset ", o)
			continue
		}

		for n := int64(0); ; n++ {
			tmp := n*increment + int64(o) + offset
			if tmp%int64(b) == 0 {
				offset += n * increment
				increment *= int64(b)
				break
			}
		}
	}

	return offset
}

func main() {
	schedule := readInputFile()
	bus := parseSchedule(schedule)
	part1 := calcPart1(bus)

	fmt.Println("Earliest bus id times wait time:", part1)
	//Earliest bus id times wait time: 207

	part2 := calcPart2(bus)
	fmt.Println("Earliest timestamp with departing offsets:", part2)
	//Earliest timestamp with departing offsets: 530015546283687
}
