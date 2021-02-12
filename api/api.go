package api

import (
	"google.golang.org/grpc"
	"log"
	_ "test-project/api/docs"
	"test-project/api/genproto"
	"test-project/utils"

	"github.com/gin-gonic/gin"
)

// New ...
// @title test-project
// @version 1.0
// @description ...
func New() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger(), gin.Recovery())

	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	utils.FailOnError(err, "Failed to connect grpc")
	service := genproto.NewGeneratorServiceClient(conn)
	router.POST("/send", func(c *gin.Context){
		in := &genproto.GSRequest{}
		err := c.ShouldBindJSON(in)
		log.Print(in)
		utils.FailOnError(err, "Failed to bind")
		_, err = service.GenerateAndSend(c, in)
		utils.FailOnError(err, "Failed to execute")
		return
	})

	return router
}
