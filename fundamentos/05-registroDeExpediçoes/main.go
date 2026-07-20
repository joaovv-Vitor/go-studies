package main

import "fmt"

func registrarExpedicao(
	expedicoes map[string]int,
	planeta string,
) {
	// Aumente em 1 a quantidade de expedições do planeta.
	expedicoes[planeta]++
}

func consultarExpedicoes(
	expedicoes map[string]int,
	planeta string,
) (int, bool) {
	// Consulte o map e retorne quantidade e existência.
	quantidade, existe := expedicoes[planeta]
	if existe {
		return quantidade, existe
	}
	return 0, false
}

func removerPlaneta(
	expedicoes map[string]int,
	planeta string,
) bool {

	_, existe := expedicoes[planeta]
	if existe {
		delete(expedicoes, planeta)
		return true
	}
	return false
}

func main() {
	expedicoes := map[string]int{
		"Marte":  2,
		"Europa": 1,
	}

	registrarExpedicao(expedicoes, "Marte")
	registrarExpedicao(expedicoes, "Titã")

	quantidade, existe := consultarExpedicoes(expedicoes, "Marte")

	if existe {
		fmt.Printf("Marte possui %d expedições\n", quantidade)
	} else {
		fmt.Println("Planeta não cadastrado")
	}
}
