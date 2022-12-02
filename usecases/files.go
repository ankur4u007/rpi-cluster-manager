package usecases

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AppendTextIfNotPresent(text string, filePath string) error {
	err := appendToFile(text, filePath)
	if err != nil {
		fmt.Printf("Filed to append file:%s\n", filePath)
		return err
	}
	return nil
}

func appendToFile(text string, filePath string) error {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0600)
	if err != nil {
		return err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	line, _, err := reader.ReadLine()
	if strings.Contains(string(line), text) {
		fmt.Printf("Skipped writing:%s to file:%s as it already exists\n", text, filePath)

	} else {
		appendedText := fmt.Sprintf("%s %s", line, text)
		if _, err = file.WriteAt([]byte(appendedText), 0); err != nil {
			return err
		} else {
			fmt.Printf("Successfully written:%s to file:%s\n", text, filePath)
		}
	}

	return nil
}
