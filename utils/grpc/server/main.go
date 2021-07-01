package main

import (
	"context"
	"goWebDemo/utils"
	"log"
	"net"

	pb "goWebDemo/utils/grpc/proto"
	"google.golang.org/grpc"
)



// server is used to implement message.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements message.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", utils.RpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}