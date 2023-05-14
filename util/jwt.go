package util

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// JWT_SECRET is the secret used to sign the jwt token
func Token(lvl uint, exp time.Time, acc uint) (string, error) {

	// Create jwt token
	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"lvl": lvl,
		"e_u": exp.Unix(), // Expiration unix
		"acc": acc,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := tk.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func TokenExpiration() time.Time {
	return time.Now().Add(time.Hour * 2)
}

// IsExpired checks if the token is expired
func IsExpired(c *fiber.Ctx) bool {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	num := claims["e_u"].(float64)
	exp := int64(num)

	return time.Now().Unix() > exp
}

// Permission checks if the user has the required permission level
func Permission(c *fiber.Ctx, perm uint) bool {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	num := claims["lvl"].(float64)
	lvl := uint(num)

	return lvl >= perm
}

// TokenDetails returns the account id and permission level of the user
func TokenDetails(c *fiber.Ctx) (acc uint, lvl uint) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	num := claims["acc"].(float64)
	acc = uint(num)

	num = claims["lvl"].(float64)
	lvl = uint(num)

	return acc, lvl
}
