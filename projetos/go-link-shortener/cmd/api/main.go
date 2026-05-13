package main

import (
	"context"
	"fmt"
	"go-link-shortener/internal/config"
	"go-link-shortener/internal/database"
	"go-link-shortener/internal/shortener"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	ctx := context.Background()

	// 1. Carrega as configurações centralizadas
	cfg := config.Load()

	// 2. Inicia o banco de dados passando a URL validada
	dbPool, err := database.NewPostgresPool(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Erro fatal no banco de dados: %v", err)
	}
	defer dbPool.Close()

	// 3. Inicia a conexão com o Redis
	redisClient, err := database.NewRedisClient(ctx, cfg.RedisURL)
	if err != nil {
		log.Fatalf("Erro ao conectar no Redis: %v", err)
	}
	defer redisClient.Close()

	// 4. Monta as dependências
	repo := shortener.NewRepository(dbPool)
	cacheRepo := shortener.NewRedisRepository(redisClient)
	svc := shortener.NewUseCase(repo, cacheRepo)
	handler := shortener.NewHandler(svc)

	// 4. Configura o Chi
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/api/shorten", handler.CreateShortURL)
	r.Get("/{code}", handler.GetOriginalURL)

	// 5. Inicia o servidor usando a porta da configuração
	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("🚀 Servidor rodando na porta %s...", cfg.Port)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
