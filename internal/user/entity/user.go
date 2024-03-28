package entity

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	UserName string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

// TODO - Verificar se o Email é válido
