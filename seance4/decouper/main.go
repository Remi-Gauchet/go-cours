package main

import (
	"fmt"
	"strings"
)

func main() {
	texte := "le chat dort\nle chien   joue"
	lignes := strings.Split(texte, "\n")
	for i, ligne := range lignes {
		mots := strings.Fields(ligne)
		fmt.Println("ligne", i+1, ":", len(mots), "mots")
		for _, mot := range mots {
			fmt.Println("-", mot)
		}
	}
}
