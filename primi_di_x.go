package main

import (
	"fmt"
	"math"
)

func main() {
	//Inserimento x
	var x int
	fmt.Scan(&x)

	for i, conta := 2, 0; i <= x; i++ {
		if i == 2 { //È 2
			conta++
		} else if i%2 == 1 { //È dispari
			//Cerca i divisori fino a sqrt(n)
			j := 0
			for j = 3; j*j <= i; j += 2 {
				if i%j == 0 {
					break
				}
			}
			//Se i<=radN significa che il ciclo è finito prima quindi non è primo
			if j*j > i {
				//È primo
				conta++
			}
		}

		asintotoX := float64(i) / math.Log(float64(x))
		rapportoAsintoto := float64(conta) / asintotoX

		fmt.Println(i, ",", conta, ",", asintotoX, ",", rapportoAsintoto)
	}
}
