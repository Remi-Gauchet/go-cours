package main

import (
	"fmt"
	"strings"
)

func direBonjour(prenom string) {
	fmt.Println("Bonjour", prenom, "!")
}

func doubler(nombre int) int {
	return nombre * 2
}

func crier(texte string) string {
	cricri := string(strings.ToUpper(texte) + " !")
	return cricri
}

func estLong(mot string) bool {
	return len(mot) > 6
}

func main() {
	direBonjour("Monique")
	fmt.Println(doubler(21))
	fmt.Println(crier("ta gueule"))
	if estLong("chat") == true {
		fmt.Println("mot long")
	} else {
		fmt.Println("mot court")
	}
}
