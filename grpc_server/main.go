package main

import (
	"context"
	"log"
	"net"

	"github.com/tkn13/grpc/common"
	pb "github.com/tkn13/grpc/student"
	"google.golang.org/grpc"
)

type StudentService struct{
	pb.UnimplementedStudentServer
}

func (s StudentService)GetStudentById(ctx context.Context,
	 in *pb.StudentRequest) (*pb.StudentResponse, error){
	log.Printf("Received: %v", in.GetId())
	return &pb.StudentResponse{
		Id: 1,
		Name: "Mameow",
		Age: 21,
	},nil
}

func main() {
	lis, err := net.Listen("tcp", common.SERVER_PORT);
	if err != nil{
		log.Fatalf("[main] cannot listen to port: %s %v", common.SERVER_PORT, err)
	}
	server := grpc.NewServer()
	pb.RegisterStudentServer(server, &StudentService{})
	log.Printf("Server was started on port: %s",common.SERVER_PORT)
	if err := server.Serve(lis); err != nil{
		log.Fatalf("failed to serve: %v", err)
	}
}