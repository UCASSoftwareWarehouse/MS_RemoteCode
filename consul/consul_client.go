package consul

import (
	"fmt"
	"google.golang.org/grpc/balancer/roundrobin"
	"remote_code/config"

	//_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"google.golang.org/grpc"
)

func GetConsulHost() string {
	return config.Conf.ConsulAddr
}

type GrpcClient struct {
	Conn      *grpc.ClientConn
	RpcTarget string
	Name      string
}

func (s *GrpcClient) RunGrpcClient() {
	conn, err := grpc.Dial(s.RpcTarget, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	s.Conn = conn
	fmt.Println("grpc client start success")
}
func (s *GrpcClient) RunConsulClient() {
	//初始化 resolver 实例
	Init()
	conn, err := grpc.Dial(
		fmt.Sprintf("%s://%s/%s", "consul", GetConsulHost(), s.Name),
		//不能block => blockkingPicker打开，在调用轮询时picker_wrapper => picker时若block则不进行robin操作直接返回失败
		//grpc.WithBlock(),
		grpc.WithInsecure(),
		//指定初始化round_robin => balancer (后续可以自行定制balancer和 register、resolver 同样的方式)
		grpc.WithBalancerName(roundrobin.Name),
		//grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
	)
	if err != nil {
		fmt.Println("dial err:", err)
		return
	}
	s.Conn = conn
	fmt.Println(fmt.Sprintf("gRpc consul client [%s] start success", s.Name))
}
