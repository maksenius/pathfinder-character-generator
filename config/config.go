package config

// Config all config variables.
type Config struct {
	HTTP     HTTPConfig
	LogLevel int
}

type HTTPConfig struct {
	Port int
}
