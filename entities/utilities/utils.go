package utilities

import (
	"fmt"
	"os"
	"strings"
)

func Exists(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		return fmt.Errorf("Path:%s does not exists", path)
	}
	return fmt.Errorf("Path:%s cannot be loaded", path)
}

func IsValueNonEmpty(value string) bool {
	if strings.TrimSpace(value) == "" {
		return false
	} else {
		return true
	}
}
