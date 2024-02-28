package biometric

import (
	"github.com/gin-gonic/gin"
	"pthw.com/asymmetric-for-biometric/internal/biometric"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc biometric.UseCase) {
	h := NewHandler(uc)

	authEndpoints := router.Group("/user")
	{
		authEndpoints.POST("/biometric", h.Create)
	}
}
