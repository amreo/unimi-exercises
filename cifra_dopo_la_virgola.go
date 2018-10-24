package main

import (
	"fmt"
)

func main() {
	var num float64

	//Vabe...
	fmt.Print("Inserire numero: ")
	fmt.Scan(&num)

	//Determinazione cifra, semplicemente la virgola viene spostata a destra e viene troncata la parte decimale
	cifra := int(num*10) % 10

	//L'ovvietà!
	fmt.Println("La prima cifra dopo la virgola è ", cifra)
}
