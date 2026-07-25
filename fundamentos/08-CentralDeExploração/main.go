package main

import "fmt"

type Explorador interface {
	Explorar() string
}

type Drone struct {
	Nome     string
	Altitude int
}

func (d Drone) Explorar() string {
	// Informe nome e altitude do drone.
	info := fmt.Sprintf(`%s está explorando a %d metros de altitude`,
		d.Nome,
		d.Altitude,
	)
	return info
}

type Rover struct {
	Nome    string
	Planeta string
}

func (r Rover) Explorar() string {
	// Informe nome e planeta do rover.
	info := fmt.Sprintf("%s está explorando a superfície de %s",
		r.Nome,
		r.Planeta,
	)
	return info
}

func iniciarExploracao(explorador Explorador) {
	fmt.Println(explorador.Explorar())
}

func main() {
	drone := Drone{
		Nome:     "Ícaro",
		Altitude: 1200,
	}

	rover := Rover{
		Nome:    "Atlas",
		Planeta: "Marte",
	}

	iniciarExploracao(drone)
	iniciarExploracao(rover)
}
