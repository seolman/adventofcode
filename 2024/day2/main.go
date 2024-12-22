package main

import (
	"bufio"
	"fmt"
	"os"
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
	var reports [][]int

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		var report []int
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err == nil {
				report = append(report, num)
			}
		}
		reports = append(reports, report)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	safeCount := 0
	for _, report := range reports {
		if isSafeReport(report) {
			safeCount++
		}
	}

	fmt.Println("Number of safe reports:", safeCount)

}

func isSafeReport(report []int) bool {
	if len(report) < 2 {
		return false
	}

	increasing := true
	decreasing := true
	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]
		if diff == 0 {
			return false
		}
		if diff > 0 {
			decreasing = false
			if diff > 3 || diff < 1 {
				return false
			}
		} else {
			increasing = false
			if diff < -3 || diff > -1 {
				return false
			}
		}
	}

	return increasing || decreasing
}

func Part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var reports [][]int

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		var report []int
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err == nil {
				report = append(report, num)
			}
		}
		reports = append(reports, report)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	safeCount := 0
	for _, report := range reports {
		if isSafeReport(report) || canBeSafeByRemovingOneLevel(report) {
			safeCount++
		}
	}

	fmt.Println("Number of safe reports:", safeCount)
}

func canBeSafeByRemovingOneLevel(report []int) bool {
	for i := 0; i < len(report); i++ {
		newReport := append([]int{}, report...)
		newReport = append(newReport[:i], newReport[i+1:]...)
		if isSafeReport(newReport) {
			return true
		}
	}
	return false
}

func main() {
  Part1()
  Part2()
}
