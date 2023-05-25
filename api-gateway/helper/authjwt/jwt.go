package authjwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UserID   uint   `json:"user_id"`
	UserRole string `json:"user_role"`
	jwt.RegisteredClaims
}