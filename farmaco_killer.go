package main

import (
	"fmt"
)

//Struct che modellano i vari tipi di dati
type somministrazione struct {
	farmaco string
	qty     float64
}

type paziente struct {
	nome             string
	cognome          string
	somministrazioni []somministrazione
}

//Funzione main
func main() {
	//Input dei dati
	limitiFarmaci := scanLimitiFarmaci()
	pazienti := scanPazienti()
	//Generazione elenco dei pazienti uccisi
	pazientiUccisi := killer(pazienti, limitiFarmaci)
	//Scrittura in out dei pazienti uccisi
	printNomeCognomePazienti(pazientiUccisi)
}

//Funzione che raccoglie la lista dei limiti di farmaci dallo stdin
func scanLimitiFarmaci() map[string]float64 {
	//Variabili che contengono l'input dallo stdin
	var (
		nome string
		qty  float64
	)

	fmt.Println("Inserire tutte le coppie di limiti in formato nome-limite fino a coppia FINE 0")

	//Creazione lista dei limiti dei farmaci
	lista := make(map[string]float64)
	//Lettura di tutti i limiti dallo stdin. Si assume l'univocità della riga
	for !scanCoppiaFarmaco(&nome, &qty) {
		lista[nome] = qty
	}
	return lista
}

//Scan dei dati dei pazienti
func scanPazienti() []paziente {
	var (
		pazienti []paziente //Lista dei pazienti inseriti
		//Dati temporanei dallo stdin
		tempPaziente paziente
		tempSomm     somministrazione
	)

	fmt.Println("Inserire pazienti e le relative assunzioni")

	//Inserimento pazienti finchè la scan si interrompe perchè è stato inserito FINE FINE
	for !scanInformazioniBasePaziente(&tempPaziente) {
		//Cancellazione dei vecchi dati delle somministrazioni
		tempPaziente.somministrazioni = make([]somministrazione, 0)
		//Inserimento di tutte le somministazioni
		for !scanCoppiaFarmaco(&tempSomm.farmaco, &tempSomm.qty) {
			tempPaziente.somministrazioni = append(tempPaziente.somministrazioni, tempSomm)
		}
		//Aggiunta del paziente
		pazienti = append(pazienti, tempPaziente)
	}
	return pazienti
}

//Funzione che restituisce la lista dei pazienti uccisi dato la lista dei pazienti e i limiti dei farmaci
func killer(pazienti []paziente, limiti map[string]float64) []paziente {
	//Variabile che tiene traccia di tutti i pazienti uccisi
	pazientiUccisi := make([]paziente, 0)

	//Per tutti i pazienti...
	for _, paz := range pazienti {
		//Viene creata la mappa che traccia la quantità di ogni farmaco assunto
		farmaciAssunti := make(map[string]float64)
		//Calcolo qty assunta di ogni farmaco
		for _, somm := range paz.somministrazioni {
			farmaciAssunti[somm.farmaco] += somm.qty
		}

		//Controllo limite assunzione
		assunzioneOk := true
		//Per tutti i farmati assunti...
		for farm, qty := range farmaciAssunti {
			//Verifica se non è stata superata la soglia
			if qty >= limiti[farm] {
				//In tal caso l'assunzione non è ok
				assunzioneOk = false
				break
			}
		}

		//Se l'assunzione non è a norma, il paziente è stato ucciso
		if !assunzioneOk {
			pazientiUccisi = append(pazientiUccisi, paz)
		}
	}

	return pazientiUccisi
}

//Funzioni di utility dal nome chiaro.
func printNomeCognomePazienti(pazienti []paziente) {
	for _, paz := range pazienti {
		fmt.Println(paz.nome, paz.cognome)
	}
}

//Recupera il nome e cognome del paziente dallo stdin. Restituisce true se è stato immesso FINE FINE
func scanInformazioniBasePaziente(paz *paziente) bool {
	fmt.Scan(&paz.nome, &paz.cognome)
	return paz.nome == "FINE" && paz.cognome == "FINE"
}

//Recupera il farmaco e quantità dallo stdin. Restituisce true se è stato immesso FINE FINE
func scanCoppiaFarmaco(nome *string, qty *float64) bool {
	fmt.Scan(nome, qty)
	return (*nome) == "FINE" && (*qty) == 0.0
}
