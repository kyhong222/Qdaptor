package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/google/go-querystring/query"
)

const baseURL string = "https://dev-icweb.ssgadm.com:9203/ic"

var session string
var handle int

const g_DN string = "5205"
const g_tenant string = "SSG_DEV"
const g_agentID string = "test04"
const g_appName string = "SSGVoicebot"

const HBPeriod int = 10 // heartbeat period
const HBErrCnt int = 18 // heartbeat error count

type OpenServerMsg struct {
	MessageType int    `json:"messagetype"`
	Method      string `json:"method"`
	Handle      int    `json:"handle"`
	Result      string `json:"result"`
	Key         string `json:"key"`
}

type RegisterMsg struct {
	MessageType int    `json:"messagetype"`
	Method      string `json:"method"`
	Result      string `json:"result"`
}

type LoginMsg struct {
	MessageType int    `json:"messagetype"`
	Method      string `json:"method"`
	Result      string `json:"result"`
}

type SetAgentStateMsg struct {
	MessageType int    `json:"messagetype"`
	Method      string `json:"method"`
	Result      string `json:"result"`
}

type OpenServerQueryOption struct {
	AppName string `url:"appname"`
}

type RegisterQueryOption struct {
	Key    string `url:"key"`
	Handle int    `url:"handle"`
	DN     string `url:"thisdn"`
	Tenant string `url:"tenantname"`
}

type LoginQueryOption struct {
	Key           string `url:"key"`
	Handle        int    `url:"handle"`
	Tenant        string `url:"tenantname"`
	AgentDN       string `url:"agentdn"`
	AgentID       string `url:"agentid"`
	AgentPassword string `url:"agentpassword"`
	AgentState    string `url:"agentstate"`
	AgentStateSub string `url:"agentstatesub"`
	PasswordType  string `url:"passwdtype"`
	MediaSet      string `url:"mediaset"`
}

type SetAgentStateQueryOption struct {
	Key           string `url:"key"`
	Handle        int    `url:"handle"`
	Tenant        string `url:"tenantname"`
	AgentID       string `url:"agentid"`
	AgentState    string `url:"agentstate"`
	AgentStateSub string `url:"agentstatesub"`
	MediaSet      string `url:"mediaset"`
}
type SetAfterCallReadyQueryOption struct {
	Key           string `url:"key"`
	Handle        int    `url:"handle"`
	Tenant        string `url:"tenantname"`
	AgentID       string `url:"agentid"`
	AgentState    string `url:"agentstate"`
	AgentStateSub string `url:"agentstatesub"`
	MediaSet      string `url:"mediaset"`
}

type HeartbeatQueryOption struct {
	Key string `url:"key"`
}

func main() {
	openServer(g_appName)
	heartbeat()
	register(g_DN)
	heartbeat()
	heartbeat()
	login(g_agentID, g_DN, g_tenant)
	heartbeat()
	heartbeat()
	// setReady(g_tenant, g_agentID)
	setAfterCallReady(g_tenant, g_agentID)
	heartbeat()
	heartbeatMaker(HBPeriod)

}

func openServer(AppName string) {
	option := OpenServerQueryOption{
		AppName,
	}

	v, _ := query.Values(option)

	url := baseURL + "/openserver?" + v.Encode()

	// openServer 호출
	resp, err := http.Get(url)
	if err != nil {
		// panic(err)
		fmt.Println("Retry>>\tSession disconnected, retry openServer() to connect")
		register(AppName)
		return
	}

	defer resp.Body.Close()

	// 결과 데이터를 res에 저장
	resJson := OpenServerMsg{}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// panic(err)
		fmt.Println("Retry>>\tSession disconnected, retry openServer() to connect")
		register(AppName)
		return
	}

	err = json.Unmarshal(data, &resJson)
	if err != nil {
		// panic(err)
		fmt.Println("Retry>>\tSession disconnected, retry openServer() to connect")
		register(AppName)
		return
	}
	fmt.Println("openServer\t", string(data))

	// session 및 handle 값 저장
	session = resJson.Key
	handle = resJson.Handle

}

func register(DN string) {
	tenantName := "SSG_DEV"
	option := RegisterQueryOption{
		session, handle, DN, tenantName,
	}

	v, _ := query.Values(option)

	url := baseURL + "/register?" + v.Encode()

	// register 호출
	reqBody := bytes.NewBufferString("") // body가 필요없으나, 파라미터라 선언.
	resp, err := http.Post(url, "application/json", reqBody)
	if err != nil {
		// panic(err)
		fmt.Println("Retry>>\tSession disconnected, retry register() to connect")
		register(DN)
		return
	}

	defer resp.Body.Close()

	// 결과 데이터를 res에 저장
	resJson := RegisterMsg{}
	data, err := ioutil.ReadAll(resp.Body) // data는 byte[]
	if err != nil {
		// panic(err)
		fmt.Println("Retry>>\tSession disconnected, retry register() to connect")
		register(DN)
		return
	}

	err = json.Unmarshal(data, &resJson)
	if err != nil {
		// panic(err)
		fmt.Println("Retry>>\tSession disconnected, retry register() to connect")
		register(DN)
		return
	}
	fmt.Println("register>>\t", string(data))
}

