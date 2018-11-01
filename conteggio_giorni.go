package main

import (
	"fmt"
)

func main() {
	var dd, mm, yyyy int

	//Input
	fmt.Print("Inserire dd mm yyyy: ")
	fmt.Scan(&dd, &mm, &yyyy)

	//Inizializzazione delle variabili di cicli. I giorni e i mesi partono da 0 per motivi di comodità
	ddi, mmi, yyyyi := 0, 0, 1970
	i := 0
	//L'algoritmo è abbastanza naive e non è particolarmente elaborato.
	//Durante la stesura dell'algoritmo mi è sfuggita la parte "Conteggio differenza giorni" che è stata aggiunta successivamente
	//Siccome io sono pigro non riscrivo l'algoritmo ma lo modifico semplicentente, nonostante si possano saltare degli anni, mesi e giorni...
	//Ha una complessità O(n) dove n è la distanza dei giorni da Unix Epoch
	for ddi < dd-1 || mmi < mm-1 || yyyyi < yyyy {
		//Stampa il giorno corrente
		fmt.Println(i, " ", ddi+1, "/", mmi+1, "/", yyyyi)

		//Avanzamento di giorno
		//Calcolo ddi
		//In base al mese
		// 0) gennaio, di 31 giorni
		// 1) febbraio, di 28 giorni (29 se l'anno è bisestile)
		// 2) marzo, di 31 giorni
		// 3) aprile, di 30 giorni
		// 4) maggio, di 31 giorni
		// 5) giugno, di 30 giorni
		// 6) luglio, di 31 giorni
		// 7) agosto, di 31 giorni
		// 8) settembre, di 30 giorni
		// 9) ottobre, di 31 giorni
		// 10) novembre, di 30 giorni
		// 11) dicembre, di 31 giorni
		// Qui converrebbe usare un array!
		lastDayOfMonth := -1
		if mmi == 0 || mmi == 2 || mmi == 4 || mmi == 6 || mmi == 7 || mmi == 9 || mmi == 11 { //Mesi di 31 giorni
			lastDayOfMonth = 30
		} else if mmi == 3 || mmi == 5 || mmi == 8 || mmi == 10 { //Mesi di 30 giorni
			lastDayOfMonth = 29
		} else if mmi == 1 && (yyyyi%400 == 0 || (yyyyi%4 == 0 && yyyyi%100 != 0)) { //Febbraio durante gli anni bisestili
			lastDayOfMonth = 28
		} else if mmi == 1 { //Febbraio
			lastDayOfMonth = 27
		} else {
			//Caso di sicurezza. L'esecuzione del programma non dovrebbe arrivare qui
			fmt.Errorf("Mese mancante o invalido!")
		}

		//Avanza il giorno, e eventualmente il mese
		if ddi == lastDayOfMonth {
			mmi += 1
		}
		ddi = (ddi + 1) % (lastDayOfMonth + 1)

		//Avanza l'anno se finisce l'anno
		if mmi == 12 {
			mmi = 0
			yyyyi += 1
		}

		//Avanza il conteggio del numero dei giorni
		i++
	}

	//Stampa in out il numero di giorni
	fmt.Println(i, " ", ddi+1, "/", mmi+1, "/", yyyyi)
}
