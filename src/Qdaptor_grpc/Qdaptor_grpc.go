package Qdaptor_grpc

import (
	pb "Qdaptor/protos/Qdaptor_grpc"
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

// var qdaptor_url string = "localhost"
var port string = "52051"

type Server struct {
	pb.UnimplementedTransactionServer
}

func (s *Server) HelloTransaction(ctx context.Context, msg *pb.TransactionMessage) (*pb.TransactionMessage, error) {
	fmt.Println("received:", msg)

	response := &pb.TransactionMessage{
		CallId:  "5205",
		Message: "test hello",
	}

	return response, nil
}

func (s *Server) RefCallTransaction(ctx context.Context, msg *pb.TransactionMessage) (*pb.TransactionMessage, error) {
	fmt.Println("received:", msg)

	response := &pb.TransactionMessage{
		CallId:  "5205",
		Message: "test refCallTransaction",
	}

	return response, nil
}

func Init() {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Print("failed to listen:")
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTransactionServer(grpcServer, &Server{})
	fmt.Println("server listening at", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		fmt.Print("failed to serve:")
		panic(err)
	}
}
