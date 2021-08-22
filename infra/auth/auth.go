package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/diegoclair/bank-transfer/util/config"
)

type Key string

func (k Key) String() string {
	return string(k)
}

const (
	UserIDKey       Key = "UserID"
	ContextTokenKey Key = "user-token"
)

var (
	TokenSigningMethod = jwt.SigningMethodRS256
)

func GenerateToken(authCfg config.AuthConfig, claims jwt.Claims) (tokenString string, err error) {

	token := jwt.NewWithClaims(TokenSigningMethod, claims)
	tokenString, err = token.SignedString([]byte(authCfg.PrivateKey))
	if err != nil {
		return tokenString, err
	}

	return tokenString, nil
}
