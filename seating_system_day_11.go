package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func parseGrid(grid []string) [][]rune {
	seats := make([][]rune, len(grid)-1)

    for i, line := range grid[:len(grid)-1] {
		if len(line) <= 0 {
			continue
		}
		seats[i] = make([]rune, len(line))
		for j, c := range line {
			seats[i][j] = c
		}
        //log.Print(line)
	}

	return seats
}

func readInputFile() []string {
	file_name := "input/day_11_input.txt"
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

func applyIteration(seats [][]rune) ([][]rune, bool) {
    new_seats := make([][]rune, len(seats))
    var is_done bool = true

    for i:=0; i<len(seats); i++ {
        new_seats[i] = make([]rune, len(seats[i]))
        for j:=0; j<len(new_seats[i]); j++ {
            total_around := 0

            for ii:=i-1; ii<=i+1; ii++ {
                if ii<0 || ii>=len(seats) {
                    continue
                }
                for jj:=j-1; jj<=j+1; jj++ {
                    if jj<0 || jj>=len(seats[ii]) || (i==ii && j==jj) {
                        continue
                    }
                    if seats[ii][jj] == '#' {
                        total_around++
                    }
                }
            }

            if seats[i][j] == 'L' {
                //empty logic
                if total_around == 0 {
                    new_seats[i][j] = '#'
                    is_done = false
                } else {
                    new_seats[i][j] = 'L'
                }
            } else if seats[i][j] == '#' {
                //occupied logic
                if total_around >= 4 {
                    new_seats[i][j] = 'L'
                    is_done = false
                } else {
                    new_seats[i][j] = '#'
                }
            } else if seats[i][j] == '.' {
                new_seats[i][j] = '.'
            } else {
                //unhandled
                log.Print("Current seat[", i, "][", j, "] = '", seats[i][j], "' is unhandled.")
            }
        }
    }

    return new_seats, is_done
}

func checkUnoccupied(seats [][]rune, i int, j int) int {
    //empty seats (L) also block view of occupied steats (#)
    total := 0

    //check left
    for jj:=j-1; jj>=0; jj-- {
        if seats[i][jj] == '#' {
            total++
            break
        } else if seats[i][jj] == 'L' {
            break
        }
    }

    //check right
    for jj:=j+1; jj<len(seats[i]); jj++ {
        if seats[i][jj] == '#' {
            total++
            break
        } else if seats[i][jj] == 'L' {
            break
        }
    }

    //check up
    for ii:=i-1; ii>=0; ii-- {
        if seats[ii][j] == '#' {
            total++
            break
        } else if seats[ii][j] == 'L' {
            break
        }
    }

    //check down 
    for ii:=i+1; ii<len(seats); ii++ {
        if seats[ii][j] == '#' {
            total++
            break
        } else if seats[ii][j] == 'L' {
            break
        }
    }

    //check \
    for n:=-1; i+n>=0 && j+n>=0; n-- {
        if seats[i+n][j+n] == '#' {
            total++
            break
        } else if seats[i+n][j+n] == 'L' {
            break
        }
    }
    for n:=1; i+n<len(seats) && j+n<len(seats[i+n]); n++ {
        if seats[i+n][j+n] == '#' {
            total++
            break
        } else if seats[i+n][j+n] == 'L' {
            break
        }
    }

    //check /
    for n:=1; i+n<len(seats) && j-n>=0; n++ {
        if seats[i+n][j-n] == '#' {
            total++
            break
        } else if seats[i+n][j-n] == 'L' {
            break
        }
    }
    for n:=1; i-n>=0 && j+n<len(seats[i-n]); n++ {
        if seats[i-n][j+n] == '#' {
            total++
            break
        } else if seats[i-n][j+n] == 'L' {
            break
        }
    }

    return total
}

func applyIterationPart2(seats [][]rune) ([][]rune, bool) {
    new_seats := make([][]rune, len(seats))
    var is_done bool = true

    for i:=0; i<len(seats); i++ {
        new_seats[i] = make([]rune, len(seats[i]))
        for j:=0; j<len(new_seats[i]); j++ {
            total_around := checkUnoccupied(seats, i, j)

            if seats[i][j] == 'L' {
                //empty logic
                if total_around == 0 {
                    new_seats[i][j] = '#'
                    is_done = false
                } else {
                    new_seats[i][j] = 'L'
                }
            } else if seats[i][j] == '#' {
                //occupied logic
                if total_around >= 5 {
                    new_seats[i][j] = 'L'
                    is_done = false
                } else {
                    new_seats[i][j] = '#'
                }
            } else if seats[i][j] == '.' {
                new_seats[i][j] = '.'
            } else {
                //unhandled
                log.Print("Current seat[", i, "][", j, "] = '", seats[i][j], "' is unhandled.")
            }
        }
    }

    return new_seats, is_done
}

func countOccupied(seats [][]rune) int {
    total := 0

    for i:=0; i<len(seats); i++ {
        for j:=0; j<len(seats[i]); j++ {
            if seats[i][j] == '#' {
                total++
            }
            fmt.Print(string(seats[i][j]))
        }
        fmt.Print("\n")
    }

    return total
}


func main() {
	grid := readInputFile()
    seats := parseGrid(grid)
    part_2_seats := seats
    done := false
	iterations := 0

    for {
        seats, done = applyIteration(seats)
		iterations++
        if done {
            break
        }
	}

    fmt.Println("Total iterations until stabilized:", iterations, "--- seats occupied:", countOccupied(seats), "\n\n")
    //Total iterations until stabilized: 134 --- seats occupied: 2211
    
    seats = part_2_seats
    iterations = 0
    for {
        seats, done = applyIterationPart2(seats)
        iterations++
        if done {
            break
        }
    }
	
    fmt.Println("Total iterations part 2 until stabilized:", iterations, "--- seats occupied:", countOccupied(seats))
    //Total iterations part 2 until stabilized: 85 --- seats occupied: 1995
}
