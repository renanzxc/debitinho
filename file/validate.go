package file

import (
	"log"

	"debitinho/color"
	"debitinho/utils"
)

// getAllowedLineTypes retorna os tipos de linhas que podem ser utilizadas por cada tipo de arquivo
func getAllowedLineTypes(fileType FileType) (allowedLineTypes []LineType) {
	switch fileType {
	case FileShippingType:
		allowedLineTypes = []LineType{LineTypeA, LineTypeC, LineTypeD, LineTypeE, LineTypeI, LineTypeJ, LineTypeK, LineTypeL, LineTypeZ}
	case FileReturnType:
		allowedLineTypes = []LineType{LineTypeA, LineTypeB, LineTypeF, LineTypeH, LineTypeT, LineTypeX, LineTypeZ}
	default:
		log.Fatal("Tipo de arquivo não identificado!")
	}

	return
}

func basicLineValidation(line string) {
	if len(line) != 150 {
		errorLine := color.AddColorSides(line, color.Red)
		log.Fatalf("A linha possui um tamanho inválido\n%s", errorLine)
	}
}

func basicFileValidation(file *File) {
	var (
		hasLines            = make(map[LineType]bool)
		errorTypeMissingMsg = "O arquivo não possui registro(s) do(s) tipo(s) "
		allowedLineTypes    = getAllowedLineTypes(file.Type)
	)

	if len(file.Lines) <= 0 {
		log.Fatal("O arquivo não possui linhas")
	}

	for _, line := range file.Lines {
		hasLines[line.Type()] = true
		if !utils.IsIn(line.Type(), allowedLineTypes...) {
			errorLine := color.AddColorSides(string(line.Type()), color.Red) + line.String()[1:]
			log.Fatalf("O arquivo de %s contêm registro do tipo %s\n%s", string(file.Type), string(line.Type()), errorLine)
		}
	}

	if !hasLines[LineTypeA] {
		errorTypeMissingMsg += "'A'"
	}
	if !hasLines[LineTypeZ] {
		errorTypeMissingMsg += "'Z'"
	}

	if !hasLines[LineTypeA] || !hasLines[LineTypeZ] {
		log.Fatal(errorTypeMissingMsg)
	}

}
