package server

import (
	"context"
	"remote_code/internal/test"
	"remote_code/pb_gen"
)

type remoteCodeServer struct {
	*pb_gen.UnimplementedRemoteCodeServiceServer
}

func newRemoteCodeServer() *remoteCodeServer {
	return &remoteCodeServer{}
}

func (c *remoteCodeServer) HelloWorld(ctx context.Context, request *pb_gen.HelloWorldRequest) (*pb_gen.HelloWorldResponse, error) {
	return test.HelloWorld(ctx,request)
	//return &pb_gen.HelloWorldResponse{ThanksText: request.GetHelloText() + ", thx."}, nil
}