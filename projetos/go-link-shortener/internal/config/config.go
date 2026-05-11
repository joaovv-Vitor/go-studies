package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config guarda todas as variáveis de ambiente mapeadas e tipadas
type Config struct {
	DatabaseURL string
	Port        string
}

// Load lê o arquivo .env e popula a struct de configuração
func Load() *Config {
	// Tenta carregar o arquivo .env. Se não achar, segue o jogo
	// (útil para rodar em produção no Docker, onde as envs já estão no sistema)
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: arquivo .env não encontrado. Usando variáveis de sistema.")
	}

	cfg := &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        os.Getenv("PORT"),
	}

	// Validações de variáveis obrigatórias
	if cfg.DatabaseURL == "" {
		log.Fatal("Erro fatal: DATABASE_URL não está configurada")
	}

	// Define uma porta padrão caso não seja informada
	if cfg.Port == "" {
		cfg.Port = "8080"
	}

	return cfg
}