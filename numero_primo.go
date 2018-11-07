package main
    
import (
	"fmt"
)

func main() {
	var n, i, cicli int

	//Input
	fmt.Print("Inserire numero: ")
	fmt.Scan(&n)

	//Verifica iniziale che sia pari ma non dispari.
	if n%2 == 0 && n != 2 {
		//Non primo, pari e diverso da 2
		fmt.Println("Non è primo. Cicli=0")
	} else {
		//Cerca i divisori fino a radN
		for i = 3; i*i <= n; i += 2 {
			cicli++

			if n%i == 0 {
				break
			}
		}

		//Se i<=radN significa che il ciclo è finito prima quindi non è primo
		if i*i <= n {
			//Non è primo
			fmt.Println("Non è primo. Cicli=", cicli)
		} else {
			//È primo
			fmt.Println("È primo. Cicli=", cicli)
		}
	}
}
