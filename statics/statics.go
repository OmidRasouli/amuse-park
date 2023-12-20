package statics

import (
	"os"
)

var (
	SecretKey string
)

func Read() {
	SecretKey = os.Getenv("SECRET_KEY")
}
