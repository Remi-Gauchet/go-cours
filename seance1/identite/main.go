package main

import "fmt"

func main() {
	prenom := "Rémi"
	ville := "Viarmes"
	age := 31
	fmt.Println("Prénom : " + prenom)
	fmt.Println("Ville : " + ville)
	fmt.Println("Âge : ", age)
	age = age + 1
	fmt.Println("L'année prochaine, j'aurai : ", age)
	fmt.Println("Léa habite à Paris.")
}
