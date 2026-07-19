package main

import "fmt"

func avaliarViagem(combustivel, consumo int) string {
	if combustivel < consumo {
		return "Não é possível realizar a viagem"
	}
	if combustivel == consumo {
		return "Viagem realizada com exatidão"
	}
	return "Viagem realizada com sucesso"
}

func calcularReserva(combustivel, consumo int) int {
	return combustivel - consumo
}

func main() {
	nave := "Aurora"
	combustivel := 60
	consumo := 85

	resultado := avaliarViagem(combustivel, consumo)

	fmt.Printf("Nave: %s\n", nave)
	fmt.Printf("Combustivel: %d\n", combustivel)
	fmt.Printf("Consumo: %d\n", consumo)
	fmt.Printf("Resultado: %s\n", resultado)

	reserva := calcularReserva(combustivel, consumo)

	fmt.Printf("Reserva de combustível: %d\n", reserva)

}
