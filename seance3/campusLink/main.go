package main

import "fmt"

func main() {
	salle, err := lireSalle("S103,Salle Knuth,A,trente")
	fmt.Println(salle, err)
}
