package main

import (
	"fmt"
)

func main() {
	var age int
	if age >= 18 {
		fmt.Println("Tu es majeur.")
	} else {
		fmt.Println("Tu es mineur.")
	}

	var note int
	if note >= 16 {
		fmt.Println("Très bien")
	}
	if note >= 10 {
		fmt.Println("Admis")
	} else {
		fmt.Println("Non admis")
	}

	var motDePasse string
	if motDePasse != "go2627" {
		fmt.Println("Accès refusé")
	} else {
		fmt.Println("Accès autorisé")
	}
}
