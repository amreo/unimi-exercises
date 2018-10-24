package main

import (
	"fmt"
	"math"
)

func main() {
	var AAx, AAy, ABx, ABy float64
	var BAx, BAy, BBx, BBy float64
	//Inserimenti vari
	fmt.Print("Inserimento coordinate del rettangolo A in formato Ax Ay Bx By: ")
	fmt.Scan(&AAx, &AAy, &ABx, &ABy)
	fmt.Print("Inserimento coordinate del rettangolo B in formato Ax Ay Bx By: ")
	fmt.Scan(&BAx, &BAy, &BBx, &BBy)

	//Eventuale inversion di dei valori che devono rispettare la proprietà:
	// Ax < Bx, Ay < By
	//Può essere fatta dopo
	AAx, ABx = math.Min(AAx, ABx), math.Max(AAx, ABx)
	AAy, ABy = math.Min(AAy, ABy), math.Max(AAy, ABy)
	BAx, BBx = math.Min(BAx, BBx), math.Max(BAx, BBx)
	BAy, BBy = math.Min(BAy, BBy), math.Max(BAy, BBy)

	//Calcolo basi, altezze
	baseA, altezzaA := (ABx - AAx), (ABy - AAy)
	baseB, altezzaB := (BBx - BAx), (BBy - BAy)

	//Calcolo aree
	areaA := baseA * altezzaA
	areaB := baseB * altezzaB

	//Visualizzazione su chi è più grande
	if areaA > areaB {
		fmt.Println("Il rettangolo A è più grande del rettangolo B")
	} else if areaA == areaB {
		fmt.Println("Il rettangolo A è grande quanto il rettangolo B")
	} else {
		fmt.Println("Il rettangolo A è più piccolo del rettangolo B")
	}

	//Println di debug
	//fmt.Println("AAx=", AAx, " AAy=", AAy, " ABx=", ABx, " ABy=", ABy, "BAx=", BAx, " BAy=", BAy, " BBx=", BBx, " BBy=", BBy)

	//Verifica la non sovrapposizione rettangoli
	if ABx <= BAx || BBx <= AAx || ABy <= BAy || BBy <= AAy {
		fmt.Println("I rettangoli non si sovrappongono")
	} else {
		fmt.Println("Un rettangolo sovrappone completamente l'altro oppure coincidono")

		//Calcolo estremi del rettangolo intersezione
		verAx, verAy := math.Max(AAx, BAx), math.Max(AAy, BAy)
		verBx, verBy := math.Min(ABx, BBx), math.Min(ABy, BBy)

		//fmt.Println("verAx=", verAx, " verAy=", verAy, "verBx=", verBx, " verBy=", verBy)
		areaSovrapposizione := (verBx - verAx) * (verBy - verAy)
		fmt.Println("Area di sovrapposizone: ", areaSovrapposizione)
	}
}
