package config

import "time"

type JwtConfig struct {
	PrivateKeyPEM 			string			`env:"PRIVATE_KEY_PEM" env-required:"true"`
	PublicKeyPEM			string			`env:"PUBLIC_KEY_PEM" env-required:"true"`

	Issuer					string			`env:"JWT_ISSUER" env-required:"true"`

	ExpirationTime 			time.Duration	`yaml:"access_expiration_time" env:"ACCESS_EXPIRATION_TIME" env-default:"30m"`
}
