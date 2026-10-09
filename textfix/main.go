package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Erreur : Veuillez entrer le fichier d'entrée puis le fichier de sortie."+"\nExemple : go run . entree.txt resultat.txt")
		os.Exit(1)
	}

	if os.Args[1] == os.Args[2] {
		fmt.Fprintln(os.Stderr, "Veillez à ce que le nom du fichier d'entrée et de sortie soient différents.")
		os.Exit(1)
	}

	contenu, err := lireFichier(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Impossible de lire le fichier")
		os.Exit(1)
	}

	err = ecrireFichier(os.Args[2], contenu)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Impossible d'écrire le fichier")
		os.Exit(1)
	}
	fmt.Println("Terminé : ", os.Args[2], " a été créé.")
}
