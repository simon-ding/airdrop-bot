package db

import (
	"airdrop-bot/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"path"
)

const (
	ArbitrumService = "arbitrum"
)

// Account 代表metamask里的一个账户，MetamaskIndex是必须的
type Account struct {
	gorm.Model
	Mnemonic   string `json:"mnemonic"`
	Address    string `json:"address"`
	PrivateKey string `json:"privateKey"`
	Services   []StaticIpAccountRelation
	Steps      []StepRun
}

// StaticIpAccountRelation 账户和ip的对应关系，Service为服务名称，如arbitrum等
type StaticIpAccountRelation struct {
	gorm.Model
	Service    string `json:"service,omitempty"`
	AccountID  uint   `json:"accountID,omitempty"`
	StaticIpID uint   `json:"staticIpID,omitempty"`
}

type StepRun struct {
	gorm.Model
	Service   string
	Step      string
	AccountID uint
}

// StaticIp 代表aws分配的一个ip，它可以和metamask里的一个或多个账号绑定
type StaticIp struct {
	gorm.Model
	Name     string                    `json:"name,omitempty"`
	Ip       string                    `json:"ip,omitempty"`
	Services []StaticIpAccountRelation `json:"services,omitempty"`
}

var DB *gorm.DB

func Open(dir string) {

	db1, err := gorm.Open(sqlite.Open(path.Join(dir, "test.db")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db1

	DB.AutoMigrate(&Account{})
	DB.AutoMigrate(&StaticIpAccountRelation{})
	DB.AutoMigrate(&StaticIp{})
	DB.AutoMigrate(&StepRun{})
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

func HasArbitrumStepRun(accountID uint, step string) bool {
	var c int64
	DB.Model(&StepRun{}).Where("account_id = ? AND step = ? AND service = ?", accountID, step, ArbitrumService).
		Count(&c)
	return c > 0
}

func HasArbitrumStepAllRun(accountID uint) bool {
	return HasArbitrumStepRun(accountID, StepArbitrumDeposit) && HasArbitrumStepRun(accountID, StepAaveSupplyAndBorrow) &&
		HasArbitrumStepRun(accountID, StepGmxSwap)
}

func UpdateAccountsByMnemonic(mnemonic, address string) {
	var a Account
	DB.Model(&Account{}).Where("mnemonic = ?", mnemonic).First(&a)
	a.Mnemonic = mnemonic
	a.Address = address
	log.Debugf("update account to %+v", a)
	DB.Save(&a)
}

func SaveAccount(a *Account) {
	DB.Save(a)
}

func GetOrAddIp(name string) StaticIp {
	var ip StaticIp
	DB.First(&ip, "name = ?", name)
	if ip.ID == 0 {
		log.Infof("ip of name %s not found, create one", name)
		ip.Name = name
		DB.Save(&ip)
	}
	return ip
}

func SaveAccountIpRelation(rel *StaticIpAccountRelation) {
	DB.Save(rel)
}

func SaveNewIP(name string, ip string) {
	s := StaticIp{
		Name: name,
		Ip:   ip,
	}
	DB.Save(&s)
}

func IpRelatedAccountNum(service string, ipID uint) int {

	var n int64
	DB.Model(&StaticIpAccountRelation{}).Where("static_ip_id = ? AND service= ?", ipID, service).Count(&n)
	return int(n)
}

func AccountHasIp(service string, accountID uint) (*StaticIp, bool) {
	var rel StaticIpAccountRelation
	DB.Model(&StaticIpAccountRelation{}).Where("account_id = ? AND service= ?", accountID, service).First(&rel)
	if rel.ID == 0 {
		return nil, false
	}
	var ip StaticIp
	DB.First(&ip, rel.StaticIpID)

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
