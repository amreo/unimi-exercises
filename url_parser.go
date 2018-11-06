package main

import (
	"fmt"
)

func main() {
	var url string

	//Lettura input
	fmt.Print("Inserire URL: ")
	fmt.Scan(&url)

	//Algoritmo è abbastanza stupido. Il miglior modo per estrarre il nome host (e altre info) è utilizzare una FSM
	//Ma sono pigro quindi cerco l'inizio e la fine del nome host.
	//Sono ancora pigro per cui utilizzo un solo ciclo for
	ricercaInizio := true
	//Per tutti i caratteri...
	for i := 0; i < len(url); i++ {

		//Trova l'inizio dell'host solo dopo i tre simboli ://.
		//L'inizio va ricercato solo fino ai 3 caratteri prima di len(url)-3 per evitare buffer overflow
		//Ovviamente va ricercato quando lo stiano ricercando
		if ricercaInizio && i <= len(url)-3 && url[i] == ':' && url[i+1] == '/' && url[i+2] == '/' {
			//Ha trovato l'inizio. Skippiamo i 3 caratteri successivi
			//+2 perché la i viene incrementata ancora alla fine dell'iterazione.
			i += 2
			ricercaInizio = false
		} else if !ricercaInizio && url[i] != '/' {
			//Stiamo cercando la fine
			//Ripeto: ho messo la condizione per evitare di mettere un break per uscire dal ciclo quando i >len(url-3)
			fmt.Printf("%c", url[i])
		} else if !ricercaInizio && url[i] == '/' {
			//Non sembra ma è un break
			i = len(url)
		}
	}

	//Un '\n' ci vuole!
	fmt.Println()
}
