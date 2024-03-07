package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Secret string
var Expire int

func InitGoDotEnv() error {
	err := godotenv.Load()

	Secret = os.Getenv("ACCESS_TOKEN_SECRET")
	value, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRY_HOUR"))
	Expire = value

	return err
}
