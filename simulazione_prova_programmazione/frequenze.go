package main

import (
	"fmt"
)

func frequenze(s []string) map[string]int {
	frequenze := make(map[string]int)

	for _, v := range s {
		frequenze[v]++
	}

	return frequenze
}

func main() {
	input := ""
	parole := make([]string, 0)
	_, err := fmt.Scan(&input)
	for err == nil {
		parole = append(parole, input)
		_, err = fmt.Scan(&input)
	}
	fmt.Println(parole)
	fmt.Println(frequenze(parole))
}
