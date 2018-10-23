package main

import (
	"fmt"
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
	if AAx > ABx {
		temp := AAx
		AAx = ABx
		ABx = temp
	}
	if AAy > ABy {
		temp := AAy
		AAy = ABy
		ABy = temp
	}
	if BAx > BBx {
		temp := BAx
		BAx = BBx
		BBx = temp
	}
	if BAy > BBy {
		temp := BAy
		BAy = BBy
		BBy = temp
	}

	//Calcolo basi, altezze
	baseA := (ABx - AAx)
	altezzaA := (ABy - AAy)
	baseB := (BBx - BAx)
	altezzaB := (BBy - BAy)

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
	fmt.Println("AAx=", AAx, " AAy=", AAy, " ABx=", ABx, " ABy=", ABy, "BAx=", BAx, " BAy=", BAy, " BBx=", BBx, " BBy=", BBy)

	//Verifica la non sovrapposizione rettangoli
	if ABx <= BAx || BBx <= AAx || ABy <= BAy || BBy <= AAy {
		fmt.Println("I rettangoli non si sovrappongono")
	} else {
		fmt.Println("Un rettangolo sovrappone completamente l'altro oppure coincidono")
		//Calcolo vertici della intersezione
		var verAx, verAy, verBx, verBy float64

		//Qui converrebbe usare la funzione math.min
		if AAx > BAx {
			verAx = AAx
		} else {
			verAx = BAx
		}
		if AAy > BAy {
			verAy = AAy
		} else {
			verAy = BAy
		}
		if ABx < BBx {
			verBx = ABx
		} else {
			verBx = BBx
		}
		if ABy < BBy {
			verBy = ABy
		} else {
			verBy = BBy
		}
		fmt.Println("verAx=", verAx, " verAy=", verAy, "verBx=", verBx, " verBy=", verBy)
		areaSovrapposizione := (verBx - verAx) * (verBy - verAy)
		fmt.Println("Area di sovrapposizone: ", areaSovrapposizione)
	}
}
