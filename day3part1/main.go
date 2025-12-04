package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type JoltagePair struct {
	value int
	index int
}

func main() {
	inputFile := os.Args[1]

	joltage := getTotalJoltage(inputFile)
	fmt.Printf("joltage: %v\n", joltage)
}

func getTotalJoltage(inputFile string) int {
	filehandle, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer filehandle.Close()

	sumJoltage := 0
	scanner := bufio.NewReader(filehandle)
	for {
		line, err := scanner.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if len(line) != 0 {
			currentLine := cleanUpLine(line)
			sumJoltage += findLineJoltage(currentLine)
		}
	}
	return sumJoltage
}

func findLineJoltage(line string) int {
	joltages := findTwoHighestJoltages(line)
	first, second := getSortedJoltages(joltages)

	return first*10 + second
}

func getSortedJoltages(joltages []*JoltagePair) (int, int) {
	if joltages[0].index > joltages[1].index {
		return joltages[1].value, joltages[0].value
	} else {
		return joltages[0].value, joltages[1].value
	}
}

func findTwoHighestJoltages(line string) []*JoltagePair {
	result := make([]*JoltagePair, 0)
	first := findHighestJoltage(line, nil)
	second := findHighestJoltage(line, first)
	return append(result, first, second)
}

func findHighestJoltage(line string, toExclude *JoltagePair) *JoltagePair {
	value, index := 0, 0
	lineCopy := line
	indexOffset := 0
	if toExclude != nil && toExclude.index == len(line) {
		lineCopy = lineCopy[:len(lineCopy)-1]
	} else if toExclude != nil {
		lineCopy = lineCopy[toExclude.index:]
		indexOffset = toExclude.index
	}
	for current := 0; current < len(lineCopy); current++ {
		number := getNumberAt(lineCopy, current)

		if number > value {
			value = number
			index = current + indexOffset + 1
		}
	}

	return &JoltagePair{
		value: value,
		index: index,
	}
}

func getNumberAt(s string, i int) int {
	currentChar := s[i : i+1]
	number, err := strconv.Atoi(currentChar)
	if err != nil {
		log.Fatal(err)
	}
	return number
}

func cleanUpLine(line string) string {
	value := strings.Replace(line, "\n", "", 1)
	value = strings.Replace(value, "\r", "", 1)

	return value
}
