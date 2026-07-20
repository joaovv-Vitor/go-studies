package main

import "fmt"

type Lutador struct {
	Nome   string
	Vida   int
	Ataque int
	Defesa int
}

func limitarMinimoZero(valor int) int {
	if valor < 0 {
		return 0
	}
	return valor
}

func calcularDano(atacante, defensor Lutador) int {
	// Dano = ataque do atacante - defesa do defensor.
	// O dano nunca pode ser negativo.
	return limitarMinimoZero(atacante.Ataque - defensor.Defesa)
}

func aplicarDano(lutador Lutador, dano int) Lutador {
	// Subtraia o dano da vida.
	// A vida nunca pode ser negativa.
	// Retorne o lutador atualizado.

	lutador.Vida = limitarMinimoZero(lutador.Vida - dano)

	return lutador
}

func main() {
	kael := Lutador{
		Nome:   "Kael",
		Vida:   100,
		Ataque: 35,
		Defesa: 10,
	}

	nyra := Lutador{
		Nome:   "Nyra",
		Vida:   80,
		Ataque: 28,
		Defesa: 12,
	}

	dano := calcularDano(kael, nyra)
	nyra = aplicarDano(nyra, dano)

	fmt.Printf("%s causou %d de dano\n", kael.Nome, dano)
	fmt.Printf("Vida restante de %s: %d\n", nyra.Nome, nyra.Vida)
}
