package usecases

import (
	"fmt"
	"os"
)

func AppendFile(text string, filePath string) error {
	err := append(text, filePath)
	if err != nil {
		fmt.Printf("Filed to append file:%s\n", filePath)
		return err
	}
	return nil
}

func append(text string, filePath string) error {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err = f.WriteString(text); err != nil {
		return err
	}
	return nil
}
