# Go Fiber - JWK for Auth

🔐 It's unofficial middleware for Go Fiber
---
Specials thanks for Go Fiber JWT [GoFiber/JWT](https://github.com/gofiber/jwt)

## 📦 What's in the box?

This is for get signature Json Web Key (JWK) [RFC 7517 JSON Web Key (JWK)](https://tools.ietf.org/html/rfc7517) and save signature.

```go
jwt.JwtMiddleware(config ...Config) func(*fiber.Ctx)
```

## 🆙 API
TBT


## 🔽 Get Dependencies

```bash
go get -u github.com/juandiii/go-jwk-security
```

## 👨🏻‍💻 Example 

```go
package main

import (
	"fmt"

	"github.com/gofiber/fiber"
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

	app.Get("/", func(c *fiber.Ctx) {

	})

	app.Listen(3000)
}

```

## ✅ TODO:
1. Improve or clean code
2. Add more functions
3. Fetch Certs after 24h expiration