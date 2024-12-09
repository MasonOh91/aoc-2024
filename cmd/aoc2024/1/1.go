package aoc2024_1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readInputFile() ([2][1000]int, error) {
	var locationsList1 []int
	var locationsList2 []int

	file, err := os.Open("./cmd/aoc2024/1/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read and print each line
	for scanner.Scan() {
		line := scanner.Text()
		locationsSplit := strings.Split(line, " ")
		location1, _ := strconv.Atoi(locationsSplit[0])
		location2, _ := strconv.Atoi(locationsSplit[len(locationsSplit)-1])
		locationsList1 = append(locationsList1, location1)
		locationsList2 = append(locationsList2, location2)
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	sort.Ints(locationsList1)
	sort.Ints(locationsList2)

	res := [2][1000]int{[1000]int(locationsList1), [1000]int(locationsList2)}
	return res, nil
}

func AOC2024_1() {
	arrays, _ := readInputFile()

	var sum int

	for index, val := range arrays[0] {
		var tempSum int
		tempSum = val - arrays[1][index]
		if tempSum < 0 {
			tempSum = tempSum * -1
		}
		sum += tempSum
	}

	fmt.Println("Total Distance:", sum)
}

func AOC2024_2() {
	arrays, _ := readInputFile()
}
