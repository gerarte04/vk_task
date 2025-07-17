package domain

import (
	"time"

	"github.com/google/uuid"
)

type AdPrice int

type Ad struct {
	Id uuid.UUID
	AuthorLogin string

	Title string
	Description string
	ImageAddress string
	Price AdPrice

	CreationTime time.Time
}
