package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/saenggeuk/grpc-golang/chat"
	"google.golang.org/grpc"
	"log"
	"net"
)

var port = flag.Int("p", 10091, "the grpc server port")

type Handler struct {
	pb.UnimplementedGreetingServiceServer
}

func (h *Handler) SayHello(ctx context.Context, in *pb.GreetingRequest) (*pb.GreetingResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.GreetingResponse{
		StatusCode: 200,
		Message:    fmt.Sprintf("Hello, %s", in.GetName()),
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreetingServiceServer(s, &Handler{})
	log.Printf("server listening at %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
