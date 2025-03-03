package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/tokenoff03/authentication-service/pkg/access_v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

var accessToken = flag.String("a", "", "access token")

const servicePort = 50059

func main() {
	flag.Parse()

	ctx := context.Background()
	md := metadata.New(map[string]string{"Authorization": "Bearer " + *accessToken})
	ctx = metadata.NewOutgoingContext(ctx, md)

	conn, err := grpc.NewClient(
		fmt.Sprintf(":%d", servicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to dial GRPC client: %v", err)
	}

	cl := access_v1.NewAccessV1Client(conn)

	_, err = cl.Check(ctx, &access_v1.CheckRequest{
		EndpointAddress: "/note_v1.NoteV1/Get",
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Access granted")
}
