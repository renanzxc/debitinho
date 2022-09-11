package file

import "log"

type FileType string

const (
	FileShippingType FileType = "remessa"
	FileReturnType   FileType = "retorno"
)

type File struct {
	Name  string
	Type  FileType
	Lines []ILine
}

func getFileType(file *File) (fType FileType) {
	for _, line := range file.Lines {
		if lineA, ok := line.(*LineA); ok {
			switch lineA.ShippingCode {
			case "1":
				fType = FileShippingType
			case "2":
				fType = FileReturnType
			}
			break
		}
	}

	if fType == "" {
		log.Fatal("Tipo de arquivo não identificado")
	}

	return
}
