package datafile

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func GetFloats(fileName string) ([3]float64, error) {
	var numbers [3]float64
	file, err := os.Open(fileName)
	if err != nil {
		return numbers, err
	}
	defer file.Close()
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if i >= 3 {
			break
		}
		line := scanner.Text()
		if line == "" {
			continue
		}
		numbers[i], err = strconv.ParseFloat(line, 64)
		if err != nil {
			return numbers, fmt.Errorf("line %d: %v", i+1, err)
		}
		i++
	}

	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	if i < 3 {
		return numbers, fmt.Errorf("expected 3 numbers, got %d", i)
	}
	return numbers, nil
}
