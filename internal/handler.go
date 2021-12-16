package internal

import (
	"context"
	"log"
	"remote_code/internal/remote"
	"remote_code/internal/test"
	"remote_code/pb_gen"
)

func HelloWorld(ctx context.Context, req *pb_gen.HelloWorldRequest) (*pb_gen.HelloWorldResponse, error) {
	return test.HelloWorld(ctx, req)
}

func DownloadRemoteCode(ctx context.Context, req *pb_gen.DownloadRemoteCodeRequest) (*pb_gen.DownloadRemoteCodeResponse, error) {
	log.Println("pypi here")
	return remote.DownloadRemoteCode(ctx, req)
}
