package middleware

import (
	"loan-app/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
)

func SetupCookieMiddleware(config *config.Config) fiber.Handler {
	return encryptcookie.New(encryptcookie.Config{
		Except:    []string{config.Security.Csrf.CookieName},
		Key:       config.Security.Cookie.Key,
		Encryptor: encryptcookie.EncryptCookie,
		Decryptor: encryptcookie.DecryptCookie,
	})
}
