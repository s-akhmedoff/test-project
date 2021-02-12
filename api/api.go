package api

import (
	"google.golang.org/grpc"
	"test-project/api/genproto"
	v1 "test-project/api/handlers/v1"
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

	handler := v1.New(service)

	router.POST("/send", handler.Send)

	return router
}
