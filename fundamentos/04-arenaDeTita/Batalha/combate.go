package batalha

func CalcularDano(ataque, defesa int) int {
	// Dano é ataque menos defesa.
	// Nunca permita dano negativo.

	dano := ataque - defesa

	return limitarMinimoZero(dano)

}

func CalcularVidaRestante(vida, dano int) int {
	return limitarMinimoZero(vida - dano)
}

func limitarMinimoZero(valor int) int {
	if valor <= 0 {
		return 0
	}
	return valor

}
