package file

import (
	"debitinho/utils"
	"log"
	"path/filepath"
)

func ParseFile(path string) (parsedFile *File) {
	parsedFile = &File{Name: filepath.Base(path)}

	linesStr, err := utils.GetFileLines(path)
	if err != nil {
		log.Fatal(err)
	}

	for ii := range linesStr {
		line := ParseLine(linesStr[ii])
		parsedFile.Lines = append(parsedFile.Lines, line)
	}

	parsedFile.Type = getFileType(parsedFile)
	basicFileValidation(parsedFile)

	return
}
