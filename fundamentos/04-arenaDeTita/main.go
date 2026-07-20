package main

import (
	"fmt"

	batalha "github.com/joaovv-Vitor/go-studies/fundamentos/04-arenaDeTita/Batalha"
)

func main() {
	nomeAtacante := "Vórtice"
	ataque := 75
	defesaInimiga := 30
	vidaInimiga := 40

	dano := batalha.CalcularDano(ataque, defesaInimiga)
	vidaRestante := batalha.CalcularVidaRestante(vidaInimiga, dano)

	fmt.Printf("Vida restante apos um ataque %d\n", vidaRestante)
	fmt.Printf("%s causou %d pontos de dano\n", nomeAtacante, dano)
}
