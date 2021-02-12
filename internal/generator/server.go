package main

import (
	"google.golang.org/grpc"
	"net"
	"test-project/internal/generator/genproto"
	"test-project/internal/generator/handler"
	"test-project/utils"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	utils.FailOnError(err, "Failed to listen port")

	s := handler.Server{}

	grpcServer := grpc.NewServer()

	genproto.RegisterGeneratorServiceServer(grpcServer, &s)

	err = grpcServer.Serve(lis)
	utils.FailOnError(err, "Failed to serve")
}
