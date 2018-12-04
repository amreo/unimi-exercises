package main

import (
	"fmt"
	"os"
)

//Funzione che dato una stringa s, calcola tutti i suoi possibili anagrammi
func anagrammi(s string) []string {
	//Gli anagrammi di una stringa vuota è per definizione una stringa vuota
	if s == "" {
		return []string{""}
	}

	//Calcola tutti gli anagrammi della sottostringa dal secondo carattere fino alla fine
	//E gli inserisce nello slice partialAnagrammi
	partialAnagrammi := anagrammi(s[1:])

	//Slice e mappa che tracciano tutti gli anagrammi prodotti
	anag := []string{}
	usati := make(map[string]bool)

	//Per ogni posizione in ogni anagramma parziale...
	for _, pAnag := range partialAnagrammi {
		for i := 0; i <= len(pAnag); i++ {
			//...viene prodotta un nuovo anagramma con la prima lettera di s nella posizione
			nuovoAnagramma := pAnag[:i] + string(s[0]) + pAnag[i:]

			//Viene inserito il nuovo anagramma se non è stato usato
			if !usati[nuovoAnagramma] {
				anag = append(anag, nuovoAnagramma)
				usati[nuovoAnagramma] = true //L'anagramma viene contrassegnato come usato
			}
		}
	}

	//Restituisce gli anagrammi trovati
	return anag
}

func main() {
	//Calcola tutti gli anagrammi di os.Args[1] e li scrive nello stdout
	for _, anag := range anagrammi(os.Args[1]) {
		fmt.Println(anag)
	}

}
