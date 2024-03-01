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

	UserBiometric := &models.UserBiometric{}
	if err := ctx.BindJSON(&UserBiometric); err != nil {
		utils.APIResponse(ctx, "Bad Request", http.StatusBadRequest, http.MethodPost, nil)
	} else {
		// UserBiometric := &models.UserBiometric{
		// 	DEVICE_ID:    deviceId,
		// 	PUBLIC_KEY:   pubKey,
		// 	BIOMETRIC_ID: utils.RandRunes(30),
		// }

		UserBiometric.BIOMETRIC_ID = utils.RandRunes(30)
		data, err := h.useCase.CreateBiometric(UserBiometric)
		if err != nil {
			utils.APIResponse(ctx, err.Error(), http.StatusBadRequest, http.MethodPost, nil)
		} else {
			utils.APIResponse(ctx, "Connect biometric successfully.", http.StatusCreated, http.MethodPost, data)
		}
	}

	fmt.Printf("\nReceived device_id: %s, public_key: %s\n", UserBiometric.DEVICE_ID, UserBiometric.PUBLIC_KEY)

}

func (h *Handler) GetChallenge(ctx *gin.Context) {
	deviceId := ctx.Param("device_id")

	if deviceId == "" {
		utils.APIResponse(ctx, "Bad Request", http.StatusBadRequest, http.MethodGet, nil)
	} else {
		data, err := h.useCase.GetChallenge(deviceId)
		if err != nil {
			utils.APIResponse(ctx, err.Error(), http.StatusBadRequest, http.MethodGet, nil)
		} else {
			utils.APIResponse(ctx, "Challenge retrieved", http.StatusCreated, http.MethodGet, data)
		}
	}
}

type ValidateBiometric struct {
	BIOMETRIC_ID string `json:"biometric_id"`
	SIGNATURE    string `json:"signature"`
}

func (h *Handler) ValidateBiometric(ctx *gin.Context) {
	// biometricId := ctx.Query("biometric_id")
	// signature := ctx.Query("signature")

	data := &ValidateBiometric{}
	if err := ctx.BindJSON(&data); err != nil {
		utils.APIResponse(ctx, "Bad Request", http.StatusBadRequest, http.MethodGet, nil)
	} else {
		data, err := h.useCase.ValidateBiometric(data.BIOMETRIC_ID, data.SIGNATURE)
		if err != nil {
			utils.APIResponse(ctx, err.Error(), http.StatusBadRequest, http.MethodGet, nil)
		} else {
			utils.APIResponse(ctx, data, http.StatusCreated, http.MethodGet, data)
		}
	}
}
