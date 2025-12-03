package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

	// mirrored sequences have to be of even length
	if len(valueString)%2 != 0 {
		return false
	}

	valueLength := len(valueString) / 2
	firstHalf := valueString[:valueLength]
	secondHalf := valueString[valueLength:]
	return firstHalf == secondHalf
}
