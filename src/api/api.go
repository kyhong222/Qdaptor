package api

import (
	"Qdaptor/logger"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/go-co-op/gocron"
	"github.com/google/go-querystring/query"
)

type Agent struct {
	APIVars           ApiVariables
	DN                string
	Tenant            string
	AgentID           string
	IVRResultResponse map[string]interface{}

	ErrorCount  int
	IsApiCalled bool
}

type ApiVariables struct {
	Session string
	Handle  int

	HBPeriod int
	HBErrCnt int

	ConnectionID string
	UCID         string
}

// var APIVars ApiVariables
// var ErrorCount int = 18 // 처음 실행하기 위해 에러카운트를 18로 시작

// var APIWaitGroup = sync.WaitGroup{}
var BaseURL string = "https://dev-icweb.ssgadm.com:9203/ic"

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

type RefCallQueryOption struct {
	Key           string `url:"key"`
	Handle        int    `url:"handle"`
	ThisDN        string `url:"thisdn"`
	DestDN        string `url:"destdn"`
	ObCallingDN   string `url:"obcallingdn"`
	ConnectionID  string `url:"connectionid"`
	PartyType     int    `url:"partytype"`
	MediaType     int    `url:"mediatype"`
	ExtensionData string `url:"extensiondata"`
}

type CallClearQueryOption struct {
	Key           string `url:"key"`
	Handle        int    `url:"handle"`
	ThisDN        string `url:"thisdn"`
	ConnectionID  string `url:"connectionid"`
	MediaType     int    `url:"mediatype"`
	ExtensionData string `url:"extensiondata"`
}

type GetQueueTrafficQueryOption struct {
	Key         string `url:"key"`
	Handle      int    `url:"handle"`
	Tenant      string `url:"tenantname"`
	QueueDN     string `url:"queuedn"`
	SkillID     int    `url:"skillid"`
	PrivateData string `url:"privatedata"`
	Mediaset    string `url:"mediaset"`
}

func (a *Agent) OpenServer(AppName string) {
	a.IsApiCalled = true
	option := OpenServerQueryOption{
		AppName,
	}

	v, _ := query.Values(option)

	url := BaseURL + "/openserver?" + v.Encode()

	// openServer 호출
	resp, err := http.Get(url)
	if err != nil {
		// panic(err)
		// fmt.Println("Retry>>\tSession disconnected, retry openServer() to connect")
		logger.Error("Session disconnected, retry openServer() to connect")
		a.OpenServer(AppName)
		return
	}

	defer resp.Body.Close()

	// 결과 데이터를 res에 저장
	resJson := OpenServerMsg{}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// panic(err)
		// fmt.Println("Retry>>\tSession disconnected, retry openServer() to connect")
		logger.Error("Session disconnected, retry openServer() to connect")
		a.OpenServer(AppName)
		return
	}

	err = json.Unmarshal(data, &resJson)
	if err != nil {
		// panic(err)
		// fmt.Println("Retry>>\tSession disconnected, retry openServer() to connect")
		logger.Error("Session disconnected, retry openServer() to connect")
		a.OpenServer(AppName)
		return
	}
	// fmt.Println("openServer\t", string(data))
	logger.Info("OpenServer",
		zap.Reflect("response", resJson),
	)

	// session 및 handle 값 저장
	a.APIVars.Session = resJson.Key
	a.APIVars.Handle = resJson.Handle

}

