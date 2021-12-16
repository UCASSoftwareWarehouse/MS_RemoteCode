package test

import (
	"context"
	"log"
	"remote_code/pb_gen"
)

func HelloWorld(ctx context.Context, req *pb_gen.HelloWorldRequest) (*pb_gen.HelloWorldResponse, error) {
	log.Println(req)
	resp := &pb_gen.HelloWorldResponse{
		ThanksText: "hello this is wanna",
	}
	return resp, nil
}
