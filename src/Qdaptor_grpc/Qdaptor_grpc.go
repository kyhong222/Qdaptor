package Qdaptor_grpc

import (
	"Qdaptor/api"
	"Qdaptor/logger"
	pb "Qdaptor/protos/Qdaptor_grpc"
	"context"
	"encoding/json"
	"net"
	"time"

	"go.uber.org/zap"

	"google.golang.org/grpc"
)

// var qdaptor_url string = "localhost"
var port string = "52051"

var QueueDN1 string = "8821"
var QueueDN2 string = "8822"

var AgentMap map[string](*api.Agent) = make(map[string](*api.Agent))

type Server struct {
	pb.UnimplementedTransactionServer
}

func (s *Server) HelloTransaction(ctx context.Context, msg *pb.TransactionMessage) (*pb.TransactionMessage, error) {
	// fmt.Println("received:", msg)
	logger.Info("Hello request is arrived",
		zap.Reflect("request", msg),
	)

	targetAgent := AgentMap[msg.CallId]

	// wait for response with for block
	for targetAgent.IVRResultResponse == nil {
		time.Sleep(1 * time.Second)
	}

	// fmt.Println(api.IVRResultResponse["ucid"].(string))
	ucid := targetAgent.IVRResultResponse["ucid"].(string)
	// fmt.Println(api.IVRResultResponse["extensiondata"].(string))
	IVRResult := targetAgent.IVRResultResponse["extensiondata"]
	b, _ := json.Marshal(IVRResult)

	extends := fusionObjectStrings(ucid, string(b))

	response := &pb.TransactionMessage{
		CallId:  msg.CallId,
		Message: extends,
	}

	logger.Info("Hello response is sent",
		zap.Reflect("response", response),
	)

	targetAgent.IVRResultResponse = nil
	return response, nil
}

func (s *Server) RefCallTransaction(ctx context.Context, msg *pb.TransactionMessage) (*pb.TransactionMessage, error) {
	logger.Info("RefCall request is arrived",
		zap.Reflect("request", msg),
	)

	targetAgent := AgentMap[msg.CallId]

	// call RefCall()
	targetAgent.RefCall(msg.CallId)

	// wait for response with for block
	for targetAgent.IVRResultResponse == nil {
		time.Sleep(1 * time.Second)
		// call HeartBeat
		// api.Heartbeat()
	}

	logger.Info("IVR response is arrived",
		zap.Reflect("IVR Response", targetAgent.IVRResultResponse),
	)

	ucid := targetAgent.IVRResultResponse["ucid"].(string)
	IVRResult := targetAgent.IVRResultResponse["extensiondata"].(string)

	extends := fusionObjectStrings(ucid, IVRResult)
	targetAgent.IVRResultResponse = nil

	response := &pb.TransactionMessage{
		CallId:  msg.CallId,
		Message: extends,
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
	targetAgent := AgentMap[msg.CallId]

	// call CallClear API
	targetAgent.CallClear(msg.Message)

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

func (s *Server) GetQueueTrafficTransaction(ctx context.Context, msg *pb.TransactionMessage) (*pb.TransactionMessage, error) {
	logger.Info("GetQueueTraffic request is arrived",
		zap.Reflect("request", msg),
	)

	targetAgent := AgentMap[msg.CallId]

	// call GetQueueTraffic for QueueDN1
	targetAgent.GetQueueTraffic(QueueDN1)

	// wait for response with for block
	for targetAgent.IVRResultResponse == nil {
		time.Sleep(1 * time.Second)
		// call HeartBeat
		// api.Heartbeat()
	}

	isReady := "false"
	if targetAgent.IVRResultResponse["readyagentcount"].(float64) != 0 {
		isReady = "true"
	} else {
		// set as nil
		targetAgent.IVRResultResponse = nil

		// call GetQueueTraffic for QueueDN2
		targetAgent.GetQueueTraffic(QueueDN2)

		// wait for response with for block
		for targetAgent.IVRResultResponse == nil {
			time.Sleep(1 * time.Second)
		}

		// wait for response with for block
		for targetAgent.IVRResultResponse == nil {
			time.Sleep(1 * time.Second)
		}

		if targetAgent.IVRResultResponse["readyagentcount"].(float64) != 0 {
			isReady = "true"
		}
	}
	targetAgent.IVRResultResponse = nil

	response := &pb.TransactionMessage{
		CallId:  msg.CallId,
		Message: isReady,
	}

	logger.Info("GetQueueTraffic response is sent",
		zap.Reflect("response", response),
	)

	return response, nil
}

func Init() {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		// fmt.Print("failed to listen:")
		// panic(err)
		logger.Fatal("failed to listen",
			zap.Error(err),
		)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTransactionServer(grpcServer, &Server{})
	// fmt.Println("server listening at", lis.Addr())
	logger.Info("server listening at " + port)

	if err := grpcServer.Serve(lis); err != nil {
		// fmt.Print("failed to serve:")
		// panic(err)
		logger.Fatal("failed to serve",
			zap.Error(err),
		)
	}
}

func fusionObjectStrings(objstring1 string, objstring2 string) (objstring3 string) {
	// var obj1_s = `{"ucid": "ucidValue"}`
	// var obj2_s = `{"extendsion":{"uei1": "a", "uei2": "b"}}`

	obj1 := make(map[string]interface{})
	obj2 := make(map[string]interface{})

	extendData := make(map[string]interface{})

	if err := json.Unmarshal([]byte(objstring2), &extendData); err != nil {
		panic(err)
	}
	// if err := json.Unmarshal([]byte(objstring1), &obj1); err != nil {
	// 	logger.Error("obj1 unmarshaling failed",
	// 		zap.Error(err),
	// 	)
	// }
	// if err := json.Unmarshal([]byte(objstring1), &obj2); err != nil {
	// 	logger.Error("obj1 unmarshaling failed",
	// 		zap.Error(err),
	// 	)
	// }

	obj1["ucid"] = objstring1
	obj2["extensiondata"] = extendData

	// fmt.Println(obj1, obj2)

	obj3 := make(map[string]interface{})
	for k, v := range obj1 {
		if _, ok := obj1[k]; ok {
			obj3[k] = v
		}
	}

	for k, v := range obj2 {
		if _, ok := obj2[k]; ok {
			obj3[k] = v
		}
	}

	obj3_b, _ := json.Marshal(obj3)
	objstring3 = string(obj3_b)
	return
}
