package main

import (
	"strings"
	"unicode"
)

func transformer(texte string) string {
	lignes := strings.Split(texte, "\n")

	for i, ligne := range lignes {
		ligne = strings.TrimRight(ligne, "\r")
		mots := strings.Fields(ligne)
		lignes[i] = strings.Join(mots, " ")
	}

	return strings.Join(lignes, "\n")
}

func capitaliser(mot string) string {
	if mot == "" {
		return mot
	}
	lettres := []rune(strings.ToLower(mot))
	lettres[0] = unicode.ToUpper(lettres[0])
	return string(lettres)
}

func appliquerMarqueurs(mots []string) []string {
	var resultat []string
	for _, mot := range mots {
		dernier := len(resultat) - 1
		switch mot {
		case "(up)":
			resultat[dernier] = strings.ToUpper(resultat[dernier])
		case "(low)":
			resultat[dernier] = strings.ToLower(resultat[dernier])
		case "(cap)":
			resultat[dernier] = capitaliser(resultat[dernier])
		default:
			resultat = append(resultat, mot)
		}
	}
	return resultat
}
