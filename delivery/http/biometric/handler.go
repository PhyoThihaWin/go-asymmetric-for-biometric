package biometric

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"pthw.com/asymmetric-for-biometric/internal/biometric"
	"pthw.com/asymmetric-for-biometric/models"
	"pthw.com/asymmetric-for-biometric/utils"
)

type Handler struct {
	useCase biometric.UseCase
}

func NewHandler(useCase biometric.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) CreateBiometric(ctx *gin.Context) {
	// UserBiometric := &models.UserBiometric{}
	// if err := ctx.BindJSON(&UserBiometric); err != nil {
	// 	utils.APIResponse(ctx, "Bad Request!", http.StatusBadRequest, http.MethodPost, nil)
	// } else {
	// 	data, _ := h.useCase.CreateBiometric(UserBiometric)
	// 	utils.APIResponse(ctx, "Connect biometric successfully.", http.StatusCreated, http.MethodPost, data)
	// }

	deviceId := ctx.Query("device_id")
	pubKey := ctx.Query("public_key")

	fmt.Printf("\nReceived device_id: %s, public_key: %s\n", deviceId, pubKey)

	if deviceId == "" && pubKey == "" {
		utils.APIResponse(ctx, "Bad Request", http.StatusBadRequest, http.MethodPost, nil)
	} else {
		UserBiometric := &models.UserBiometric{
			DEVICE_ID:    deviceId,
			PUBLIC_KEY:   pubKey,
			BIOMETRIC_ID: utils.RandRunes(30),
		}

		data, err := h.useCase.CreateBiometric(UserBiometric)
		if err != nil {
			utils.APIResponse(ctx, err.Error(), http.StatusBadRequest, http.MethodPost, nil)
		} else {
			utils.APIResponse(ctx, "Connect biometric successfully.", http.StatusCreated, http.MethodPost, data)
		}
	}

}
