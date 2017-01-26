package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

func main() {
	name := []string{"Alex", "Rachid", "Omar", "khalid", "manuel"}
	tpl, err := template.ParseFiles("letter.xmas")
	if err != nil {
		log.Println(err)

	}

	err = tpl.Execute(os.Stdout, name)
	if err != nil {

		fmt.Println("Problem with execution of template", err)
	}
}
