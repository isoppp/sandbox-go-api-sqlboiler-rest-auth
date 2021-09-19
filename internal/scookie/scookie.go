package scookie

import (
	"github.com/gorilla/securecookie"
)

func NewSecureCookie() *securecookie.SecureCookie {
	var hashKey = []byte("jkb2kJU4C6ad11DOFElCYMhtF8kvw75n")
	return securecookie.New(hashKey, nil)
}
