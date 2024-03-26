package entity

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID
	Name     string
	UserName string
	Email    string
	Password string
}

// TODO - Verificar se o Email é válido
