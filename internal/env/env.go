package env

import "github.com/joho/godotenv"

func LoadEnv() error {
	err := godotenv.Load(".env")
	return err
}