func (a *Agent) Register() {
	a.IsApiCalled = true
	tenantName := a.Tenant

	option := RegisterQueryOption{
		a.APIVars.Session, a.APIVars.Handle, a.DN, tenantName,
	}

	v, _ := query.Values(option)

	url := BaseURL + "/register?" + v.Encode()

	// register 호출
	reqBody := bytes.NewBufferString("") // body가 필요없으나, 파라미터라 선언.
	resp, err := http.Post(url, "application/json", reqBody)
	if err != nil {
		// panic(err)
		// fmt.Println("Retry>>\tSession disconnected, retry register() to connect")
		logger.Error("Session disconnected, retry register() to connect")
		a.Register()
		return
	}

	defer resp.Body.Close()

	// 결과 데이터를 res에 저장
	resJson := RegisterMsg{}
	data, err := ioutil.ReadAll(resp.Body) // data는 byte[]
	if err != nil {
		// panic(err)
		// fmt.Println("Retry>>\tSession disconnected, retry register() to connect")
		logger.Error("Session disconnected, retry register() to connect")
		a.Register()
		return
	}

	err = json.Unmarshal(data, &resJson)
	if err != nil {
		// panic(err)
		// fmt.Println("Retry>>\tSession disconnected, retry register() to connect")
		logger.Error("Session disconnected, retry register() to connect")
		a.Register()
		return
	}
	// fmt.Println("register>>\t", string(data))
	logger.Info("Register",
		zap.Reflect("response", resJson),
	)

}

func (a *Agent) Login() {
	a.IsApiCalled = true

	option := LoginQueryOption{
		a.APIVars.Session,
		a.APIVars.Handle,
		a.Tenant,
		a.DN,
		a.AgentID,
		"",
		"40", // 40 is ready
		"0",  // state sub is 0
		"4",  // 4 is SHA-2(512)
		"",   // mediaset is blank
	}

	v, _ := query.Values(option)

	url := BaseURL + "/agentlogin?" + v.Encode()

	// login 호출
	reqBody := bytes.NewBufferString("") // body가 필요없으나, 파라미터라 선언.
	resp, err := http.Post(url, "application/json", reqBody)
	if err != nil {
		// panic(err)
		// fmt.Println("Retry>>\tSession disconnected, retry login() to connect")
		logger.Error("Session disconnected, retry login() to connect")
		a.Login()
		return
	}

	defer resp.Body.Close()

	// 결과 데이터를 res에 저장
	resJson := RegisterMsg{}
	data, err := ioutil.ReadAll(resp.Body) // data는 bytep[]]
	if err != nil {
		// panic(err)
		// fmt.Println("Retry>>\tSession disconnected, retry login() to connect")
		logger.Error("Session disconnected, retry login() to connect")
		a.Login()
		return
	}

	err = json.Unmarshal(data, &resJson)
	if err != nil {
		// panic(err)
		// fmt.Println("Retry>>\tSession disconnected, retry login() to connect")
		logger.Error("Session disconnected, retry login() to connect")
		a.Login()
		return
	}

	// fmt.Println("login>>\t", string(data))
	logger.Info("Login",
		zap.Reflect("response", resJson),
	)
}

func (a *Agent) SetReady() {
	a.IsApiCalled = true

	option := SetAgentStateQueryOption{
		a.APIVars.Session,
		a.APIVars.Handle,
		a.Tenant,
		a.AgentID,
		"40", // 40 is ready
		"0",  // state sub is 0
		"",   // mediaset is blank
	}

	v, _ := query.Values(option)

	url := BaseURL + "/setagentstate?" + v.Encode()

	// login 호출
	reqBody := bytes.NewBufferString("") // body가 필요없으나, 파라미터라 선언.
	resp, err := http.Post(url, "application/json", reqBody)
	if err != nil {
		// panic(err)
		// fmt.Println("Retry>>\tSession disconnected, retry setReady() to connect")
		logger.Error("Session disconnected, retry setReady() to connect")
		a.SetReady()
		return
	}

	defer resp.Body.Close()

	// 결과 데이터를 res에 저장
	resJson := RegisterMsg{}
	data, err := ioutil.ReadAll(resp.Body) // data는 bytep[]
	if err != nil {
		// panic(err)
		// fmt.Println("Retry>>\tSession disconnected, retry setReady() to connect")
		logger.Error("Session disconnected, retry setReady() to connect")
		a.SetReady()
		return
	}

	err = json.Unmarshal(data, &resJson)
	if err != nil {
		// panic(err)
		// fmt.Println("Retry>>\tSession disconnected, retry setReady() to connect")
		logger.Error("Session disconnected, retry setReady() to connect")
		a.SetReady()
		return
	}

	// fmt.Println("setReady>>\t", string(data))
	logger.Info("SetReady",
		zap.Reflect("response", resJson),
	)

	// fmt.Println(url)
}

