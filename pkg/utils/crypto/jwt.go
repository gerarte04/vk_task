package crypto

import (
	"crypto/ed25519"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	method = jwt.SigningMethodEdDSA
)

type TokenClaims struct {
	Iss string
	Sub string
	Exp time.Time
}

func ValidateJwtToken(tokenStr string, publicKey []byte) (uuid.UUID, error) {
	const op = "ValidateJwtToken"

	keyFunc := func(token *jwt.Token) (any, error) {
		if token.Method != method {
			return nil, ErrUnexpectedSigningMethod
		}

		return ed25519.PublicKey(publicKey), nil
	}

	token, err := jwt.Parse(tokenStr, keyFunc,
		jwt.WithValidMethods([]string{method.Alg()}),
		jwt.WithExpirationRequired(),
	)

	if errors.Is(err, jwt.ErrTokenExpired) {
		return uuid.Nil, fmt.Errorf("%s: %w", op, ErrTokenExpired)
	} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
		return uuid.Nil, fmt.Errorf("%s: %w", op, ErrTokenSignatureInvalid)
	} else if err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, ErrTokenParsingFailed)
	}

	sub, err := token.Claims.GetSubject()
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	usrId, err := uuid.Parse(sub)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	return usrId, nil
}

func GenerateJwtToken(cm TokenClaims, privateKey []byte) (string, error) {
	const op = "GenerateJwtToken"

	claims := jwt.RegisteredClaims{
		Issuer: cm.Iss,
		Subject: cm.Sub,
		ExpiresAt: jwt.NewNumericDate(cm.Exp),
		ID: uuid.New().String(),
	}

	token := jwt.NewWithClaims(method, claims)
	tokenStr, err := token.SignedString(ed25519.PrivateKey(privateKey))

	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return tokenStr, nil
}
