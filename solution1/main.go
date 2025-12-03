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

func main() {
	inputFile := os.Args[1]

	filehandle, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer filehandle.Close()

	position := 50
	amountZeroes := 0
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
			currentStep := parseLine(line)
			position = (position + currentStep) % 100
			if position == 0 {
				amountZeroes++
			}
		}
	}

	fmt.Println(amountZeroes)
}

func parseLine(line string) int {
	direction := line[:1]
	value := line[1:]
	value = strings.Replace(value, "\n", "", 1)
	value = strings.Replace(value, "\r", "", 1)

	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal(err)
	}
	intValue = (intValue % 100)

	if direction == "R" {
		return intValue
	} else {
		return -intValue
	}
}
