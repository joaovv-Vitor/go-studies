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

func (l Lutador) CalcularDano(defensor Lutador) int {
	// Retorne ataque menos defesa, limitado a zero.
	return limitarMinimoZero(l.Ataque - defensor.Defesa)
}

func (l *Lutador) ReceberDano(dano int) {
	// Modifique diretamente a vida.
	// Ela nunca pode ficar negativa.

	l.Vida = limitarMinimoZero(l.Vida - dano)

}

func (l Lutador) EstaVivo() bool {
	// Retorne se a vida é maior que zero.

	if l.Vida > 0 {
		return true
	}

	return false
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

	dano := kael.CalcularDano(nyra)
	nyra.ReceberDano(dano)

	fmt.Printf("%s causou %d de dano\n", kael.Nome, dano)
	fmt.Printf("Vida restante de %s: %d\n", nyra.Nome, nyra.Vida)
	fmt.Printf("%s está viva: %t\n", nyra.Nome, nyra.EstaVivo())
}
