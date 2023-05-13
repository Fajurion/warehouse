package util

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"math/big"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Standard for successful requests
func SuccessfulRequest(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"success": true,
	})
}

// Standard for failed requests
func FailedRequest(c *fiber.Ctx, error string, err error) error {
	return c.Status(200).JSON(fiber.Map{
		"success": false,
		"error":   error,
	})
}

// Standard for invalid requests (invalid payload)
func InvalidRequest(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusBadRequest)
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// Generates a string of random characters
func GenerateToken(tkLength int32) string {

	s := make([]rune, tkLength)

	length := big.NewInt(int64(len(letters)))

	for i := range s {

		number, _ := rand.Int(rand.Reader, length)
		s[i] = letters[number.Int64()]
	}

	return string(s)
}

// HashPassword hashes a password
func HashPassword(password string) string {

	// Get hash of password
	hash := sha256.Sum256([]byte(password))

	// Convert byte[] to string
	return base64.StdEncoding.EncodeToString(hash[:])
}

var Validator *validator.Validate = validator.New()
