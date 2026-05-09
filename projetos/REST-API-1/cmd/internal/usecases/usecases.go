package usecases

import (
	"REST-API-1/cmd/internal/models"
	"REST-API-1/cmd/internal/repositories"
	"errors"
	"log/slog"

	"github.com/google/uuid"
)

type UseCases struct {
	repos *repositories.Repositories
}

func New(repos *repositories.Repositories) *UseCases {
	return &UseCases{
		repos: repos,
	}
}

func (u UseCases) Add(newUser models.CreateUserRequest) (uuid.UUID, error) {
	exists := u.repos.Users.EmailExist(newUser.Email)

	if exists {
		slog.Error("this user already exist", "email", newUser.Email)
		return uuid.Nil, errors.New("user already exist")
	}

	repoReq := models.Users{
		ID:   uuid.New(),
		Name: newUser.Name,
	}

	u.repos.Users.Add(repoReq)

	return repoReq.ID, nil
}

func (u *UseCases) GetAll() []models.Users {
	return u.repos.Users.GetAll()
}
