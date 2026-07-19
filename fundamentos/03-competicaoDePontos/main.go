package main

import "fmt"

const pontosPorVitoria = 100
const pontosPorEmpate = 25
const penalidadePorFalha = 15

func calcularPontuacao(vitorias, empates, falhas int) int {
	// Calcule:
	// pontos das vitórias
	// + pontos dos empates
	// - penalidades pelas falhas
	return vitorias*pontosPorVitoria + empates*pontosPorEmpate - falhas*penalidadePorFalha
}

func calcularMedia(pontuacao, rodadas int) (float64, bool) {
	if rodadas == 0 {
		return 0, false
	}
	mediaPontuacao := float64(pontuacao) / float64(rodadas)
	return float64(mediaPontuacao), true
}

func main() {
	robo := "Ferrugem-7"
	vitorias := 3
	empates := 2
	falhas := 4
	rodadas := 9

	pontuacao := calcularPontuacao(vitorias, empates, falhas)
	media, calculada := calcularMedia(pontuacao, rodadas)

	fmt.Printf("Robô: %s\n", robo)
	fmt.Printf("Pontuação: %d\n", pontuacao)
	if calculada {
		fmt.Printf("Media do robô foi de %.2f\n", media)
	} else {
		fmt.Println("Não foi possível calcular a média")
	}
}
