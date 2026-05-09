package handlers

import (
	usecases "REST-API-1/cmd/internal/usecases"
	"fmt"
	"log/slog"
	"net/http"
)

type Handlers struct {
	usecases *usecases.UseCases
}

func New(usecases *usecases.UseCases) *Handlers {
	return &Handlers{
		usecases: usecases,
	}
}

func (h *Handlers) Listen(port int) error {
	h.registerUserEndpint()

	slog.Info("listening on", "port", port)

	return http.ListenAndServe(
		fmt.Sprintf(":%v", port),
		nil,
	)
}
