package consul

import (
	"context"
	"log"
	"remote_code/config"
	"remote_code/pb_gen"
	//pb_gen2 "MS_Local/pb_gen"
	"testing"
)

func TestGetConsulHost(t *testing.T) {
	client := &GrpcClient{Name: "remote_code"}
	config.InitConfigDefault()
	client.RunConsulClient()
	remote_code := pb_gen.NewRemoteCodeServiceClient(client.Conn)
	request := &pb_gen.HelloWorldRequest{
		HelloText: "wanna",
	}
	resp := &pb_gen.HelloWorldResponse{}
	resp, err := remote_code.HelloWorld(context.Background(), request)
	log.Println(resp)
	log.Println(err)

}
