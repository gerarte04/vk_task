package types

import "errors"

var (
	ErrBadPageNumber = errors.New("page number has bad format or not specified")
	ErrBadPriceValue = errors.New("bad price value, must be positive and less than max")
	ErrBadTitleLength = errors.New("too large title")
	ErrBadDescriptionLength = errors.New("too large description")
	ErrBadImageFormat = errors.New("bad image format (too big size)")

	ErrBadLoginFormat = errors.New("bad login format")
	ErrBadPasswordFormat = errors.New("bad password format")
)
