package main

import (
	"fmt"
)

//MAPPA costante (più o meno) che mappa i voti alle loro rispettive fascia
//Si può implementare una versione leggermente più efficiente con una stringa piuttosto con una runa,
//dove ogni lettera nella stringa è indicizzata da un voto.
var FASCIA_DI_VALORI map[int]rune = map[int]rune{
	18: 'D', 19: 'D', 20: 'D', 21: 'D',
	22: 'C', 23: 'C', 24: 'C',
	25: 'B', 26: 'B', 27: 'B',
	28: 'A', 29: 'A', 30: 'A',
}

//Funzione main
func main() {
	frequenzaFasce := make(map[rune]int) //mappa che tiene traccia delle fascie
	numeroValori := 0                    //Numero di valori immessi finora
	input := -1                          //Valore immesso dall'input

	//Ciclo di inserimento degli input finché diverso da 0
	for input != 0 {
		fmt.Scan(&input)
		//Se l'input è nella fascia dei voti validi
		if input >= 18 && input <= 30 {
			//Lo aggiunge alla rispettiva fascia
			frequenzaFasce[FASCIA_DI_VALORI[input]]++
			numeroValori++
		}
	}

	//Stampa le info sulle frequenze
	fmt.Printf("A : %d %%\n", frequenzaFasce['A']*100/numeroValori)
	fmt.Printf("B : %d %%\n", frequenzaFasce['B']*100/numeroValori)
	fmt.Printf("C : %d %%\n", frequenzaFasce['C']*100/numeroValori)
	fmt.Printf("D : %d %%\n", frequenzaFasce['D']*100/numeroValori)
}
