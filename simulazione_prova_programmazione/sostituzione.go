package main

import (
	"fmt"
	"os"
	"strconv"
)

//La funzione è veramente brutta
func sostituisci(s, old, new []byte, n int) []byte {
	//Strings.Replace andrà bene?

	//Scorre tutto lo slice fino alla fine per trovare una occorrenza di old
	for i := 0; i < len(s); i++ {

		j := 0
		//Verifica la presenza di una occorrenza
		for j = i; j < i+len(old) && j < len(s); j++ {
			//Controlla matching
			if s[j] != old[j-i] {
				break
			}
		}
		//Se l'occorrenza è stata trovata...
		if j == i+len(old) {
			//Se è la prima occorrenza viene sostituita e ciao.
			if n == 1 {
				//Sovrascrive l'occorrenza di old con new
				for k := i; k < i+len(new); k++ {
					s[k] = new[k-i]
				}
				break
			} else {
				n--
			}
			//Avanza l'indice.
			i += len(old) - 1
		}
	}

	return s
}

func main() {
	s := []byte(os.Args[1])
	old := []byte(os.Args[2])
	new := []byte(os.Args[3])
	numeroOccorrenza, _ := strconv.Atoi(os.Args[4])

	fmt.Println(string(s))
	s = sostituisci(s, old, new, numeroOccorrenza)
	fmt.Println(string(s))
}
