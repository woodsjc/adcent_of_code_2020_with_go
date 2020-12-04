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

func parsePassport(passport string) bool {
	valid := []bool{false, false, false, false, false, false, false}

	for _, l := range strings.Split(passport, "\n") {
		entries := strings.Split(l, " ")

		for _, e := range entries {
			e = strings.TrimSpace(e)
			if len(e) < 3 {
				continue
			}
			e = e[:3]

			if e == "byr" {
				valid[0] = true
			} else if e == "iyr" {
				valid[1] = true
			} else if e == "eyr" {
				valid[2] = true
			} else if e == "hgt" {
				valid[3] = true
			} else if e == "hcl" {
				valid[4] = true
			} else if e == "ecl" {
				valid[5] = true
			} else if e == "pid" {
				valid[6] = true
			} else if e == "cid" {
				continue
			} else {
				log.Print("Invalid entry:", e)
			}
		}
	}

	for _, v := range valid {
		if !v {
			return false
		}
	}
	return true
}

func parsePassport2(passport string) bool {
	valid := []bool{false, false, false, false, false, false, false}
	hair_re := regexp.MustCompile("^#[0-9a-f]{6}$")
	eye_re := regexp.MustCompile("^:(amb|blu|brn|gry|grn|hzl|oth)$")
	pid_re := regexp.MustCompile("^:[0-9]{9}$")

	for _, l := range strings.Split(passport, "\n") {
		entries := strings.Split(l, " ")

		for _, e := range entries {
			e = strings.TrimSpace(e)
			if len(e) < 3 {
				continue
			}
			data := string(e)
			e = e[:3]

			if e == "byr" {
				data = data[3:]
				if string(data[0]) == ":" {
					d, err := strconv.Atoi(data[1:])
					if err == nil && d <= 2002 && d >= 1920 {
						valid[0] = true
					}
				}
			} else if e == "iyr" {
				data = data[3:]
				if string(data[0]) == ":" {
					d, err := strconv.Atoi(data[1:])
					if err == nil && d <= 2020 && d >= 2010 {
						valid[1] = true
					}
				}
			} else if e == "eyr" {
				data = data[3:]
				if string(data[0]) == ":" {
					d, err := strconv.Atoi(data[1:])
					if err == nil && d <= 2030 && d >= 2020 {
						valid[2] = true
					}
				}
			} else if e == "hgt" {
				data = data[3:]
				if string(data[0]) == ":" {
					if data[len(data)-2:] == "cm" {
						d, err := strconv.Atoi(data[1 : len(data)-2])
						if err == nil && d <= 193 && d >= 150 {
							valid[3] = true
						}
					} else if data[len(data)-2:] == "in" {
						d, err := strconv.Atoi(data[1 : len(data)-2])
						if err == nil && d <= 76 && d >= 59 {
							valid[3] = true
						}
					}
				}
			} else if e == "hcl" {
				data = data[3:]
				if string(data[0]) == ":" {
					if hair_re.MatchString(data[1:]) {
						valid[4] = true
					}
				}
			} else if e == "ecl" {
				data = data[3:]
				if eye_re.MatchString(data) {
					valid[5] = true
				}
			} else if e == "pid" {
				data = data[3:]
				if pid_re.MatchString(data) {
					valid[6] = true
				}
			} else if e == "cid" {
				continue
			} else {
				log.Print("Invalid entry:", e)
			}
		}
	}

	for _, v := range valid {
		if !v {
			return false
		}
	}
	return true
}

func readInputFile() []string {
	file_name := "input/day_4_input.txt"
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
		if parsePassport(p) {
			valid++
		}
		if parsePassport2(p) {
			valid2++
		}
	}

	fmt.Println("Number of valid passports", valid)
	//part1 --- Number of valid passports 208

	fmt.Println("Part 2 - Number of valid passports", valid2)
	//part2 ---- Part 2 - Number of valid passports 167
}
