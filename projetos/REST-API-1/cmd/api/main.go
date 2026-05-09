package main

import (
	"REST-API-1/cmd/internal/handlers"
	"REST-API-1/cmd/internal/repositories"
	"REST-API-1/cmd/internal/usecases"
)

func main(){
	repos := repositories.New()

	useCases := usecases.New(repos)
	
	h := handlers.New(useCases)

	h.Listen(8080)
}