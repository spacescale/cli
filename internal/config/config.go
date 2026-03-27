package config

import "os"

const defaultEnvironment = "development"

type Config struct {
	Environment string
}

func Load() Config {
	environment := os.Getenv("SPACESCALE_ENV")
	if environment == "" {
		environment = defaultEnvironment
	}

	return Config{
		Environment: environment,
	}
}
