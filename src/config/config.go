package config

import (
	"Qdaptor/Qdaptor_grpc"
	"Qdaptor/api"
	"Qdaptor/logger"
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"

	"go.uber.org/zap"
)

func Init() {
	var configFilePath string = "../config.json"
	jsonFile, err := os.Open(configFilePath)
	// if we os.Open returns an error then handle it
	if err != nil {
		println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	agentArray := make(map[string]interface{})
	err = json.Unmarshal([]byte(byteValue), &agentArray)

	if err != nil {
		println(err)
	}
	// fmt.Println(reflect.TypeOf(agentArray["agents"]))
	for _, agent := range agentArray["agents"].([]interface{}) {
		// fmt.Println(reflect.TypeOf(agent))
		var err error

		agent_map := agent.(map[string]interface{})
		agent_struct := api.Agent{}
		agent_struct.DN = agent_map["DN"].(string)
		agent_struct.Tenant = agent_map["Tenant"].(string)
		agent_struct.AgentID = agent_map["AgentID"].(string)
		agent_struct.APIVars.HBPeriod, err = strconv.Atoi(agent_map["HBPeriod"].(string))
		if err != nil {
			logger.Fatal("HBPeriod in config.json is not string",
				zap.Error(err))
		}
		agent_struct.APIVars.HBErrCnt, err = strconv.Atoi(agent_map["HBErrCnt"].(string))
		if err != nil {
			logger.Fatal("HBPeriod in config.json is not string",
				zap.Error(err))
		}
		Qdaptor_grpc.AgentMap[agent_struct.DN] = &agent_struct

		// fmt.Println(Qdaptor_grpc.AgentMap[agent_struct.DN])
	}
}
