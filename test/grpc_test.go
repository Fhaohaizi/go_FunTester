package test

import (
	"context"
	"fmt"
	"funtester/ftool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"testing"
	"time"
	// 导入刚才我们生成的代码所在的proto包。
	pb "funtester/proto"
)

const address = "127.0.0.1:12345"

// TestGrpcClient
// @Description: gRPC客户端Demo
// @param t
func TestGrpcClient(t *testing.T) {
	conn, _ := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	defer conn.Close()
	c := pb.NewHelloServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, _ := c.ExecuteHi(ctx, &pb.HelloRequest{
		Name: "FunTester",
		Age:  23,
	})
	fmt.Printf("msg: %s\n", r.Msg)

}
func TestGrpcClient2(t *testing.T) {
	conn, _ := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	defer conn.Close()
	c := pb.NewHelloServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, _ := c.ExecuteHi(ctx, &pb.HelloRequest{
		Name: "FunTester",
		Age:  23,
	})
	fmt.Printf("msg: %s\n", r.Msg)

}

type Ser struct {
}

func (s *Ser) ExecuteHi(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	// 创建一个HelloReply消息，设置Message字段，然后直接返回。
	return &pb.HelloResponse{Msg: "Hello " + in.Name + ftool.Date()}, nil
}

func TestGrpcService(t *testing.T) {
	// 监听127.0.0.1:12345
	lis, err := net.Listen("tcp", "127.0.0.1:12345")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 实例化grpc服务端
	s := grpc.NewServer()

	pb.RegisterHelloServiceServer(s, &Ser{})

	// 往grpc服务端注册反射服务
	reflection.Register(s)

	// 启动grpc服务
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type StreamDemo struct {
}

func (s *StreamDemo) ExecuteHi(req *pb.DemoRequest, ser pb.StreamDemo_ExecuteHiServer) {

}
