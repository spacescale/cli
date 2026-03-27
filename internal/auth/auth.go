package auth

import "os"

const TokenEnvVar = "SPACESCALE_TOKEN"

type Credentials struct {
	Token string
}

func LoadFromEnv() Credentials {
	return Credentials{
		Token: os.Getenv(TokenEnvVar),
	}
}
