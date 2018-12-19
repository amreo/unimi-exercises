package main

import (
	"fmt"
	"os"
	"strconv"
)

type orario struct {
	secondi int
	minuti  int
	ore     int
}

func newOrario(secondi int, minuti int, ore int) (orario, bool) {
	if secondi >= 0 && secondi <= 59 && minuti >= 0 && minuti <= 59 && ore >= 0 && ore <= 23 {
		return orario{secondi: secondi, minuti: minuti, ore: ore}, true
	} else {
		return orario{}, false
	}
}

func tic(orarioCorrente orario) orario {
	orarioFuturo := orarioCorrente
	orarioFuturo.secondi = (orarioFuturo.secondi + 1) % 60
	if orarioFuturo.secondi == 0 {
		orarioFuturo.minuti = (orarioFuturo.minuti + 1) % 60
		if orarioFuturo.minuti == 0 {
			orarioFuturo.ore = (orarioFuturo.ore + 1) % 24
		}
	}

	return orarioFuturo
}

func main() {
	secondi, _ := strconv.Atoi(os.Args[1])
	minuti, _ := strconv.Atoi(os.Args[2])
	ore, _ := strconv.Atoi(os.Args[3])
	steps, _ := strconv.Atoi(os.Args[4])
	oraCorrente, ok := newOrario(secondi, minuti, ore)

	if !ok {
		fmt.Println("parametri non validi")
		return
	}

	fmt.Printf("%d:%d:%d\n", oraCorrente.secondi, oraCorrente.minuti, oraCorrente.ore)
	for i := 0; i < steps; i++ {
		oraCorrente = tic(oraCorrente)
	}
	fmt.Printf("%d:%d:%d\n", oraCorrente.secondi, oraCorrente.minuti, oraCorrente.ore)
}
