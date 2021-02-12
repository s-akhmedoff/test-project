package v1

import (
	"github.com/gin-gonic/gin"
	"test-project/api/models"
)

// Ping godoc
// @ID ping
// @Router /ping [GET]
// @Summary returns "pong" message
// @Description this returns "pong" message to show service is working
// @Accept json
// @Produce json
// @Success 200 {object} models.SuccessResponse{data=string} "desc"
// @Failure 500 {object} models.ErrorResponse{error=string}
func (h *Handler) Ping(c *gin.Context) {
	c.JSON(200, models.SuccessResponse{
		Code:    200,
		Message: "ok",
		Data:    "pong",
	})
	return
}