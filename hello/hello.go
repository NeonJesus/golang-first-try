package main

import (
	"fmt"
	"log"
	"example/greetings"
)

func main() {

	log.SetPrefix("Greetings: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Darrin"}

	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}