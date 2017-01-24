package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	rachidF, err := template.ParseFiles("file.rac")
	if err != nil {
		log.Fatal(err)

	}

	err = rachidF.ExecuteTemplate(os.Stdout, "file.rac", nil)
	if err != nil {

		log.Fatal(err)
	}

	rachidF, err = rachidF.ParseFiles("file.oma", "file.red")
	if err != nil {
		log.Fatalln(err)
	}
	err = rachidF.ExecuteTemplate(os.Stdout, "file.red", nil)
	if err != nil {

		log.Fatal(err)
	}
	err = rachidF.ExecuteTemplate(os.Stdout, "file.oma", nil)
	if err != nil {

		log.Fatal(err)
	}
	err = rachidF.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
