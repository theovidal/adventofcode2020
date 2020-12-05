package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	// Reading the content from the input file
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	text := string(file)
	seats := strings.Split(text, "\n")

	// Checking all the seats
	highestID := 0
	ids := []int{}
	for _, seat := range seats {
		if seat == "" {
			continue
		}
		row := findNumber('F', 'B', seat[0:7], 128)
		column := findNumber('L', 'R', seat[7:], 8)
		id := row*8 + column

		ids = append(ids, id)
		if id > highestID {
			highestID = id
		}
	}

	fmt.Printf("Part 1 - The highest seat ID is %d\n", highestID)

	sort.Ints(ids)
	for index, id := range ids {
		if ids[index+1] != id+1 {
			fmt.Printf("Part 2 - Your seat ID is %d\n", id+1)
			return
		}
	}
}

// findNumber finds the row or column number of the seat
func findNumber(lowerHalf, upperHalf rune, sequence string, number int) int {
	rows := []int{}
	for i := 0; i < number; i++ {
		rows = append(rows, i)
	}
	for _, half := range sequence {
		limit := (len(rows) / 2)
		if half == lowerHalf {
			rows = rows[:limit]
		} else {
			rows = rows[limit:]
		}
	}
	return rows[0]
}
