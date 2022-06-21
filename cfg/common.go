package cfg

//从节点心跳消息
type Heartbeat struct {
	NodeName  string `json:"nodeName"`
	NodeIp    string `json:"nodeIp"`
	DnsName   string `json:"dnsName"`
	AWSRegion string `json:"awsRegion"`
}

type HeartbeatResp struct {
	Account string `json:"account"`
	Task    string `json:"task"`
	TrackID uint   `json:"trackID"`
}

type TaskStatus struct {
	TrackID uint   `json:"trackID"`
	Result  Result `json:"result"`
	Reason  string `json:"reason"`
}

type Result int

const (
	ResultSuccess Result = iota
	ResultFail
)
