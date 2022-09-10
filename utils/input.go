package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func GetUserInput(msg string) (text string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(msg)
	fmt.Print("-> ")
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	text = strings.Replace(text, "\n", "", -1)

	return
}
