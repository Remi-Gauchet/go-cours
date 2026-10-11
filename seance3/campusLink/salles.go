package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Salle struct {
	ID       string
	Nom      string
	Batiment string
	Capacite int
}

func lireSalle(ligne string) (Salle, error) {
	champs := strings.Split(ligne, ",")
	if len(champs) != 4 {
		return Salle{}, fmt.Errorf("4 champs attendus, %d trouvés", len(champs))
	}
	for i := range champs {
		champs[i] = strings.TrimSpace(champs[i])
	}
	if champs[0] == "" {
		return Salle{}, fmt.Errorf("l'identifiant est vide")
	}
	capacite, err := strconv.Atoi(champs[3])
	if err != nil {
		return Salle{}, fmt.Errorf("la capacité %q n'est pas un nombre", champs[3])
	}
	if capacite <= 0 {
		return Salle{}, fmt.Errorf("la capacité doit être positive, %d trouvé", capacite)
	}
	return Salle{ID: champs[0], Nom: champs[1], Batiment: champs[2], Capacite: capacite}, nil
}

func lireSalles(texte string) ([]Salle, []string) {
	var salles []Salle
	var problemes []string
	lignes := strings.Split(texte, "\n")
	for i, ligne := range lignes {
		if i == 0 || strings.TrimSpace(ligne) == "" {
			continue
		}
		salle, err := lireSalle(ligne)
		if err != nil {
			problemes = append(problemes, fmt.Sprintf("ligne %d : %v", i+1, err))
			continue
		}
		salles = append(salles, salle)
	}
	return salles, problemes
}
