package config

type JWTConfig struct {
	Secret        string
	AccessExpire  int // minutes
	RefreshExpire int // minutes
}
