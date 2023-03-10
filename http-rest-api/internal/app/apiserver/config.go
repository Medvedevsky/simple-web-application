package apiserver

type Config struct {
	WebAddress string `toml:"web_address"` // адрес на котором запускается веб сервер
	LogLevel   string `toml:"log_level"`
}

func NewConfig() *Config {
	return &Config{
		WebAddress: ":8080",
		LogLevel:   "debug",
	}
}
