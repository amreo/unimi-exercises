package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Inserire N: ")
	fmt.Scan(&n)

	//Primi N quadrati
	fmt.Print("Somma dei primi N quadrati: ")
	somma := 0
	for i := 1; i <= n; i++ {
		//fmt.Println("#", i, " ", i*i)
		somma += i * i
	}
	fmt.Println(somma)

	//Conto alla rovescia fino a 1
	fmt.Println("Conto alla rovescia da n a 1")
	for i := n; i > 0; i-- {
		fmt.Println(i)
	}

	//Tutte le potenze minori di 2 < n
	fmt.Println("Potenze di 2 fino a <n")
	//Si puo usare int(math.pow(2, i)) che 1 << i
	for i := uint(0); (1 << i) <= n; i++ {
		fmt.Println(1 << i)
	}
}
