package main

import (
	"context"
	"fmt"
	"flag"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	v1 "github.com/laughtt/templatesGrpc/service+sql/api/protov1"
	"google.golang.org/grpc"
)

func main() {
	// get configuration
	address := flag.String("server", "", "gRPC server in format host:port")
	flag.Parse()
	fmt.Println(*address)
	// Set up a connection to the server.
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := v1.NewAuthServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	t := time.Now().In(time.UTC)
	reminder, _ := ptypes.TimestampProto(t)
	//pfx := t.Format(time.RFC3339Nano)

	req1 := &v1.CreateRequest{
		Id: 64,
		Message: "bitch",
		Reminder: reminder,
	}
	res1 , err := c.SendMessage(ctx , req1)
	fmt.Println(res1)
}
