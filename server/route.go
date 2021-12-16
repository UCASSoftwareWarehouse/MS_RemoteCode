package server

import (
	"context"
	"log"
	"remote_code/internal/remote"
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
	log.Println("hello here")
	return test.HelloWorld(ctx, request)
}

func (c *remoteCodeServer) DownloadRemoteCode(ctx context.Context, request *pb_gen.DownloadRemoteCodeRequest) (*pb_gen.DownloadRemoteCodeResponse, error) {
	log.Println("pypi here")
	return remote.DownloadRemoteCode(ctx, request)
}

func (c *remoteCodeServer) DownloadAptDeb(ctx context.Context, request *pb_gen.DownloadAptDebRequest) (*pb_gen.DownloadAptDebResponse, error) {
	log.Println("apt here")
	return remote.DownloadAptDeb(ctx, request)
}
