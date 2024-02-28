package biometric

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"pthw.com/asymmetric-for-biometric/internal/biometric"
)

type Bookmark struct {
	ID    string `json:"id"`
	URL   string `json:"url"`
	Title string `json:"title"`
}

type Handler struct {
	useCase biometric.UseCase
}

func NewHandler(useCase biometric.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) Create(context *gin.Context) {
	context.Status(http.StatusOK)
}
