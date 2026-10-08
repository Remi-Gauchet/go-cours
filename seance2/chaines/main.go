package main

import (
	"fmt"
	"strings"
)

func main() {
	var mot string = "Bonjour"
	fmt.Println(strings.Contains(mot, "jour"))

	var space string = " trop d'espace ici "
	fmt.Println(strings.TrimSpace(space))

	var animal string = "le chat et le chien"
	animal = strings.ReplaceAll(animal, "chat", "putois")
	animal = strings.ReplaceAll(animal, "chien", "chat")
	animal = strings.ReplaceAll(animal, "putois", "chien")
	fmt.Println(animal)

	var phrase string
	if strings.Contains(phrase, "urgent") {
		fmt.Println("Message urgent !")
	}
}
