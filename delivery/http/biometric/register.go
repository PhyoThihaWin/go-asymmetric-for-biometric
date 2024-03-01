package biometric

import (
	"github.com/gin-gonic/gin"
	"pthw.com/asymmetric-for-biometric/internal/biometric"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc biometric.UseCase) {
	handler := NewHandler(uc)

	authEndpoints := router.Group("/user")
	{
		authEndpoints.POST("/biometric", handler.CreateBiometric)
		authEndpoints.GET("/biometric/challenge/:device_id", handler.GetChallenge)
		authEndpoints.POST("/biometric/verify", handler.ValidateBiometric)
	}
}
