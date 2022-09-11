package utils

import (
	"bufio"
	"log"
	"os"
)

func IsIn[T comparable](str T, strs ...T) bool {
	for ii := range strs {
		if str == strs[ii] {
			return true
		}
	}

	return false
}

func IsInByte(str byte, strs ...byte) bool {
	for ii := range strs {
		if str == strs[ii] {
			return true
		}
	}

	return false
}

func GetFileLines(filePath string) (lines []string, err error) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}
