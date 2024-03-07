package biometric

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"pthw.com/asymmetric-for-biometric/config"
	"pthw.com/asymmetric-for-biometric/internal/biometric"
	"pthw.com/asymmetric-for-biometric/models"
	"pthw.com/asymmetric-for-biometric/utils"
	"pthw.com/asymmetric-for-biometric/utils/jwt"
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

	UserBiometric := &models.UserBiometric{}
	if err := ctx.BindJSON(&UserBiometric); err != nil {
		utils.ApiErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, err.Error())
	} else {
		UserBiometric.BIOMETRIC_ID = utils.RandRunes(30)
		data, err := h.useCase.CreateBiometric(UserBiometric)

		token, _ := jwt.CreateAccessToken(data, config.Secret, config.Expire)
		response := map[string]interface{}{"token": token, "data": data}

		if err != nil {
			utils.ApiErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, err.Error())
		} else {
			utils.APIResponse(ctx, "Connect biometric successfully.", http.StatusCreated, http.MethodPost, response)
		}
	}

	fmt.Printf("\nReceived device_id: %s, public_key: %s\n", UserBiometric.DEVICE_ID, UserBiometric.PUBLIC_KEY)

}

func (h *Handler) GetChallenge(ctx *gin.Context) {
	deviceId := ctx.Param("device_id")

	if deviceId == "" {
		utils.ApiErrorResponse(ctx, http.StatusBadRequest, http.MethodGet, "Bad Request")
	} else {
		data, err := h.useCase.GetChallenge(deviceId)
		if err != nil {
			utils.ApiErrorResponse(ctx, http.StatusBadRequest, http.MethodGet, err.Error())
		} else {
			utils.APIResponse(ctx, "Challenge retrieved", http.StatusOK, http.MethodGet, data)
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
		utils.ApiErrorResponse(ctx, http.StatusBadRequest, http.MethodGet, err.Error())
	} else {
		data, err := h.useCase.ValidateBiometric(data.BIOMETRIC_ID, data.SIGNATURE)
		if err != nil {
			utils.ApiErrorResponse(ctx, http.StatusBadRequest, http.MethodGet, err.Error())
		} else {
			utils.APIResponse(ctx, data, http.StatusOK, http.MethodGet, data)
		}
	}
}
