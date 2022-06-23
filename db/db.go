package db

import (
	"airdrop-bot/cfg"
	"airdrop-bot/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	log2 "log"
	"os"
	"path"
	"time"
)

const (
	ArbitrumService = "arbitrum"
)

// Account 代表metamask里的一个账户
type Account struct {
	gorm.Model
	Mnemonic   string `json:"mnemonic"`
	Address    string `json:"address"`
	PrivateKey string `json:"privateKey"`
	StaticIpID uint
	Steps      []StepRun
}

type StepRun struct {
	gorm.Model
	Service    string
	Step       string
	Status     Status
	Reason     string
	AccountID  uint
	NodeID     uint
	StaticIpID uint
}

type Status int

const (
	StatusPending Status = iota
	StatusRunning
	StatusSuccess
	StatusFailed
)

// StaticIp 代表aws分配的一个ip，它可以和metamask里的一个或多个账号绑定
type StaticIp struct {
	gorm.Model
	Name     string    `json:"name,omitempty"`
	Ip       string    `json:"ip,omitempty"`
	Accounts []Account `json:"accounts,omitempty"`
	NodeID   uint
}

type Node struct {
	gorm.Model
	NodeName  string
	Region    string
	DnsName   string
	NodeIp    string
	StaticIps []StaticIp
}

var DB *gorm.DB

func Open(dir string) {
	newLogger := logger.New(
		log2.New(os.Stdout, "\r\n", log2.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)

	db1, err := gorm.Open(sqlite.Open(path.Join(dir, "test.db")), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db1

	DB.AutoMigrate(&Account{})
	DB.AutoMigrate(&StaticIp{})
	DB.AutoMigrate(&StepRun{})
	DB.AutoMigrate(&Node{})
}

func StepBeenDone(accountId uint, step string) bool {
	var s StepRun
	DB.First(&s, "account_id = ? AND step = ?", accountId, step)
	if s.ID == 0 {
		return false
	}
	if s.Status == StatusSuccess {
		return true
	}
	DB.Delete(&s, "id = ?", s.ID)
	log.Infof("delete step, rerun: %+v", s)
	return false
}

func FindFirstPendingTask(nodeID uint) *StepRun {
	var s StepRun
	DB.First(&s, "status = ? AND node_id = ?", StatusPending, nodeID)
	if s.ID == 0 {
		return nil
	}
	return &s
}

func AddOrUpdateNode(c *cfg.Heartbeat) Node {
	var n Node
	DB.First(&n, "node_name = ?", c.NodeName)
	if n.ID == 0 || n.DnsName != c.DnsName || n.NodeName != c.NodeName || n.Region != c.AWSRegion {
		n.NodeName = c.NodeName
		n.DnsName = c.DnsName
		n.Region = c.AWSRegion
		DB.Save(&n)
		log.Infof("update node: %+v", n)
	}
	return n
}

func FindAccount(id int) Account {
	var a Account
	DB.First(&a, id)
	if a.ID == 0 {
		log.Errorf("account with id %d not found", id)
	}
	return a
}

func AccountNum() int {
	var c int64
	DB.Model(&Account{}).Count(&c)
	return int(c)
}

func SaveArbitrumStep(accountID uint, step string) {
	s := StepRun{
		Service:   ArbitrumService,
		Step:      step,
		AccountID: accountID,
	}
	log.Infof("arbitrum service step %s has done", step)
	DB.Save(&s)
}

func SaveAccount(a *Account) {
	DB.Save(a)
}

func GetOrAddIp(name string, nodeId uint) StaticIp {
	var ip StaticIp
	DB.First(&ip, "name = ?", name)
	if ip.ID == 0 {
		log.Infof("ip of name %s not found, create one", name)
		ip.Name = name
		ip.NodeID = nodeId
		DB.Save(&ip)
		log.Infof("save new ip: %+v", ip)
	}
	return ip
}

func IpRelatedAccountNum(ipID uint) int {
	var n int64
	DB.Model(&Account{}).Where("static_ip_id = ?", ipID).Count(&n)
	return int(n)
}

func AccountHasIp(staticIpId uint) (*StaticIp, bool) {
	if staticIpId == 0 {
		return nil, false
	}
	var ip StaticIp
	DB.First(&ip, staticIpId)

	return &ip, true
}

func SaveOrUpdateIp(name, address string) {
	var ip StaticIp
	DB.First(&ip, "name = ?", name)
	if ip.ID == 0 || ip.Ip != address {
		//not found or need update
		ip.Name = name
		ip.Ip = address
		DB.Save(&ip)
	}
}

func PickANode() Node {
	var c int64
	DB.Model(&Node{}).Count(&c)
	for i := 0; i < int(c); i++ {
		if !nodeHasRunningTask(uint(i + 1)) {
			var n Node
			DB.First(&n, i+1)

			return n
		}
	}
	return Node{}
}

func nodeHasRunningTask(nodeID uint) bool {
	var c int64
	DB.Model(&StepRun{}).Where("status = ?", nodeID).Count(&c)
	return c != 0
}
