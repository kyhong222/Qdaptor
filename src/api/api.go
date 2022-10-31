package api

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

type ApiVariables struct {
	BaseURL string
	Session string
	Handle  int

	HBPeriod int
	HBErrCnt int
}

var APIVars ApiVariables
var ErrorCount int = 18 // 처음 실행하기 위해 에러카운트를 18로 시작

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

func OpenServer(AppName string) {
	option := OpenServerQueryOption{
		AppName,
	}

	v, _ := query.Values(option)

	url := APIVars.BaseURL + "/openserver?" + v.Encode()

	// openServer 호출
	resp, err := http.Get(url)
	if err != nil {
		// panic(err)
		fmt.Println("Retry>>\tSession disconnected, retry openServer() to connect")
		OpenServer(AppName)
		return
	}

	defer resp.Body.Close()

	// 결과 데이터를 res에 저장
	resJson := OpenServerMsg{}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// panic(err)
		fmt.Println("Retry>>\tSession disconnected, retry openServer() to connect")
		OpenServer(AppName)
		return
	}

	err = json.Unmarshal(data, &resJson)
	if err != nil {
		// panic(err)
		fmt.Println("Retry>>\tSession disconnected, retry openServer() to connect")
		OpenServer(AppName)
		return
	}
	fmt.Println("openServer\t", string(data))

	// session 및 handle 값 저장
	APIVars.Session = resJson.Key
	APIVars.Handle = resJson.Handle

}

func Register(DN string) {
	tenantName := "SSG_DEV"
	option := RegisterQueryOption{
		APIVars.Session, APIVars.Handle, DN, tenantName,
	}

	v, _ := query.Values(option)

	url := APIVars.BaseURL + "/register?" + v.Encode()

	// register 호출
	reqBody := bytes.NewBufferString("") // body가 필요없으나, 파라미터라 선언.
	resp, err := http.Post(url, "application/json", reqBody)
	if err != nil {
		// panic(err)
		fmt.Println("Retry>>\tSession disconnected, retry register() to connect")
		Register(DN)
		return
	}

	defer resp.Body.Close()

	// 결과 데이터를 res에 저장
	resJson := RegisterMsg{}
	data, err := ioutil.ReadAll(resp.Body) // data는 byte[]
	if err != nil {
		// panic(err)
		fmt.Println("Retry>>\tSession disconnected, retry register() to connect")
		Register(DN)
		return
	}

	err = json.Unmarshal(data, &resJson)
	if err != nil {
		// panic(err)
		fmt.Println("Retry>>\tSession disconnected, retry register() to connect")
		Register(DN)
		return
	}
	fmt.Println("register>>\t", string(data))
}

