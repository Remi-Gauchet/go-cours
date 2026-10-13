package main

import (
	"fmt"
	"os"
)

// main lance le programme textfix qui consiste à récupérer un string dans un fichier texte (.txt) à lui appliquer des modifications, puis à le réécrire dans un autre fichier texte
// fonctionne en définissant les noms des deux fichiers en argument
func main() {
	// Je commence par verifier les arguments
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Erreur : Veuillez entrer le fichier d'entrée puis le fichier de sortie."+"\nExemple : go run . entree.txt resultat.txt")
		os.Exit(1)
	}

	if os.Args[1] == os.Args[2] {
		fmt.Fprintln(os.Stderr, "Veillez à ce que le nom du fichier d'entrée et de sortie soient différents.")
		os.Exit(1)
	}

	//Je vais chercher mon contenu
	contenu, err := lireFichier(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Impossible de lire le fichier", err)
		os.Exit(1)
	}

	// Je modifie mon contenu
	contenu = transformer(contenu)

	// Enregistre le contenu modifié
	err = ecrireFichier(os.Args[2], contenu)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("Terminé : ", os.Args[2], " a été créé.")
}
