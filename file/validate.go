package file

import (
	"fmt"
	"log"

	"debitinho/color"
	"debitinho/utils"
)

type validation func(string) *utils.ErrorValidation

func validRegistersShipping() []string {
	return []string{"A", "C", "D", "E", "I", "J", "K", "L", "Z"}
}

func defaultValidations() map[string]validation {
	return map[string]validation{
		"Tipo de registro inválido": func(line string) *utils.ErrorValidation {
			if !utils.IsInStr(string(line[0]), validRegistersShipping()...) {
				lineColor := color.AddColorSides(string(line[0]), color.Red) + line[1:]
				return &utils.ErrorValidation{
					Err:  fmt.Errorf(`o arquivo de remessa contêm registro do tipo "%s"`, string(line[0])),
					Line: &lineColor,
				}
			}

			return nil
		},
		"Tamanho de resgitro inválido": func(line string) *utils.ErrorValidation {
			if len(line) != 150 {
				lineColor := color.AddColorSides(line, color.Red)
				return &utils.ErrorValidation{
					Err:  fmt.Errorf(""),
					Line: &lineColor,
				}
			}

			return nil
		},
	}
}

func defaultLineValidation(line string) {
	for validationName, validt := range defaultValidations() {
		errValid := validt(line)
		if errValid != nil && errValid.Err != nil {
			if errValid.Line != nil {
				log.Fatalf("%s: %s \n%s", validationName, errValid.Err.Error(), *errValid.Line)
			} else {
				log.Fatalf("%s: %s", validationName, errValid.Err.Error())
			}
		}
	}
}
