package config

type HttpConfig struct {
	Host				string 						`yaml:"host" env:"HTTP_HOST" env-required:"true"`
	Port 				string 						`yaml:"port" env:"HTTP_PORT" env-required:"true"`
}
