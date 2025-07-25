package config

import (
	pkgConfig "marketplace/pkg/config"
	"marketplace/pkg/database/postgres"
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
	MaxPrice				int					`yaml:"max_price" env-default:"10000000"`
	MaxTitleLength			int					`yaml:"max_title_length" env-default:"100"`
	MaxDescriptionLength	int					`yaml:"max_description_length" env-default:"2000"`
	MaxImageSize			int					`yaml:"max_image_size" env-default:"256"`

	MinLoginLength			int					`yaml:"min_login_length" env-default:"3"`
	MaxLoginLength			int					`yaml:"max_login_length" env-default:"30"`
	MinPasswordLength		int					`yaml:"min_password_length" env-default:"8"`
	MaxPasswordLength		int					`yaml:"max_password_length" env-default:"30"`
	SpecialSymbols			string				`yaml:"special_symbols" env-default:"!@#$%^&*?/"`

	DebugMode				bool				`yaml:"debug_mode" env-default:"true"`
}

type PathConfig struct {
	ApiPath					string				`yaml:"api" env-required:"true"`
	RegisterPath			string				`yaml:"register" env-required:"true"`
	LoginPath				string				`yaml:"login" env-required:"true"`
	CreateAdPath			string				`yaml:"create_ad" env-required:"true"`
	GetFeedPath				string				`yaml:"get_feed" env-required:"true"`
}

type Config struct {
	HttpCfg	pkgConfig.HttpConfig 		`yaml:"http"`
	PostgresCfg postgres.PostgresConfig	`yaml:"postgres"`
	JwtCfg JwtConfig 					`yaml:"jwt"`
	CryptCfg CryptConfig 				`yaml:"crypt"`
	SvcCfg ServiceConfig 				`yaml:"service"`
	PathCfg PathConfig					`yaml:"paths"`
}
