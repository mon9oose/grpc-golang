package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/saenggeuk/grpc-golang/client/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

var (
	port = flag.Int("p", 10091, "the grpc server port")
	name = flag.String("name", "ygpark", "name to greet")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(
		fmt.Sprintf("localhost:%d", *port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to conn: %v", err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatalf("filed to close conn: %v", err)
		}
	}()
	c := pb.NewGreetingServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := c.SayHello(ctx, &pb.GreetingRequest{Name: *name})
	if err != nil {
		log.Fatalf("failed to grpc comm: %v", err)
	}
	log.Printf("[%d] Greeting : %s", res.GetStatusCode(), res.GetMessage())
}
