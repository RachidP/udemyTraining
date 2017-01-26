package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	var name string
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	name = "Prof. Go lang "
	err = tpl.Execute(os.Stdout, name)
	if err != nil {
		log.Fatalln(err)
	}
}

// Do not use the above code in production
// We will learn about efficiency improvements soon!
