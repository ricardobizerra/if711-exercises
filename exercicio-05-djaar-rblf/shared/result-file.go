package shared

import (
	"bufio"
	"fmt"
	"os"
)

func WriteRTTValue(fileName string, elapsedTime float64) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

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
