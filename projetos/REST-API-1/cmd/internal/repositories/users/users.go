package users

import (
	"REST-API-1/cmd/internal/models"
)

type Users struct {
	users []models.Users
}

func New() *Users {
	return &Users{users: make([]models.Users, 0)}

}

func (u Users) GetAll() []models.Users{
	return u.users
}

func (u Users) EmailExist(email string) bool{
	for _, v := range u.users {
		if v.Email == email{
			return true
		}
	}
	return false
}

func (u Users) Add(newUser models.Users){
	u.users = append(u.users, newUser)
}