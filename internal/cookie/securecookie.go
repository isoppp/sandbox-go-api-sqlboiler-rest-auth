package cookie

import (
	"sandbox-go-api-sqlboiler-rest-auth/internal/config"

	"github.com/gorilla/securecookie"
)

func NewSecureCookie(cfg *config.Config) *securecookie.SecureCookie {
	var hashKey = []byte(cfg.SecretKey)
	return securecookie.New(hashKey, nil)
}
