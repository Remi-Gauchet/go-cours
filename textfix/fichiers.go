package main

import (
	"fmt"
	"os"
)

func lireFichier(chemin string) (string, error) {
	contenu, err := os.ReadFile(chemin)
	if err != nil {
		return "", fmt.Errorf("Impossible de lire le fichier %s", chemin)
	}
	return string(contenu), nil
}
