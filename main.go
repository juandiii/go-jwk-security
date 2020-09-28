package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/juandiii/go-jwk-security/jwt"
	"github.com/juandiii/go-jwk-security/security"
)

type Server struct {
	JwtKey *security.JwtKeys
}

func main() {
	server := &Server{JwtKey: &security.JwtKeys{JwtURL: "https://sso.example.net/realm/protocol/openid-connect/certs"}}
	err := server.JwtKey.GetKeys()

	if err != nil {
		fmt.Println(err)
		return
	}

	app := fiber.New()

	app.Use(jwt.JwtMiddleware(jwt.Config{
		KeyFunc: server.JwtKey.GetKey,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Entró klk")
	})

	app.Listen(":3000")
}
