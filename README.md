# JWK Security

🔐 It's unoffical middleware for Go Fiber
---
Specials thanks for Go Fiber JWT [GoFiber/JWT](https://github.com/gofiber/jwt)

This is for get signature Json Web Key (JWK) [RFC 7517 JSON Web Key (JWK)](https://tools.ietf.org/html/rfc7517) and save signature.

```go
security.JwtMiddleware(config ...Config) func(*fiber.Ctx)
```

## Get Dependencies

```bash
go get -u github.com/juandiii/go-jwk-security
```

## TODO ✅:
1. Improve or clean code
2. Add more functions
3. Fetch Certs after 24h expiration