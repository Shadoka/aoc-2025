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

func (sd *SafeDial) Increment(amount int) {
	fmt.Println("test")
	absPosition := max(sd.Position, -sd.Position)
	sd.RotationCounter += (amount + absPosition) / 100
	sd.Position = (sd.Position + amount) % 100
	if sd.Position == 0 {
		sd.ZeroCounter++
	}
}

func (sd *SafeDial) Decrement(amount int) {
	negPosition := min(sd.Position, -sd.Position)
	rotations := (negPosition - amount) / 100
	sd.RotationCounter += rotations
	newPosition := (sd.Position - amount) % 100
	if rotations > 0 {
		sd.Position = (100 - newPosition) * -1
	} else {
		sd.Position = newPosition
	}
	if sd.Position == 0 {
		sd.ZeroCounter++
	}
}

func main() {
	inputFile := os.Args[1]
	// inputFile := "example.txt"

	dial := SafeDial{
		Position:        50,
		RotationCounter: 0,
		ZeroCounter:     0,
	}
	countZeroesWithDial(inputFile, &dial)
	fmt.Printf("total zeroes: %v\n", dial.RotationCounter)
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
			_, completeValue := parseLine(line)
			if completeValue > 0 {
				dial.Increment(completeValue)
			} else {
				dial.Decrement(completeValue)
			}
		}
	}

	fmt.Printf("ending zeroes: %v, rotating zeroes: %v\n", dial.ZeroCounter, dial.RotationCounter)
}

func countZeroes(inputFile string) int {
	filehandle, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer filehandle.Close()

	position := 50
	rotatingZeroes := 0
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
			currentStep, completeValue := parseLine(line)
			rotatingZeroes += countRotatingZeroes(completeValue, position)
			position = (position + currentStep) % 100

			if position == 0 {
				amountZeroes++
			}
		}
	}

	fmt.Printf("ending zeroes: %v, rotating zeroes: %v\n", amountZeroes, rotatingZeroes)
	return amountZeroes + rotatingZeroes
}

func countRotatingZeroes(completeValue int, position int) int {
	result := 0

	fullRotations := completeValue / 100
	result += max(fullRotations, -fullRotations)

	modulValue := completeValue % 100
	newPosition := position + modulValue
	if (position > 0 && newPosition < 0) || (position < 0 && newPosition > 0) {
		result += 1
	}

	absNewPosition := max(newPosition/100, -(newPosition / 100))
	result += absNewPosition

	return result
}

func parseLine(line string) (int, int) {
	direction := line[:1]
	value := line[1:]
	value = strings.Replace(value, "\n", "", 1)

	original, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal(err)
	}
	intValue := (original % 100)

	if direction == "R" {
		return intValue, original
	} else {
		return -intValue, -original
	}
}
