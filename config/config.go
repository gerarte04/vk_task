package config

import (
	pkgConfig "marketplace/pkg/config"
	"time"
)

type JwtConfig struct {
	PrivateKeyPEM 			string				`env:"PRIVATE_KEY_PEM" env-required:"true"`
	PublicKeyPEM			string				`env:"PUBLIC_KEY_PEM" env-required:"true"`

	Issuer					string				`yaml:"issuer" env-required:"true"`
	ExpirationTime 			time.Duration		`yaml:"expiration_time" env-required:"true"`
}

type CryptConfig struct {
	HashingCost				int					`yaml:"hashing_cost" env-default:"14"`
}

type ServiceConfig struct {
	PageSize				int					`yaml:"page_size" env-default:"5"`
}

type Config struct {
	HttpCfg	pkgConfig.HttpConfig 	`yaml:"http"`
	JwtCfg JwtConfig 				`yaml:"jwt"`
	CryptCfg CryptConfig 			`yaml:"crypt"`
	SvcCfg ServiceConfig 			`yaml:"service"`
}
