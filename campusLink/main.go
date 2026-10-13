package main

import (
	"fmt"
	"os"
)

func main() {
	var usage string = "Pour lancer le programme, vous devez entrer en argument en premier import-salles suivi d'un espace et le nom de fichier que vous désirez en format .csv\n Exemple : go run . import-salles <fichier.csv>"
	var err error
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}
	commande := os.Args[1]

	if commande == "aide" {
		fmt.Println(usage)
	} else if commande == "import-salles" {
		err = commandeImportSalles(os.Args[2:])
	} else {
		err = fmt.Errorf("commande inconnue : %s", commande)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, "Erreur :", err)
		os.Exit(1)
	}
}

func commandeImportSalles(fichiers []string) error {
	if len(fichiers) == 0 {
		return fmt.Errorf("aucun fichier indiqué")
	}
	contenu, err := lireFichier(fichiers[0])
	if err != nil {
		return err
	}
	salles, problemes := lireSalles(contenu)
	for _, salle := range salles {
		fmt.Printf("%s - %s - bâtiment %s - %d places\n", salle.ID, salle.Nom, salle.Batiment, salle.Capacite)
	}
	for _, problem := range problemes {
		fmt.Fprintln(os.Stderr, "Ligne ignorée :", problem)
	}
	fmt.Printf("%d salle(s) importée(s), %d ligne(s) ignorée(s).", len(salles), len(problemes))
	return nil
}
