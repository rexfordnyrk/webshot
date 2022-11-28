package capture
type Config struct {
	Size string
	Width	int
	Height	int
	Format 	string
}

func (c Config) ValidateConfig() error {
	return nil
}

func NewConfig() Config {
	return Config{
		Size:   "windowed",
		Width:  1440,
		Height: 900,
		Format: "png",
	}
}
