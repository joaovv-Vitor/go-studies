package main

import (
	"errors"
	"fmt"
)

func calcularSuprimentos(
	tripulantes,
	dias int,
) (int, error) {
	// Se tripulantes for menor ou igual a zero:
	// retorne 0 e um erro.
	if tripulantes <= 0 {
		return 0, errors.New("quantidade de tripulantes deve ser maior que zero")
	}

	// Se dias for menor ou igual a zero:
	// retorne 0 e um erro.
	if dias <= 0 {
		return 0, errors.New("quantidade de dias deve ser maior que zero")
	}

	// Cada tripulante consome 3 unidades por dia.
	// Retorne o total e nil.
	return tripulantes * dias * 3, nil
}

func main() {
	total, err := calcularSuprimentos(5, -1)

	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Printf("Suprimentos necessários: %d\n", total)
}
