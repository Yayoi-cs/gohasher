package ioHelper

import (
	"bufio"
	"fmt"
	"os"
)

func ReadHashFile(filename string) (string, error) {
	fmt.Println("Reading File..")
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error: ")
		return "", err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	lineCount := 0
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		lineCount++
		if lineCount > 1 {
			return "", fmt.Errorf("Too much lines\n")
		}
	}
	if err = scanner.Err(); err != nil {
		return "", err
	}
	return line, nil
}
