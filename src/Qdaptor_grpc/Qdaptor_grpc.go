package Qdaptor_grpc

import (
	"Qdaptor/api"
	"Qdaptor/logger"
	pb "Qdaptor/protos/Qdaptor_grpc"
	"context"
	"fmt"
	"net"

	"go.uber.org/zap"

	"google.golang.org/grpc"
)

// var qdaptor_url string = "localhost"
var port string = "52051"

type Server struct {
	pb.UnimplementedTransactionServer
}

func (s *Server) HelloTransaction(ctx context.Context, msg *pb.TransactionMessage) (*pb.TransactionMessage, error) {
	// fmt.Println("received:", msg)
	logger.Info("Hello request is arrived",
		zap.Reflect("request", msg),
	)

	api.APIWaitGroup.Add(1)
	api.APIWaitGroup.Wait()

	response := &pb.TransactionMessage{
		CallId:  msg.CallId,
		Message: api.APIVars.UCID,
	}

	logger.Info("Hello response is sent",
		zap.Reflect("response", response),
	)

	return response, nil
}

func (s *Server) RefCallTransaction(ctx context.Context, msg *pb.TransactionMessage) (*pb.TransactionMessage, error) {
	logger.Info("RefCall request is arrived",
		zap.Reflect("request", msg),
	)

	// call RefCall()
	api.RefCall(msg.CallId)

	// call HeartBeat
	api.Heartbeat()

	api.APIWaitGroup.Add(1)
	api.APIWaitGroup.Wait()

	logger.Info("IVR response is arrived",
		zap.Reflect("IVR Response", api.IVRResultResponse),
	)

	IVRResult := api.IVRResultResponse["extensiondata"].(string)
	api.IVRResultResponse = nil

	response := &pb.TransactionMessage{
		CallId:  msg.CallId,
		Message: IVRResult,
	}

	logger.Info("RefCall response is sent",
		zap.Reflect("response", response),
	)

	return response, nil
}

func (s *Server) CallClearTransaction(ctx context.Context, msg *pb.TransactionMessage) (*pb.TransactionMessage, error) {
	logger.Info("CallClear request is arrived",
		zap.Reflect("request", msg),
	)

	// msg.Message에 종료 타입이 넘어올 예정

	// TODO: call API: isAbleToTransfer

	// call CallClear API
	api.CallClear(msg.Message)

	// TODO: make callClear UEIs and response it first,
	response := &pb.TransactionMessage{
		CallId:  msg.CallId,
		Message: "callClear api is called", // 호 종료 관련 UEI가 들어갈 수도 있음. 관련 ExtensionData를 api쪽에 구현할지, 여기구현할지 미정
	}

	logger.Info("CallClear response is sent",
		zap.Reflect("response", response),
	)

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
