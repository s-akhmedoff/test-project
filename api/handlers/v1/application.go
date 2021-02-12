package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	"test-project/api/genproto"
	"test-project/utils"
)

func (h *Handler) Send(c *gin.Context) {
	in := &genproto.GSRequest{}
	err := c.ShouldBindJSON(in)
	log.Print(in)
	utils.FailOnError(err, "Failed to bind")
	_, err = h.client.GenerateAndSend(c,in)
	utils.FailOnError(err, "Failed to execute")
	return
}