package domain

import "github.com/google/uuid"

type User struct {
	Id 				uuid.UUID	`json:"id"`
	Login 			string		`json:"login"`
	PasswordHash 	[]byte		`json:"-"`
}
