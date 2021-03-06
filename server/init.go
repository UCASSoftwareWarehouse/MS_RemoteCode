package server

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"remote_code/config"
	"remote_code/consul"
	"remote_code/pb_gen"
)

func StartServer() {
	appName := config.Conf.AppName
	println(appName)
	addr := fmt.Sprintf("%s:%d", config.Conf.Host, config.Conf.Port)
	log.Printf("%s Dialing addr: %s", appName, addr)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("failed to listen: %v", err)
	}
	var options []grpc.ServerOption
	options = append(options, grpc.MaxSendMsgSize(5*1024*1024*1024*1024), grpc.MaxRecvMsgSize(5*1024*1024*1024*1024))
	grpcServer := grpc.NewServer(options...)
	consul.MustRegisterGRPCServer(grpcServer)
	pb_gen.RegisterRemoteCodeServiceServer(grpcServer, newRemoteCodeServer())
	log.Printf("%s ready to server at %s...", appName, addr)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Printf("grpcServer Serve failed, err=[%v]", err)
	}
}
