package main

import (
	"fmt"
	"strings"
)

func main() {
	var enMajuscules []string
	fruits := []string{"concombre", "banane", "aubergine"}
	fruits = append(fruits, "courge")
	for i := range fruits {
		enMajuscules = append(enMajuscules, strings.ToUpper(fruits[i]))
	}
	for i, majfruit := range enMajuscules {
		fmt.Println(i+1, "-", majfruit)
	}

	var phrase string = "le chat dort"
	mots := strings.Fields(phrase) // ["le", "chat", "dort"]
	fmt.Println(len(mots))
	fmt.Println(strings.Join(mots, " "))
}
