package v1

import (
	"test-project/api/genproto"
	"test-project/api/models"

	"github.com/gin-gonic/gin"
)

// Handler ...
type Handler struct {
	client genproto.GeneratorServiceClient
}

// New ...
func New(c genproto.GeneratorServiceClient) *Handler {
	return &Handler{
		client: c,
	}
}

func (h *Handler) handleSuccessResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, models.SuccessResponse{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func (h *Handler) handleErrorResponse(c *gin.Context, code int, message string, err interface{}) {
	c.JSON(code, models.ErrorResponse{
		Code:    code,
		Message: message,
		Error:   err,
	})
}