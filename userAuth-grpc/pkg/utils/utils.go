package utils

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"math/rand"
	"time"
)

func Wrap(err error, message string) error {
	return errors.Wrap(err, "==> "+message)
}

// GenerateRandom4DigitCode generate 4 digit code for OTP
func GenerateRandom4DigitCode() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(9000) + 1000
}

func GetUDID() string {
	return uuid.New().String()
}
