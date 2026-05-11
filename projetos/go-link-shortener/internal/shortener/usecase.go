package shortener

import (
	"context"
	"go-link-shortener/pkg/hash"
	"time"
	// Ajuste para o nome do seu módulo
)

type useCase struct {
	repo Repository
}

func NewUseCase(repo Repository) UseCase {
	return &useCase{repo: repo}
}

func (u *useCase) CreateShortURL(ctx context.Context, originalURL string) (*URL, error) {
	// 1. Gera o hash a partir da URL original
	shortCode := hash.GenerateShortCode(originalURL)

	// 2. Verifica se esse código já existe no banco
	existingURL, err := u.repo.GetByCode(ctx, shortCode)
	
	if err == nil {
		// Se o código existe e a URL é a mesma, não precisamos salvar de novo!
		// Retornamos a URL que já está no banco.
		if existingURL.OriginalURL == originalURL {
			return existingURL, nil
		}
		
		// Se a URL for DIFERENTE, significa que tivemos uma colisão de Hash.
		// concatenamos o tempo atual na string para forçar um hash totalmente novo.
		shortCode = hash.GenerateShortCode(originalURL + time.Now().String())
	} else if err != ErrURLNotFound {
		// "não encontrado"
		return nil, err
	}

	// 3. Monta a nova entidade (se chegou aqui, o shortCode está livre ou a URL não existia)
	url := &URL{
		OriginalURL: originalURL,
		ShortCode:   shortCode,
		CreatedAt:   time.Now(),
	}

	// 4. Salva no banco de dados
	if err := u.repo.Save(ctx, url); err != nil {
		return nil, err
	}

	return url, nil
}

func (u *useCase) GetOriginalURL(ctx context.Context, code string) (string, error) {
	url, err := u.repo.GetByCode(ctx, code)
	if err != nil {
		return "", err
	}
	
	return url.OriginalURL, nil
}