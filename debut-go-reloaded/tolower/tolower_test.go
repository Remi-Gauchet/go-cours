package tolower

import (
	"fmt"
	"testing"
)

func TestToLower(t *testing.T) {
	chaine := []string{"HELLO", "Été", "déjà VU"}
	for i := range chaine {
		fmt.Println(ToLower(chaine[i]))
	}
}
