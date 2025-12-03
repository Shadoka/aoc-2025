package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var factorTable = map[int][]int{
	1:  {},
	2:  {1},
	3:  {1},
	4:  {1, 2},
	5:  {1},
	6:  {1, 2, 3},
	7:  {1},
	8:  {1, 2, 4},
	9:  {1, 3},
	10: {1, 2, 5},
}

func main() {
	inputFile := os.Args[1]
	// inputFile := "example.txt"
	sum := countInvalidIds(inputFile)
	fmt.Printf("sum of invalid ids: %v", sum)
}

func countInvalidIds(inputFile string) uint64 {
	filehandle, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer filehandle.Close()

	scanner := bufio.NewReader(filehandle)
	line, err := scanner.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	line = line[:len(line)-2]
	var sum uint64 = 0
	if len(line) != 0 {
		idRanges := strings.Split(line, ",")
		for _, idRange := range idRanges {
			sum += processIdRange(idRange)
		}
	}
	return sum
}

func processIdRange(idRange string) uint64 {
	rangeValues := strings.Split(idRange, "-")

	from, err := strconv.ParseUint(rangeValues[0], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	to, err := strconv.ParseUint(rangeValues[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	mirroredSequences := findMirroredValues(from, to)
	var sum uint64 = 0
	for _, seq := range mirroredSequences {
		sum += seq
	}
	return sum
}

func findMirroredValues(from uint64, to uint64) []uint64 {
	sequences := make([]uint64, 0)
	// we do this because go doesnt have a 'range from...to' syntax
	rangeEnd := to - from + 1
	for current := range rangeEnd {
		value := current + from
		if isMirrored(value) {
			sequences = append(sequences, value)
		}
	}
	return sequences
}

func isMirrored(value uint64) bool {
	valueString := strconv.FormatUint(value, 10)
	valueLength := len(valueString)

	factors := factorTable[valueLength]
	for _, factor := range factors {
		if splitAndCheckEquality(valueString, factor) {
			return true
		}
	}
	return false
}

func splitAndCheckEquality(value string, factor int) bool {
	// split
	parts := make([]string, 0)
	partsAmount := len(value) / factor
	previousStart := 0
	for i := range partsAmount {
		end := factor * (i + 1)
		parts = append(parts, value[previousStart:end])
		previousStart = end
	}

	// check
	isEquals := true
	for i := partsAmount - 1; i > 0; i-- {
		isEquals = isEquals && parts[i] == parts[i-1]
	}
	return isEquals
}
