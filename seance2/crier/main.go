package main

import (
	"fmt"
	"os"
	"strings"
)

func crier(texte string) string {
	cricri := string(strings.ToUpper(texte) + " !")
	return cricri
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Erreur : il faut au moins un mot.")
		fmt.Println("Exemple : go run ./crier attention")
		os.Exit(1)
	} else {
		fmt.Println(crier(strings.Join(os.Args[1:], " ")))
	}
}
