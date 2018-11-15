package main

import (
	"fmt"
)

//Struc che contiene una data
type data struct {
	dd   int //[0; 30]
	mm   int //[0; 11]
	yyyy int
}

func main() {
	//Inserimento date
	data1 := inputData("Inserire data1: ")
	data2 := inputData("Inserire data2: ")

	//Scrive sullo stdout la distanza delle date
	fmt.Println("Distanza giorni: ", calcolaDistanza(data1, data2))
}

///Funzione che restituisce la data immessa dall'utente.
//msg: messaggio da dire all'utente prima della immmissione della data
func inputData(msg string) (d data) {
	fmt.Print(msg)
	fmt.Scan(&d.dd, &d.mm, &d.yyyy)
	//Sistema il valore del giorno e del mese. L'utente ragiona i giorni in termini di [1; 31], noi [0; 30]
	//Stessa cosa con i mesi
	d.dd--
	d.mm--

	return
}

//Funzione che restituisce la distanza in giorni delle date. Se la prima data > seconda data restituisce un numero negativo
func calcolaDistanza(data1 data, data2 data) int {
	//Verifica l'ordine delle date con compare
	//Ho messo < 0 | > 0 per motivi di abitudine. Tendo ad assumere il risultato di compare implementation dependent
	//In questo caso il codice funzionerebbe anche se la condizione è == -1 o == 1
	if val := compare(data1, data2); val < 0 {
		return calcolaDistanzaRaw(data1, data2)
	} else if val > 0 {
		return -calcolaDistanzaRaw(data2, data1)
	} else {
		return 0
	}
}

//Restituisce -1 se data1<data2; 0 se data1==data2; 1 se data1>data2
func compare(data1 data, data2 data) int {
	//Confronta in ordine yyyy, mm, dd e restituisce -1, 1 o 0 in base al caso
	switch {
	case data1.yyyy < data2.yyyy:
		return -1
	case data1.yyyy > data2.yyyy:
		return 1
	case data1.mm < data2.mm:
		return -1
	case data1.mm > data2.mm:
		return 1
	case data1.dd < data2.dd:
		return -1
	case data1.dd > data2.dd:
		return 1
	default:
		return 0
	}
}

//Calcolo distanza anni assumendo che data1 < data2
//Non ho trovato un nome migliore per la funzione...
func calcolaDistanzaRaw(data1 data, data2 data) int {
	//L'algoritmo è abbastanza naïve/stupido. Continua a spostare la data1 finche raggiunge data2
	//Il ciclo si può implementare con un for terario ma esce con le tre parti un pochino illeggibili
	data := data1 //Il cursore che tiene traccia della data corrente
	i := 0        //Numero di giorni passati
	for compare(data, data2) < 0 {
		//Avanza la data di un giorno
		data = nextData(data)
		//Incrementa il numero di giorni passati
		i++
	}

	//Restituisce il numero di giorni
	return i
}

//Funzione che restituisce la data successiva a data
func nextData(data data) data {
	//Avanza di un giorno
	data.dd++
	//Verifica che ha superato l'ultimo giorno del mese
	if data.dd > ultimoGiorniMese(data.mm, data.yyyy) {
		//Aggiorna il mese
		data.mm++
		data.dd = 0
	}
	//Verifica che non ha superato il mese
	if data.mm > 11 {
		//avanza di un anno
		data.mm = 0
		data.yyyy++
	}

	//Restitusce la data aggiornata
	return data
}

//Funzione che data un mese e l'anno, restituisce l'ultimo giorno del mese
func ultimoGiorniMese(mm int, yyyy int) int {
	//In base al mese
	switch mm {
	case 0, 2, 4, 6, 7, 9, 11:
		return 30
	case 3, 5, 8, 10:
		return 29
	case 1:
		{
			if isBisestile(yyyy) {
				return 28
			} else {
				return 27
			}
		}
	default:
		fmt.Errorf("ultimiGiorniMese() richiamato con mese errato")
		return -1
	}

}

//Funzione che restituisce true se l'anno è bisestile, altrimenti false
func isBisestile(yyyy int) bool {
	return yyyy%400 == 0 || (yyyy%4 == 0 && yyyy%100 != 0)
}
