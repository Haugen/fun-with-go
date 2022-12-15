package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"White", "Fluffy", "Clouds"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	for name, message := range messages {
		name, err := greetings.ToUpper(name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(name, message)
	}

	fmt.Println(messages)
}
