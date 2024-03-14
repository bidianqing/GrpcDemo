package main

import (
	"context"
	"fmt"
	"golangapp/client/Protos/greet"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/*
https://github.com/protocolbuffers/protobuf

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative .\Protos\greet\greet.proto
*/

func main() {
	target := "localhost:7075"
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	greetClient := greet.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	reply, err := greetClient.SayHello(ctx, &greet.HelloRequest{
		Name: "golang",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(reply.GetMessage())
}
