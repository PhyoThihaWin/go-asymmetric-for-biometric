package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"pthw.com/asymmetric-for-biometric/utils"
	"pthw.com/asymmetric-for-biometric/utils/jwt"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := jwt.IsAuthorized(authToken, secret)
			if authorized {
				// userID, err := jwt.ExtractIDFromToken(authToken, secret)
				// if err != nil {
				// 	utils.ValidatorErrorResponse(ctx, http.StatusUnauthorized, ctx.Request.Method, err.Error())
				// 	return
				// }
				// // ctx.Set("x-user-id", userID)
				ctx.Next()
				return
			}
			utils.ApiErrorResponse(ctx, http.StatusUnauthorized, ctx.Request.Method, err.Error())
			return
		}
		utils.ApiErrorResponse(ctx, http.StatusUnauthorized, ctx.Request.Method, "Not authorized")
	}
}
