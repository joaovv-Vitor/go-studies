package shortener

import (
	"context"
	"time"
)

// URL representa a entidade principal do nosso domínio.
type URL struct {
	ID          int       `json:"-"` // Omitimos o ID numérico nas respostas JSON
	OriginalURL string    `json:"original_url"`
	ShortCode   string    `json:"short_code"`
	CreatedAt   time.Time `json:"created_at"`
}

// ShortenRequest é uma struct auxiliar para mapear o JSON que chega no Handler.
type ShortenRequest struct {
	URL string `json:"url"`
}

// Repository define o contrato para a camada de persistência.
// O nosso arquivo postgres.go vai implementar exatamente essas assinaturas.
type Repository interface {
	Save(ctx context.Context, url *URL) error
	GetByCode(ctx context.Context, code string) (*URL, error)
}

type UseCase interface {
	CreateShortURL(ctx context.Context, originalURL string) (*URL, error)
	GetOriginalURL(ctx context.Context, code string) (string, error)
}

type CacheRepository interface {
	Set(ctx context.Context, key string, value string, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
}
