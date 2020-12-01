package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	// Reading the content from the input file
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	text := string(content)
	inputs := strings.Split(text, "\n")

	partOneFound, partTwoFound := false, false

	// Iterating over the inputs to find two of them that sum is 2020
	for _, input1Str := range inputs {
		for _, input2Str := range inputs {
			input1, _ := strconv.Atoi(input1Str)
			input2, _ := strconv.Atoi(input2Str)

			if input1+input2 == 2020 && !partOneFound {
				fmt.Printf("Part 1 answer is %d\n", input1*input2)
				partOneFound = true
			}

			if !partTwoFound {
				// A third loop to find the answer to part 2 in the same script
				for _, input3Str := range inputs {
					input3, _ := strconv.Atoi(input3Str)
					if input1+input2+input3 == 2020 {
						fmt.Printf("Part 2 answer is %d\n", input1*input2*input3)
						partTwoFound = true
					}
				}
			}

			// All the parts are found - We can get out of the script!
			if partOneFound && partTwoFound {
				return
			}
		}
	}
}
