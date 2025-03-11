package shared

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func WriteRTTValue(fileName string, elapsedTime float64) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(fmt.Sprintf("%f\n", elapsedTime))

	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	writer.Flush()
}

func ReadRTTValues(filename string) ([]float64, error) {
	var values []float64

	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		num, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		values = append(values, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return values, nil
}
