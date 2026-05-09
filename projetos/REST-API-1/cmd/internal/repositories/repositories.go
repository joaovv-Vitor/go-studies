package repositories

import (
	"REST-API-1/cmd/internal/models"
	"REST-API-1/cmd/internal/repositories/users"
)

type Repositories struct {
	Users interface {
		GetAll() []models.Users
		Add(newUser models.Users)
		EmailExist(email string) bool
	}
}

func New() *Repositories {
	return &Repositories{
		Users: users.New(),
	}
}