package config

const (
	InMemory = "in-memory"
	MongoDB  = "mongo-db"
)

type Config struct {
}

func (c *Config) Validate() bool {
	return false
}
