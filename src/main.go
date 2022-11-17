package main

import (
	"Qdaptor/Qdaptor_grpc"
	"Qdaptor/config"
	"sync"
)

// const baseURL string = "https://dev-icweb.ssgadm.com:9203/ic"

// const g_DN string = "5205"
// const g_tenant string = "SSG_DEV"
// const g_agentID string = "test04"
// const g_appName string = "SSGVoicebot"

// const HBPeriod int = 10 // heartbeat period
// const HBErrCnt int = 18 // heartbeat error count

var wg sync.WaitGroup

func main() {
	defer wg.Wait()
	go Qdaptor_grpc.Init()

	config.Init()
	for _, agent := range Qdaptor_grpc.AgentMap {
		go agent.Start()
	}
	wg.Add(1)
}