func (a *Agent) SetAfterCallReady() {
	a.IsApiCalled = true

	option := SetAfterCallReadyQueryOption{
		a.APIVars.Session,
		a.APIVars.Handle,
		a.Tenant,
		a.AgentID,
		"40", // 40 is ready
		"0",  // state sub is 0
		"",   // mediaset is blank
	}

	v, _ := query.Values(option)

	url := BaseURL + "/setaftcallstate?" + v.Encode()

	// login 호출
	reqBody := bytes.NewBufferString("") // body가 필요없으나, 파라미터라 선언.
	resp, err := http.Post(url, "application/json", reqBody)
	if err != nil {
		// panic(err)
		// fmt.Println("Retry>>\tSession disconnected, retry setAfterCallReady() to connect")
		logger.Error("Session disconnected, retry setAfterCallReady() to connect")
		a.SetAfterCallReady()
		return
	}

	defer resp.Body.Close()

	// 결과 데이터를 res에 저장
	resJson := RegisterMsg{}
	data, err := ioutil.ReadAll(resp.Body) // data는 bytep[]
	if err != nil {
		// panic(err)
		// fmt.Println("Retry>>\tSession disconnected, retry setAfterCallReady() to connect")
		logger.Error("Session disconnected, retry setAfterCallReady() to connect")
		a.SetAfterCallReady()
		return
	}

	err = json.Unmarshal(data, &resJson)
	if err != nil {
		// panic(err)
		// fmt.Println("Retry>>\tSession disconnected, retry setAfterCallReady() to connect")
		logger.Error("Session disconnected, retry setAfterCallReady() to connect")
		a.SetAfterCallReady()
		return
	}

	// fmt.Println("setAfterCallReady>>\t", string(data))
	logger.Info("SetAfterCallReady",
		zap.Reflect("response", resJson),
	)
	// fmt.Println(url)
}

