package domain

import (
	"time"

	"github.com/google/uuid"
)

type AdPrice int

type Ad struct {
	Id 				uuid.UUID	`swaggerignore:"true"`
	AuthorLogin 	string		`swaggerignore:"true"`

	Title 			string		`json:"title"`
	Description 	string		`json:"description"`
	ImageAddress 	string		`json:"image_address"`
	Price 			AdPrice		`json:"price"`

	CreationTime 	time.Time	`swaggerignore:"true"`
}
