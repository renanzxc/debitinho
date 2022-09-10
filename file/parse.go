package file

import (
	"debitinho/utils"
	"log"
)

func Parse(filePath string) (parsedFile *File) {
	lines, err := utils.GetFileLines(filePath)
	if err != nil {
		log.Fatal(err)
	}

	parsedFile = &File{Name: filePath}
	parsedFile.identifyFileType(lines)
	for ii := range lines {
		line := NewLine(lines[ii])
		parsedFile.Lines = append(parsedFile.Lines, line)
	}

	return
}
