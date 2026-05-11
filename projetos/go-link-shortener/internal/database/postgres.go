package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// NewPostgresPool cria e valida a conexão com o banco de dados
func NewPostgresPool(ctx context.Context, connString string) (*pgxpool.Pool, error) {
	// Cria a configuração do pool com a string de conexão
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer parse da config do banco: %w", err)
	}

	// configs da pool
	config.MaxConns = 10
	config.MaxConnIdleTime = time.Minute * 5

	// Inicializa o pool
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar no banco: %w", err)
	}

	// Faz um Ping para garantir que o banco está online e acessível
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("banco de dados não respondeu ao ping: %w", err)
	}

	return pool, nil
}
