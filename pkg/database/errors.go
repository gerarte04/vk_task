package database

import "errors"

var (
	ErrForeignKeyViolation = errors.New("Foreign key violation")
	ErrUniqueViolation = errors.New("Unique violation")
	ErrUndocumented = errors.New("Undocumented database error")
)
