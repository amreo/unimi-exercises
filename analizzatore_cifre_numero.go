package main

import (
	"fmt"
)

func main() {
	var num int

	//Va be, lo sappiamo cosa fanno queste due linee...
	fmt.Print("Inserire numero/anno: ")
	fmt.Scan(&num)

	//Verifica che sia un multiplo di 1000
	if num%1000 == 0 {
		fmt.Println("Finisce per 000")
	} else {
		fmt.Println("Non finisce per 000")
	}

	//Calcolo cifre
	//La cifra meno significativa è num0!
	num2, num1, num0 := num/100%10, num/10%10, num%10
	somma := num0 + num1 + num2
	fmt.Println("La somma delle cifre è ", somma)

	//Dato che il numero è anche un anno, viene verificato che sia bisestile
	if num%400 == 0 || (num%4 == 0 && num%100 != 0) {
		fmt.Println("L'anno è bisestile")
	} else {
		fmt.Println("L'anno non è bisestile")
	}
}