func Login(agnetID string, DN string, tenant string) {

	option := LoginQueryOption{
		APIVars.Session,
		APIVars.Handle,
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

	url := APIVars.BaseURL + "/agentlogin?" + v.Encode()

	// login 호출
	reqBody := bytes.NewBufferString("") // body가 필요없으나, 파라미터라 선언.
	resp, err := http.Post(url, "application/json", reqBody)
	if err != nil {
		// panic(err)
		fmt.Println("Retry>>\tSession disconnected, retry login() to connect")
		Login(agnetID, DN, tenant)
		return
	}

	defer resp.Body.Close()

	// 결과 데이터를 res에 저장
	resJson := RegisterMsg{}
	data, err := ioutil.ReadAll(resp.Body) // data는 bytep[]]
	if err != nil {
		// panic(err)
		fmt.Println("Retry>>\tSession disconnected, retry login() to connect")
		Login(agnetID, DN, tenant)
		return
	}

	err = json.Unmarshal(data, &resJson)
	if err != nil {
		// panic(err)
		fmt.Println("Retry>>\tSession disconnected, retry login() to connect")
		Login(agnetID, DN, tenant)
		return
	}

	fmt.Println("login>>\t", string(data))
}

func SetReady(tenant string, agentID string) {
	option := SetAgentStateQueryOption{
		APIVars.Session,
		APIVars.Handle,
		tenant,
		agentID,
		"40", // 40 is ready
		"0",  // state sub is 0
		"",   // mediaset is blank
	}

	v, _ := query.Values(option)

	url := APIVars.BaseURL + "/setagentstate?" + v.Encode()

	// login 호출
	reqBody := bytes.NewBufferString("") // body가 필요없으나, 파라미터라 선언.
	resp, err := http.Post(url, "application/json", reqBody)
	if err != nil {
		// panic(err)
		fmt.Println("Retry>>\tSession disconnected, retry setReady() to connect")
		SetReady(tenant, agentID)
		return
	}

	defer resp.Body.Close()

	// 결과 데이터를 res에 저장
	resJson := RegisterMsg{}
	data, err := ioutil.ReadAll(resp.Body) // data는 bytep[]
	if err != nil {
		// panic(err)
		fmt.Println("Retry>>\tSession disconnected, retry setReady() to connect")
		SetReady(tenant, agentID)
		return
	}

	err = json.Unmarshal(data, &resJson)
	if err != nil {
		// panic(err)
		fmt.Println("Retry>>\tSession disconnected, retry setReady() to connect")
		SetReady(tenant, agentID)
		return
	}

	fmt.Println("setReady>>\t", string(data))

	// fmt.Println(url)
}

func SetAfterCallReady(tenant string, agentID string) {
	option := SetAfterCallReadyQueryOption{
		APIVars.Session,
		APIVars.Handle,
		tenant,
		agentID,
		"40", // 40 is ready
		"0",  // state sub is 0
		"",   // mediaset is blank
	}

	v, _ := query.Values(option)

	url := APIVars.BaseURL + "/setaftcallstate?" + v.Encode()

	// login 호출
	reqBody := bytes.NewBufferString("") // body가 필요없으나, 파라미터라 선언.
	resp, err := http.Post(url, "application/json", reqBody)
	if err != nil {
		// panic(err)
		fmt.Println("Retry>>\tSession disconnected, retry setAfterCallReady() to connect")
		SetAfterCallReady(tenant, agentID)
		return
	}

	defer resp.Body.Close()

	// 결과 데이터를 res에 저장
	resJson := RegisterMsg{}
	data, err := ioutil.ReadAll(resp.Body) // data는 bytep[]
	if err != nil {
		// panic(err)
		fmt.Println("Retry>>\tSession disconnected, retry setAfterCallReady() to connect")
		SetAfterCallReady(tenant, agentID)
		return
	}

	err = json.Unmarshal(data, &resJson)
	if err != nil {
		// panic(err)
		fmt.Println("Retry>>\tSession disconnected, retry setAfterCallReady() to connect")
		SetAfterCallReady(tenant, agentID)
		return
	}

	fmt.Println("setAfterCallReady>>\t", string(data))

	// fmt.Println(url)
}

func Heartbeat() {
	option := HeartbeatQueryOption{
		APIVars.Session,
	}

	v, _ := query.Values(option)

	url := APIVars.BaseURL + "/heartbeat?" + v.Encode()

	// heartbeat 호출
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Retry>>\tSession disconnected, retry heartbeat() to connect")
		fmt.Println(err)
		ErrorCount++
		return
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body) // data는 byte[]
	if err != nil {
		fmt.Println("Retry>>\tSession disconnected, retry heartbeat() to connect")
		fmt.Println(err)
		ErrorCount++
		return
	}

	var objmap map[string]interface{}
	if err = json.Unmarshal(data, &objmap); err != nil {
		fmt.Println("Retry>>\tSession disconnected, retry heartbeat() to connect")
		fmt.Println(err)
		ErrorCount++
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

func HeartbeatMaker(period int) {

	c := gocron.NewScheduler(time.UTC)

	// 4
	c.Every(period).Seconds().Do(func() {
		Heartbeat()
		// Heartbeat()
	})

	// 5
	c.StartBlocking()

}

func Init(url string, HBP int, HBC int, appName string, DN string, tenant string, agentID string) {
	APIVars.BaseURL = url
	APIVars.HBPeriod = HBP
	APIVars.HBErrCnt = HBC

	OpenServer(appName)
	time.Sleep(1 * time.Second)
	Heartbeat()
	time.Sleep(1 * time.Second)
	Register(DN)
	time.Sleep(1 * time.Second)
	Heartbeat()
	time.Sleep(1 * time.Second)
	Heartbeat()
	time.Sleep(1 * time.Second)
	Login(agentID, DN, tenant)
	time.Sleep(1 * time.Second)
	Heartbeat()
	time.Sleep(1 * time.Second)
	Heartbeat()
	time.Sleep(1 * time.Second)
	SetReady(tenant, agentID)
	time.Sleep(1 * time.Second)
	SetAfterCallReady(tenant, agentID)
	time.Sleep(1 * time.Second)
	Heartbeat()
	time.Sleep(1 * time.Second)
	HeartbeatMaker(APIVars.HBPeriod)

	for ErrorCount >= APIVars.HBErrCnt {
		// 에러카운트가 HBErrCnt을 넘으면, 종료
		return
	}
}