func login(agnetID string, DN string, tenant string) {

	option := LoginQueryOption{
		session,
		handle,
		tenant,
		DN,
		agnetID,
		"",
		"40", // 40 is ready
		"0",  // state sub is 0
		"4",  // 4 is SHA-2(512)
		"",   // mediaset is blank
	}

	v, _ := query.Values(option)

	url := baseURL + "/agentlogin?" + v.Encode()

	// login 호출
	reqBody := bytes.NewBufferString("") // body가 필요없으나, 파라미터라 선언.
	resp, err := http.Post(url, "application/json", reqBody)
	if err != nil {
		// panic(err)
		fmt.Println("Retry>>\tSession disconnected, retry login() to connect")
		login(agnetID, DN, tenant)
		return
	}

	defer resp.Body.Close()

	// 결과 데이터를 res에 저장
	resJson := RegisterMsg{}
	data, err := ioutil.ReadAll(resp.Body) // data는 bytep[]]
	if err != nil {
		// panic(err)
		fmt.Println("Retry>>\tSession disconnected, retry login() to connect")
		login(agnetID, DN, tenant)
		return
	}

	err = json.Unmarshal(data, &resJson)
	if err != nil {
		// panic(err)
		fmt.Println("Retry>>\tSession disconnected, retry login() to connect")
		login(agnetID, DN, tenant)
		return
	}

	fmt.Println("login>>\t", string(data))
}

func setReady(tenant string, agentID string) {
	option := SetAgentStateQueryOption{
		session,
		handle,
		tenant,
		agentID,
		"40", // 40 is ready
		"0",  // state sub is 0
		"",   // mediaset is blank
	}

	v, _ := query.Values(option)

	url := baseURL + "/setagentstate?" + v.Encode()

	// login 호출
	reqBody := bytes.NewBufferString("") // body가 필요없으나, 파라미터라 선언.
	resp, err := http.Post(url, "application/json", reqBody)
	if err != nil {
		// panic(err)
		fmt.Println("Retry>>\tSession disconnected, retry setReady() to connect")
		setReady(tenant, agentID)
		return
	}

	defer resp.Body.Close()

	// 결과 데이터를 res에 저장
	resJson := RegisterMsg{}
	data, err := ioutil.ReadAll(resp.Body) // data는 bytep[]
	if err != nil {
		// panic(err)
		fmt.Println("Retry>>\tSession disconnected, retry setReady() to connect")
		setReady(tenant, agentID)
		return
	}

	err = json.Unmarshal(data, &resJson)
	if err != nil {
		// panic(err)
		fmt.Println("Retry>>\tSession disconnected, retry setReady() to connect")
		setReady(tenant, agentID)
		return
	}

	fmt.Println("setReady>>\t", string(data))

	// fmt.Println(url)
}

func setAfterCallReady(tenant string, agentID string) {
	option := SetAfterCallReadyQueryOption{
		session,
		handle,
		tenant,
		agentID,
		"40", // 40 is ready
		"0",  // state sub is 0
		"",   // mediaset is blank
	}

	v, _ := query.Values(option)

	url := baseURL + "/setaftcallstate?" + v.Encode()

	// login 호출
	reqBody := bytes.NewBufferString("") // body가 필요없으나, 파라미터라 선언.
	resp, err := http.Post(url, "application/json", reqBody)
	if err != nil {
		// panic(err)
		fmt.Println("Retry>>\tSession disconnected, retry setAfterCallReady() to connect")
		setAfterCallReady(tenant, agentID)
		return
	}

	defer resp.Body.Close()

	// 결과 데이터를 res에 저장
	resJson := RegisterMsg{}
	data, err := ioutil.ReadAll(resp.Body) // data는 bytep[]
	if err != nil {
		// panic(err)
		fmt.Println("Retry>>\tSession disconnected, retry setAfterCallReady() to connect")
		setAfterCallReady(tenant, agentID)
		return
	}

	err = json.Unmarshal(data, &resJson)
	if err != nil {
		// panic(err)
		fmt.Println("Retry>>\tSession disconnected, retry setAfterCallReady() to connect")
		setAfterCallReady(tenant, agentID)
		return
	}

	fmt.Println("setAfterCallReady>>\t", string(data))

	// fmt.Println(url)
}

func heartbeat() {
	option := HeartbeatQueryOption{
		session,
	}

	v, _ := query.Values(option)

	url := baseURL + "/heartbeat?" + v.Encode()

	// heartbeat 호출
	resp, err := http.Get(url)
	if err != nil {
		// panic(err)
		fmt.Println("Retry>>\tSession disconnected, retry heartbeat() to connect")
		heartbeat()
		return
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body) // data는 byte[]
	if err != nil {
		// panic(err)
		fmt.Println("Retry>>\tSession disconnected, retry heartbeat() to connect")
		heartbeat()
		return
	}

	var objmap map[string]interface{}
	if err = json.Unmarshal(data, &objmap); err != nil {
		// panic(err)
		fmt.Println("Retry>>\tSession disconnected, retry heartbeat() to connect")
		heartbeat()
		return
	}

	fmt.Println("heartbeat>>\t", string(data))

	// setReady 비활성화, setAfterCallReady로 대체됨.
	// // heartbeat에서 agentState != 40이 감지될 경우
	// // fmt.Println("agentState:", objmap["agentstate"])
	// // fmt.Println("agentState:", reflect.TypeOf(objmap["agentstate"]))
	// if objmap["agentstate"] != nil {
	// 	if int(objmap["agentstate"].(float64)) != 40 {
	// 		setReady("SSG_DEV", "test02")
	// 	}
	// }
}

func heartbeatMaker(period int) {

	c := gocron.NewScheduler(time.UTC)

	// 4
	c.Every(period).Seconds().Do(func() {
		heartbeat()
		heartbeat()
	})

	// 5
	c.StartBlocking()

}
