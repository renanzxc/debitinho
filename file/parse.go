package file

import (
	"debitinho/utils"
	"errors"
	"log"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
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

func parseLine(originline string, line ILine) (err error) {
	reflectValue := reflect.Indirect(reflect.ValueOf(line))
	reflectType := reflect.TypeOf(line).Elem()
	if reflectType.Kind() != reflect.Struct {
		return errors.New("tipo inválido")
	}
	for ii := 0; ii < reflectType.NumField(); ii++ {
		var posicaoInicial, posicaoFinal int64

		if tagPosicao := reflectType.Field(ii).Tag.Get("posicao"); tagPosicao != "" {
			reflectValueField := reflectValue.Field(ii)
			tagPosicaoArr := strings.Split(tagPosicao, ":")

			switch len(tagPosicaoArr) {
			case 1:
				posicaoInicial, err = strconv.ParseInt(tagPosicaoArr[0], 10, 64)
				if err != nil {
					return errors.New("valor inicial inválido")
				}

				reflectValueField.SetString(string(originline[posicaoInicial]))
			case 2:
				posicaoFinal, err = strconv.ParseInt(tagPosicaoArr[1], 10, 64)
				if err != nil {
					return errors.New("valor final inválido")
				}

				posicaoInicial, err = strconv.ParseInt(tagPosicaoArr[0], 10, 64)
				if err != nil {
					return errors.New("valor inicial inválido")
				}

				reflectValueField.SetString(string(originline[posicaoInicial:posicaoFinal]))
			default:
				return errors.New("valor inválido da tag")
			}
		}

	}

	return
}
