package jwtModelsClaims

import "github.com/dgrijalva/jwt-go"

type JwtCustomClaims struct {
	ID       int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Username string `boil:"username" json:"username" toml:"username" yaml:"username"`
	Email    string `boil:"email" json:"email" toml:"email" yaml:"email"`

	jwt.StandardClaims
}
