package auth

import "github.com/dgrijalva/jwt-go"

type Key string

func (k Key) String() string {
	return string(k)
}

const UserIDKey Key = "UserID"

var (
	// TokenSigningMethod is the auth token signing algorithm
	TokenSigningMethod = jwt.SigningMethodRS256
)
