package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"pthw.com/asymmetric-for-biometric/models"
)

type JwtCustomClaims struct {
	UserBiometric models.UserBiometric
	jwt.MapClaims
}
