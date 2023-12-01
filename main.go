package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("calibration_input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()
	var total int64
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		row := scanner.Text()
		numerals := extractStringNumbers(row)
		val, _ := strconv.Atoi(numerals)
		total += int64(val)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	fmt.Println(total)
}

func extractStringNumbers(input string) string {
	reg := regexp.MustCompile("[0-9]+")
	matches := reg.FindAllString(input, -1)
	num := ""
	for _, match := range matches {
		num += match
	}

	return num
}
