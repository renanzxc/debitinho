package file

import (
	"debitinho/utils"
	"log"
)

type FileType string

const (
	ShippingType FileType = "remessa"
	ReturnType   FileType = "retorno"
)

type File struct {
	Name  string
	Type  FileType
	Lines []ILine
}

func (f *File) identifyFileType(lines []string) {
	for _, line := range lines {
		if len(line) > 2 && line[0] == 'A' {
			switch line[1] {
			case '1':
				f.Type = ShippingType
			case '2':
				f.Type = ReturnType
			}
			break
		}
	}

	if f.Type == "" {
		userInput := utils.GetUserInput("Qual o tipo de arquivo? 1 - Remessa | 2 - Retorno")
		switch userInput {
		case "1":
			f.Type = ShippingType
		case "2":
			f.Type = ReturnType
		default:
			log.Fatal("tipo de arquivo inválido")
		}
	}
}
