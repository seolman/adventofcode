package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var leftList, rightList []int

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "   ")
		left, err1 := strconv.Atoi(parts[0])
		right, err2 := strconv.Atoi(parts[1])
		if err1 == nil && err2 == nil {
			leftList = append(leftList, left)
			rightList = append(rightList, right)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	totalDistance := 0
	for i := 0; i < len(leftList); i++ {
		distance := abs(leftList[i] - rightList[i])
		totalDistance += distance
	}

	fmt.Println(totalDistance)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
