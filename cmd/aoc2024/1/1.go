package aoc2024_1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func AOC2024_1() {
	var locationsList1 []int
	var locationsList2 []int

	file, err := os.Open("./cmd/aoc2024/1/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
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

	fmt.Println((locationsList1))
	fmt.Println((locationsList2))

	sort.Ints(locationsList1)
	sort.Ints(locationsList2)

	var sum int

	for index, _ := range locationsList1 {
		fmt.Println((index))
		var tempSum int
		tempSum = locationsList1[index] - locationsList2[index]
		if tempSum < 0 {
			tempSum = tempSum * -1
		}
		sum += tempSum
	}

	fmt.Println(sum)
}
