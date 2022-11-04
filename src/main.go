package main

import (
	Q "Qdaptor/Qdaptor_grpc"
	"Qdaptor/api"
)

const baseURL string = "https://dev-icweb.ssgadm.com:9203/ic"

const g_DN string = "5205"
const g_tenant string = "SSG_DEV"
const g_agentID string = "test04"
const g_appName string = "SSGVoicebot"

const HBPeriod int = 10 // heartbeat period
const HBErrCnt int = 18 // heartbeat error count

func main() {
	go func() {
		Q.Init()
	}()
	// api.Start(baseURL, HBPeriod, HBErrCnt, g_appName, g_DN, g_tenant, g_agentID)
	for {
		if api.ErrorCount >= HBErrCnt {
			api.ErrorCount = 0
			api.Start(baseURL, HBPeriod, HBErrCnt, g_appName, g_DN, g_tenant, g_agentID)
		}
	}
}
