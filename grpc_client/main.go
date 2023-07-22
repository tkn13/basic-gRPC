package main

import (
	"context"
	"log"
	"time"

	"github.com/tkn13/grpc/common"
	pb "github.com/tkn13/grpc/student"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost"+common.SERVER_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewStudentClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetStudentById(ctx, &pb.StudentRequest{Id: 1})
	if err != nil {
		log.Fatalf("could not send request: %v", err)
	}
	log.Printf("data: %v\n", r.String())
	log.Printf("ID: %v\n", r.GetId())
	log.Printf("Name: %v\n", r.GetName())
	log.Printf("Age: %v\n", r.GetAge())
}
