package main

import "fmt"

type person struct {
	id   int
	name string
}

type agent struct {
	person
	areofWork string
	ritered   bool
}

func main() {
	p := person{1, "Pippo"}
	sa := agent{person: person{2, "Pluto"}}
	fmt.Printf("The person: %s has got id=%d \n", p.name, p.id)
	p.pSpeak()
	sa.saSpeak()
}

func (p person) pSpeak() {
	fmt.Println("I am a person")

}

func (p agent) saSpeak() {
	fmt.Println("I'm a a secret agent")

}
