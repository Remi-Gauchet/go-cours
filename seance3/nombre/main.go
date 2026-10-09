package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Erreur : veuillez renseigner un nombre entier.")
		os.Exit(1)
	}
	texte := os.Args[1]
	nombre, err := strconv.Atoi(texte)
	if err != nil {
		fmt.Println("Erreur :", texte, " n'est pas un nombre entier.")
		os.Exit(1)
	}
	fmt.Println("Le double est", nombre*2)
}
