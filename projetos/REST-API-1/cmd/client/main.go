package main

import (
	"REST-API-1/cmd/internal/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	req := models.CreateUserRequest{
		Name:  "Jovi",
		Email: "jovi@gmail.com",
	}

	b, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("http://localhost:8080/users", "application/json", bytes.NewReader(b))
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusCreated {
		panic("error to create user")
	}

	var responseApi models.CreateUserRespose
	if err := json.NewDecoder(resp.Body).Decode(&responseApi); err != nil {
		panic(err)
	}

	fmt.Println("new user created", responseApi)

}
