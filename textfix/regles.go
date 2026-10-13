package main

import (
	"strings"
	"unicode"
)

// transformer prend un string et le remanie en lui appliquant des fonctions de mise en forme ou de mise en page
func transformer(texte string) string {
	lignes := strings.Split(texte, "\n")

	for i, ligne := range lignes {
		ligne = strings.TrimRight(ligne, "\r")
		mots := strings.Fields(ligne)
		mots = appliquerMarqueurs(mots)
		lignes[i] = strings.Join(mots, " ")
	}

	return strings.Join(lignes, "\n")
}

// capitaliser met une majuscule au début du mot et le reste en mjnuscule
func capitaliser(mot string) string {
	if mot == "" {
		return mot
	}
	lettres := []rune(strings.ToLower(mot))
	lettres[0] = unicode.ToUpper(lettres[0])
	return string(lettres)
}

// appliquerMarqueurs détecte dans le string certains "marqueurs" prédéfinis et modifie le string en fonction d'eux
func appliquerMarqueurs(mots []string) []string {
	var resultat []string
	for _, mot := range mots {
		if mot == "(up)" || mot == "(low)" || mot == "(cap)" {
			if len(resultat) == 0 {
				continue
			}
			dernier := len(resultat) - 1
			// puis on choisit la transformation selon le marqueur
			switch mot {
			case "(up)":
				resultat[dernier] = strings.ToUpper(resultat[dernier])
			case "(low)":
				resultat[dernier] = strings.ToLower(resultat[dernier])
			case "(cap)":
				resultat[dernier] = capitaliser(resultat[dernier])
			}
		} else {
			resultat = append(resultat, mot)
		}
	}
	return resultat
}
