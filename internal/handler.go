package internal

import (
	"context"
	"remote_code/internal/test"
	"remote_code/pb_gen"
)

func HelloWorld(ctx context.Context, req *pb_gen.HelloWorldRequest) (*pb_gen.HelloWorldResponse, error) {
	return test.HelloWorld(ctx,req)
}
