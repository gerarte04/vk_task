package utils

import (
	"errors"
	pkgCrypto "marketplace/pkg/utils/crypto"
	"net/http"
	"strings"
)

var (
	ErrIncorrectAuthHeaderFormat = errors.New("incorrect authorization header format")
)

func ProcessAuthHeader(r *http.Request, publicKey []byte) (string, error) {
	strs := strings.Split(r.Header.Get("Authorization"), " ")

	if len(strs) != 2 || strs[0] != "Bearer" {
		return "", ErrIncorrectAuthHeaderFormat
	}

	sub, err := pkgCrypto.ValidateJwtToken(strs[1], publicKey)
	if err != nil {
		return "", err
	}

	return sub, nil
}
