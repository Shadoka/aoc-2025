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

type SafeDial struct {
	Position        int
	RotationCounter int
	ZeroCounter     int
}

func (sd *SafeDial) Right(amount int) {
	rotations := (amount + sd.Position) / 100
	// our zero counter will get incremented
	if (amount+sd.Position)%100 == 0 {
		rotations--
	}
	sd.RotationCounter += rotations
	sd.CalculatePosition(amount)
}

func (sd *SafeDial) Left(amount int) {
	rotations := 0
	fullRotations := amount / 100
	modulAmount := amount % 100

	rotations += fullRotations

	if sd.Position != 0 && modulAmount > sd.Position {
		rotations++
	}
	sd.RotationCounter += rotations

	invertedAmount := 100 - modulAmount
	amount = amount - modulAmount + invertedAmount
	sd.CalculatePosition(amount)
}

func (sd *SafeDial) CalculatePosition(amount int) {
	sd.Position = (sd.Position + amount) % 100
	if sd.Position == 0 {
		sd.ZeroCounter++
	}
}

func main() {
	inputFile := os.Args[1]

	dial := SafeDial{
		Position:        50,
		RotationCounter: 0,
		ZeroCounter:     0,
	}
	countZeroesWithDial(inputFile, &dial)
	fmt.Printf("ending zeroes: %v, rotating zeroes: %v\n", dial.ZeroCounter, dial.RotationCounter)
	fmt.Printf("total zeroes: %v\n", dial.RotationCounter+dial.ZeroCounter)
	fmt.Printf("position: %v\n", dial.Position)
}

func countZeroesWithDial(inputFile string, dial *SafeDial) {
	filehandle, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer filehandle.Close()

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
			completeValue := parseLine(line)
			absValue := abs(completeValue)
			if completeValue > 0 {
				dial.Right(absValue)
			} else {
				dial.Left(absValue)
			}
		}
	}
}

func abs(value int) int {
	return max(value, -value)
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

	if direction == "R" {
		return intValue
	} else {
		return -intValue
	}
}
