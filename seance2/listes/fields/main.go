package main

import (
	"fmt"
	"strings"
)

func main() {
	var phrase string = "le chat dort"
	mots := strings.Fields(phrase) // ["le", "chat", "dort"]
	fmt.Println(len(mots))
	fmt.Println(strings.Join(mots, " "))
}
