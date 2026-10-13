package main

import (
	"fmt"
	"os"
)

// LireFichier ouvre un fichier.txt et le convertit en string
func lireFichier(chemin string) (string, error) {
	contenu, err := os.ReadFile(chemin)
	if err != nil {
		return "", fmt.Errorf("Impossible de lire le fichier %s : %w", chemin, err)
	}
	return string(contenu), nil
}

// ecrireFichier convertir un string et le range dans un fichier .txt
func ecrireFichier(chemin string, texte string) error {
	err := os.WriteFile(chemin, []byte(texte), 0644)
	if err != nil {
		return fmt.Errorf("Impossible d'écrire le fichier %s : %w", chemin, err)
	}
	return nil
}
