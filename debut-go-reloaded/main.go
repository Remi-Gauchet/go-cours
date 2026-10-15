package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Remi-Gauchet/Go-reloaded.git/archives/tolower"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Entrez à la suite le nom de votre texte, et le nom du texte corrigé")
		fmt.Fprintln(os.Stderr, "Exemple : go run . sample.txt result.txt")
		os.Exit(1)
	}
	lecturebyte, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("Erreur lors de lecture du fichier : %v", err)
	}

	contenustring := string(lecturebyte)

	// ici je vais soumettre la variable "contenustring" à mes fonctions pour procéder aux changements
	tolower.ToLower(contenustring)

	err = os.WriteFile(os.Args[2], []byte(contenustring), 0644)
	if err != nil {
		log.Fatalf("Erreur lors de l'écriture du fichier : %v", err)
	}
	fmt.Fprintln(os.Stderr, "Fichier corrigé avec succès.")
	os.Exit(0)
}
