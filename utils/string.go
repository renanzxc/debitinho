package utils

import (
	"bufio"
	"log"
	"os"
)

func IsInStr(str string, strs ...string) bool {
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
