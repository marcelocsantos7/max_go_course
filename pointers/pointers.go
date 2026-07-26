package main

import (
	"fmt"
)

func main() {
	age := 32
	agePointer := &age // cria o pointero para variavel age

	/*
		em uma variavel normal, usa o &var
		para obter o endereço da memória da variável

		E *var para DESREFERENCIAR o ponteiro e usar o valor
		no endereço da variável
	*/

	fmt.Println("Age:", *agePointer) // desreferenciar => usar o VALOR no endereço da mémoria
	fmt.Println("Adult years:", getAdultYears(agePointer))
}

func getAdultYears(age *int) int {
	return *age - 18
}
