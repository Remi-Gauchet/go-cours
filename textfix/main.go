package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Veuillez entrer deux arguments : le fichier d'entrée puis le fichier de sortie.")
		os.Exit(1)
	}
	fmt.Println(lireFichier(os.Args[1]))
}
