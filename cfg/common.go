package cfg

import "time"

//从节点心跳消息
type Heartbeat struct {
	NodeName  string `json:"nodeName"`
	DnsName   string `json:"dnsName"`
	AWSRegion string `json:"awsRegion"`
}

type HeartbeatResp struct {
	Account string        `json:"account"`
	Task    string        `json:"task"`
	TrackID int           `json:"trackID"`
	Timeout time.Duration `json:"timeout"`
}

type TaskStatus struct {
	TrackID int    `json:"trackID"`
	Result  Result `json:"result"`
	Reason  string `json:"reason"`
}

type Result int

const (
	ResultSuccess Result = iota
	ResultFail
)
