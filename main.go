package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	myInput := os.Getenv("INPUT_MYINPUT")

	output := fmt.Sprintf("Hello %s", myInput)

	f, err := os.OpenFile(os.Getenv("GITHUB_OUTPUT"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(fmt.Sprintf(`myOutput=%s`, output)); err != nil {
		log.Println(err)
	}

}