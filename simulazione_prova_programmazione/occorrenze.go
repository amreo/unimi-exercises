package main

import (
	"fmt"
	"os"
)

func occorrenze(s string, r rune) int {
	//In base alla lunghezza della stringa viene deciso l'occorrenza
	if len(s) == 0 { //Stringa vuota
		return 0
	} else if len(s) == 1 { //Solo un carattere
		if s[0] == byte(r) {
			return 1
		} else {
			return 0
		}
	} else { //len > 0
		c := 0
		//Verifica occorrenza primo carattere
		if s[0] == byte(r) {
			c++
		}
		//Verifica occorrenza ultimo carattere
		if s[len(s)-1] == byte(r) {
			c++
		}
		//Calcolo occorrenze totali
		return c + occorrenze(s[1:len(s)-1], r)
	}
}

func main() {
	s := os.Args[2]
	r := rune(os.Args[1][0])
	fmt.Printf("occorrenze di %c in %s\n%d\n", r, s, occorrenze(s, r))

}
