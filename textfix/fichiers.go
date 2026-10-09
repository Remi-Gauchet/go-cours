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

func ecrireFichier(chemin string, texte string) error {
	err := os.WriteFile(chemin, []byte(texte), 0644)
	if err != nil {
		return fmt.Errorf("Impossible d'écrire le fichier %s", chemin)
	}
	return nil
}
