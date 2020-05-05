package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"laracom/demo-service/api"
	pb "laracom/demo-service/proto/demo"
	"log"
	"net"
)

const (
	address = "localhost:9091"
	grpcPort = ":9999"
	httpPort = "8000"
	appName = "Demo Service"
)

type DemoService struct {

}

func (ds *DemoService) SayHello(ctx context.Context, req *pb.DemoRequest) (*pb.DemoResponse, error) {
	rsp := &pb.DemoResponse{Text: "你好, " + req.Name}
	return rsp, nil
}

func main() {
	// 通过-mode参数判断启动哪个模式的代码
	mode := flag.String("mode", "grpc", "mode:grpc/http/client")
	flag.Parse()
	fmt.Println("run mode:", *mode)
	switch *mode {
	case "http":
		fmt.Printf("Starting %v\n", appName)
		api.StartWebServer(httpPort)
	case "client":
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("连接到 gRPC 服务器失败: %v", err)
		}
		defer conn.Close()
		client := pb.NewDemoServiceClient(conn)
		req := &pb.DemoRequest{Name: "学院君"}
		rsp, err := client.SayHello(context.Background(), req)
		if err != nil {
			log.Fatalf("调用 gRPC 服务接口失败: %v", err)
		}
		log.Printf("%s", rsp.Text)
	case "grpc":
		fallthrough
	default:
		listener, err := net.Listen("tcp", grpcPort)
		if err != nil {
			log.Fatalf("监听指定端口失败: %v", err)
		}

		server := grpc.NewServer()
		pb.RegisterDemoServiceServer(server, &DemoService{})

		reflection.Register(server)

		if err := server.Serve(listener); err != nil {
			log.Fatalf("服务启动失败: %v", err)
		}
	}
}
