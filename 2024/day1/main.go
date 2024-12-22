package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Part1() {
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
		distance := int(math.Abs(float64(leftList[i] - rightList[i])))
		totalDistance += distance
	}

	fmt.Println(totalDistance)
}

func Part2() {
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

	rightCount := make(map[int]int)
	for _, num := range rightList {
		rightCount[num]++
	}

	similarityScore := 0
	for _, num := range leftList {
		if count, exists := rightCount[num]; exists {
			similarityScore += num * count
		}
	}

	fmt.Println(similarityScore)
}

func main() {
  Part1()
  Part2()
}