func (a *Agent) Heartbeat() {
	option := HeartbeatQueryOption{
		a.APIVars.Session,
	}

	v, _ := query.Values(option)

	url := BaseURL + "/heartbeat?" + v.Encode()
	logger.Info("Heartbeat called",
		zap.String("url", url),
	)

	// heartbeat 호출
	resp, err := http.Get(url)
	if err != nil {
		// fmt.Println("Retry>>\tSession disconnected, retry heartbeat() to connect")
		// fmt.Println(err)
		a.ErrorCount++
		logger.Error("Session disconnected, retry heartbeat() to connect",
			zap.Error(err),
			zap.Int("ErrorCount", a.ErrorCount),
		)
		return
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body) // data는 byte[]
	if err != nil {
		// fmt.Println("Retry>>\tSession disconnected, retry heartbeat() to connect")
		// fmt.Println(err)
		a.ErrorCount++
		logger.Error("Session disconnected, retry heartbeat() to connect",
			zap.Error(err),
			zap.Int("ErrorCount", a.ErrorCount),
		)
		return
	}

	var objmap map[string]interface{}
	if err = json.Unmarshal(data, &objmap); err != nil {
		// fmt.Println("Retry>>\tSession disconnected, retry heartbeat() to connect")
		// fmt.Println(err)
		a.ErrorCount++
		logger.Error("Session disconnected, retry heartbeat() to connect",
			zap.Error(err),
			zap.Int("ErrorCount", a.ErrorCount),
		)
		return
	}

	// fmt.Println("heartbeat>>\t", string(data))
	logger.Info("HeartBeat",
		zap.Reflect("response", objmap),
	)

	// heartbeat message processing
	switch objmap["messagetype"].(float64) {
	case 1:
		go a.Heartbeat()
	case 2: // messagetype: 2 is IVR response
		switch objmap["method"].(float64) {
		case 1051: //  method: 1051 is getQueueTraffic
			// get full IVR response
			logger.Info("IVRResultResponse is arrived(1051)",
				zap.Reflect("response", objmap))
			a.IVRResultResponse = objmap
		}
		go a.Heartbeat()
	case 3: // messagetype: 3 is IVR event
		switch objmap["method"].(float64) {
		case 2000: //  method: 2000 is ringing
			// get connectionID & UCID
			logger.Info("IVRResultResponse is arrived(2000)",
				zap.Reflect("response", objmap))
			a.APIVars.ConnectionID = objmap["connectionid"].(string)
			a.APIVars.UCID = objmap["ucid"].(string)
			a.IVRResultResponse = objmap
		// case 2001: //  method: 2001 is established
		// 	// get connectionID & UCID
		// 	logger.Info("IVRResultResponse is arrived(2001)",
		// 		zap.Reflect("response", objmap))
		// 	APIVars.ConnectionID = objmap["connectionid"].(string)
		// 	APIVars.UCID = objmap["ucid"].(string)
		// 	IVRResultResponse = objmap
		// case 2002: //  method: 2002 is release
		// 	// get connectionID & UCID
		// 	APIVars.ConnectionID = objmap["connectionid"].(string)
		// 	APIVars.UCID = objmap["ucid"].(string)
		// 	IVRResultResponse = objmap
		case 2010: //  method: 2010 is party delete(means IVR is ended)
			// get full IVR Response
			logger.Info("IVRResultResponse is arrived(2010)",
				zap.Reflect("response", objmap))
			a.IVRResultResponse = objmap
		}
		go a.Heartbeat()
	}

	// // messagetype: 2 is IVR response
	// if objmap["messagetype"].(float64) == 2 {
	// 	switch objmap["method"].(float64) {
	// 	case 1051: //  method: 1051 is getQueueTraffic
	// 		// get full IVR response
	// 		IVRResultResponse = objmap
	// 	}
	// 	// APIWaitGroup.Done()
	// }

	// // messagetype: 3 is IVR event
	// if objmap["messagetype"].(float64) == 3 {
	// 	switch objmap["method"].(float64) {
	// 	case 2000: //  method: 2000 is ringing
	// 		// get connectionID & UCID
	// 		APIVars.ConnectionID = objmap["connectionid"].(string)
	// 		APIVars.UCID = objmap["ucid"].(string)
	// 		IVRResultResponse = objmap
	// 	case 2001: //  method: 2001 is established
	// 		// get connectionID & UCID
	// 		APIVars.ConnectionID = objmap["connectionid"].(string)
	// 		APIVars.UCID = objmap["ucid"].(string)
	// 		IVRResultResponse = objmap
	// 	case 2002: //  method: 2002 is ringing
	// 		// get connectionID & UCID
	// 		APIVars.ConnectionID = objmap["connectionid"].(string)
	// 		APIVars.UCID = objmap["ucid"].(string)
	// 		IVRResultResponse = objmap
	// 	case 2010: //  method: 2010 is party delete(means IVR is ended)
	// 		// get full IVR Response
	// 		IVRResultResponse = objmap
	// 	}
	// 	Heartbeat()
	// 	// APIWaitGroup.Done()
	// }

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

func (a *Agent) HeartbeatMaker() {

	c := gocron.NewScheduler(time.UTC)

	// 4
	c.Every(a.APIVars.HBPeriod).Seconds().Do(func() {
		if a.IsApiCalled {
			a.IsApiCalled = false
		} else {
			go a.Heartbeat()
		}
	})

	// 5
	c.StartBlocking()

}

func (a *Agent) RefCall(destDN string) {
	// key: EC8ACECF-5F92-4BEB-B93F-32896C2F0450
	// handle: 974
	// thisdn: 5205
	// destdn: 8993
	// obcallingdn:
	// connectionid: 6bab040d03a410
	// partytype: 4
	// mediatype: 0
	// extensiondata: {"NANSAGTID":["2022002106"]}

	option := RefCallQueryOption{
		a.APIVars.Session,
		a.APIVars.Handle,
		a.DN,
		destDN,
		"",
		a.APIVars.ConnectionID,
		4, // 4 is Mute
		0,
		"", // 브리지텍 측에서 NANSAGTID 필요없다 하심.
	}

	v, _ := query.Values(option)

	url := BaseURL + "/singlesteptransfer?" + v.Encode()

	// singlesteptransfer 호출
	resp, err := http.Get(url)
	if err != nil {
		logger.Error("singleStepTransfer failed",
			zap.Error(err),
		)
		return
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body) // data는 byte[]
	if err != nil {
		logger.Error("singleStepTranfer response read failed",
			zap.Error(err),
		)
		return
	}

	var objmap map[string]interface{}
	if err = json.Unmarshal(data, &objmap); err != nil {
		logger.Error("singleStepTranfer response Unmarshal failed ",
			zap.Error(err),
		)
		return
	}

	// fmt.Println("heartbeat>>\t", string(data))
	logger.Info("singleStepTranfer",
		zap.Reflect("response", objmap),
	)
}

func (a *Agent) CallClear(extensionData string) {
	// key: 78D6FD84-F167-42CE-A717-67268F5F6530
	// handle: 44
	// thisdn: 5205
	// connectionid:
	// mediatype: 0
	// extensiondata

	// TODO: make extensionData

	option := CallClearQueryOption{
		a.APIVars.Session,
		a.APIVars.Handle,
		a.DN,
		a.APIVars.ConnectionID,
		0,
		extensionData,
	}

	v, _ := query.Values(option)

	url := BaseURL + "/clearcall?" + v.Encode()

	// callClear 호출
	resp, err := http.Get(url)
	if err != nil {
		logger.Error("callClear failed",
			zap.Error(err),
		)
		return
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body) // data는 byte[]
	if err != nil {
		logger.Error("callClear response read failed",
			zap.Error(err),
		)
		return
	}

	var objmap map[string]interface{}
	if err = json.Unmarshal(data, &objmap); err != nil {
		logger.Error("callClear response Unmarshal failed ",
			zap.Error(err),
		)
		return
	}

	// fmt.Println("heartbeat>>\t", string(data))
	logger.Info("callClear",
		zap.Reflect("response", objmap),
	)
}

// deprecated on 221116
func (a *Agent) GetQueueTraffic(QueueDN string) {
	// key: BB1D0BF5-D949-4BAD-B972-13A2FB85C7B0
	// handle: 14
	// tenantname:
	// queuedn: 8821
	// skillid: 0
	// privatedata:
	// mediaset:

	option := GetQueueTrafficQueryOption{
		a.APIVars.Session,
		a.APIVars.Handle,
		"",
		QueueDN,
		0,
		"",
		"",
	}

	v, _ := query.Values(option)

	url := BaseURL + "/getqueuetraffic?" + v.Encode()

	// GetQueueTraffic 호출
	resp, err := http.Get(url)
	if err != nil {
		logger.Error("getQueueTraffic failed",
			zap.Error(err),
		)
		return
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body) // data는 byte[]
	if err != nil {
		logger.Error("getQueueTraffic response read failed",
			zap.Error(err),
		)
		return
	}

	var objmap map[string]interface{}
	if err = json.Unmarshal(data, &objmap); err != nil {
		logger.Error("getQueueTraffic response Unmarshal failed ",
			zap.Error(err),
		)
		return
	}

	logger.Info("getQueueTraffic",
		zap.Reflect("response", objmap),
	)

}

func (a *Agent) Start() {
	a.ErrorCount = a.APIVars.HBErrCnt
	for {
		if a.ErrorCount >= a.APIVars.HBErrCnt {
			a.ErrorCount = 0
			a.Run()
		}
	}
}

func (a *Agent) Run() {
	appName := "SSG_VoiceBot_" + a.DN + "_" + a.AgentID

	a.OpenServer(appName)
	time.Sleep(1 * time.Second)
	a.Register()
	time.Sleep(1 * time.Second)
	a.Login()
	time.Sleep(1 * time.Second)
	a.SetReady()
	time.Sleep(1 * time.Second)
	a.SetAfterCallReady()
	time.Sleep(1 * time.Second)
	a.HeartbeatMaker()

	for a.ErrorCount >= a.APIVars.HBErrCnt {
		// 에러카운트가 HBErrCnt을 넘으면, 종료
		time.Sleep(10 * time.Second)
		return
	}
}
