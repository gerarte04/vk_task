package crypto

import "errors"

var (
	ErrTokenParsingFailed = errors.New("token parsing failed, check if it's correct")
	ErrTokenSignatureInvalid = errors.New("token signature is invalid")
	ErrUnexpectedSigningMethod = errors.New("unexpected signing method")
	ErrTokenExpired = errors.New("Token has been expired")
)
