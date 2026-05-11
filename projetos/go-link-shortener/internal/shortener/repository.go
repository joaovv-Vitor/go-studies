package shortener

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrURLNotFound = errors.New("url não encontrada")

// repository (com letra minúscula) é a nossa implementação concreta e privada.
type repository struct {
	pool *pgxpool.Pool
}

// NewRepository retorna a interface Repository definida no seu model.go.
// Visualmente, fica super limpo quando você chamar no main.go.
func NewRepository(pool *pgxpool.Pool) Repository {
	return &repository{pool: pool}
}

func (r *repository) Save(ctx context.Context, url *URL) error {
	query := `
		INSERT INTO urls (original_url, short_code, created_at)
		VALUES ($1, $2, $3)
		RETURNING id`

	err := r.pool.QueryRow(ctx, query, url.OriginalURL, url.ShortCode, url.CreatedAt).Scan(&url.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetByCode(ctx context.Context, code string) (*URL, error) {
	query := `
		SELECT id, original_url, short_code, created_at
		FROM urls
		WHERE short_code = $1`

	var url URL
	err := r.pool.QueryRow(ctx, query, code).Scan(
		&url.ID,
		&url.OriginalURL,
		&url.ShortCode,
		&url.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrURLNotFound
		}
		return nil, err
	}

	return &url, nil
}
