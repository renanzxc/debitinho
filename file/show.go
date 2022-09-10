package file

import (
	"encoding/json"
	"fmt"
	"log"
)

func ShowFile(file *File) {
	for ii := range file.Lines {
		lineByte, err := json.MarshalIndent(file.Lines[ii], "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(lineByte))
	}

}
