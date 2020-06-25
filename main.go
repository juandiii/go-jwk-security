// Thanks to Fiber JWT https://github.com/gofiber/jwt

package security

import (
	"errors"
	"reflect"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
)

type Config struct {
	keyFunc        jwt.Keyfunc
	Claims         jwt.Claims
	SuccessHandler func(c *fiber.Ctx)
	ErrorHandler   func(c *fiber.Ctx, err error)
}

// JwtMiddleware
func JwtMiddleware(config ...Config) func(*fiber.Ctx) {
	var cfg Config
	headers := "header:" + fiber.HeaderAuthorization
	authScheme := "Bearer"

	if len(config) > 0 {
		cfg = config[0]
	}

	if cfg.SuccessHandler == nil {
		cfg.SuccessHandler = func(c *fiber.Ctx) {
			c.Next()
		}
	}

	if cfg.ErrorHandler == nil {
		cfg.ErrorHandler = func(c *fiber.Ctx, err error) {
			if err.Error() == "Missing or malformed JWT" {
				c.SendStatus(fiber.StatusBadRequest)
				c.SendString("Missing or malformed JWT")
			} else {
				c.Status(fiber.StatusUnauthorized)
				c.SendString("Invalid or expired JWT")
			}
		}
	}

	if cfg.keyFunc == nil {
		// cfg.keyFunc = cfg.keyFunc
	}

	parts := strings.Split(headers, ":")

	t := jwtFromHeaders(parts[1], authScheme)

	return func(c *fiber.Ctx) {
		auth, err := t(c)
		cfg.Claims = jwt.MapClaims{}

		token := new(jwt.Token)

		if _, ok := cfg.Claims.(jwt.MapClaims); ok {
			token, err = jwt.Parse(auth, cfg.keyFunc)
		} else {
			t := reflect.ValueOf(cfg.Claims).Type().Elem()
			claims := reflect.New(t).Interface().(jwt.Claims)
			token, err = jwt.ParseWithClaims(auth, claims, cfg.keyFunc)
		}

		if err == nil && token.Valid {
			c.Locals("user", token)
			cfg.SuccessHandler(c)
			return
		}

		cfg.ErrorHandler(c, err)
		return
	}
}

func jwtFromHeaders(header string, authScheme string) func(c *fiber.Ctx) (string, error) {
	return func(c *fiber.Ctx) (string, error) {
		auth := c.Get(header)

		l := len(authScheme)
		if len(auth) > l+1 && auth[:l] == authScheme {
			return auth[l+1:], nil
		}

		return "", errors.New("Missing or malformed JWT")
	}
}
