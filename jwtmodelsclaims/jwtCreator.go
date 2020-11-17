package jwtModelsClaims

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWTCreator it return token string and err
func JWTCreator(userClaims JwtCustomClaims) (string, error) {

	claims := &JwtCustomClaims{
		userClaims.ID,
		userClaims.Email,
		userClaims.Username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secretToken"))
	if err != nil {
		return "", err
	}
	return t, err

}
