package biometric

import (
	"os"

	"github.com/gin-gonic/gin"
	"pthw.com/asymmetric-for-biometric/internal/biometric"
	"pthw.com/asymmetric-for-biometric/middleware"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc biometric.UseCase) {
	handler := NewHandler(uc)

	authEndpoints := router.Group("/user")
	{
		authEndpoints.POST("/biometric", handler.CreateBiometric)
	}

	jwtEndpoints := authEndpoints
	secret := os.Getenv("ACCESS_TOKEN_SECRET")

	jwtEndpoints.Use(middleware.JwtAuthMiddleware(secret))
	{
		jwtEndpoints.GET("/biometric/challenge/:biometric_id", handler.GetChallenge)
		jwtEndpoints.POST("/biometric/verify", handler.ValidateBiometric)
	}
}
