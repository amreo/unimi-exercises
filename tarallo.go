package main

import (
	"fmt"
	"math"
)

func main() {
	var k int
	var e float64
	fmt.Println("Inserisci k")
	fmt.Scan(&k)
	/*calcolo approssimativo con n=10000000 (approssimativo)*/
	e = math.Pow(1.0+1.0/10000000.0, 10000000.0)
	println(math.Floor(e*math.Pow10(k)) / math.Pow10(k))
}
