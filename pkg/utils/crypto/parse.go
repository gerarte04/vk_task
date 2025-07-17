package crypto

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func ParsePrivateKeyFromPEM(data string) ([]byte, error) {
	const op = "ParsePrivateKeyFromPEM"

	block, _ := pem.Decode([]byte(data))
    key, err := x509.ParsePKCS8PrivateKey(block.Bytes)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

    return key.(ed25519.PrivateKey), nil
}

func ParsePublicKeyFromPEM(data string) ([]byte, error) {
	const op = "ParsePublicKeyFromPEM"

	block, _ := pem.Decode([]byte(data))
	key, err := x509.ParsePKIXPublicKey(block.Bytes)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

    return key.(ed25519.PublicKey), nil
}
