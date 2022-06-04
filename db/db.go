package db

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	ArbitrumService = "arbitrum"
)

// Account 代表metamask里的一个账户，MetamaskIndex是必须的
type Account struct {
	gorm.Model
	MetamaskIndex int    `json:"metamaskIndex"`
	Address       string `json:"address"`
	PrivateKey    string `json:"privateKey"`
	Services      []StaticIpAccountRelation
}

// StaticIpAccountRelation 账户和ip的对应关系，Service为服务名称，如arbitrum等
type StaticIpAccountRelation struct {
	gorm.Model
	Service    string `json:"service,omitempty"`
	AccountID  uint   `json:"accountID,omitempty"`
	StaticIpID uint   `json:"staticIpID,omitempty"`
}

// StaticIp 代表aws分配的一个ip，它可以和metamask里的一个或多个账号绑定
type StaticIp struct {
	gorm.Model
	Name     string                    `json:"name,omitempty"`
	Ip       string                    `json:"ip,omitempty"`
	Services []StaticIpAccountRelation `json:"services,omitempty"`
}

var db *gorm.DB

func Open() {

	db1, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db = db1

	db.AutoMigrate(&Account{})
	db.AutoMigrate(&StaticIpAccountRelation{})
	db.AutoMigrate(&StaticIp{})
}

func SaveNewAccount(metamaskIndex int, address, privateKey string) {
	a := Account{
		MetamaskIndex: metamaskIndex,
		Address:       address,
		PrivateKey:    privateKey,
	}
	db.Save(&a)
}

func SaveNewIP(name string, ip string) {
	s := StaticIp{
		Name: name,
		Ip:   ip,
	}
	db.Save(&s)
}

func SaveIpAccountRelation(service string, metamaskID int, ipName string) error {
	var a Account
	db.First(&a, "metamaskIndex = ?", metamaskID)
	if a.MetamaskIndex == 0 {
		return fmt.Errorf("metamask account of id %d not found", metamaskID)
	}
	var ip StaticIp
	db.First(&ip, "name = ?", ipName)
	if ip.Name == "" {
		return fmt.Errorf("static ip of name %s not found", ipName)
	}

	rel := StaticIpAccountRelation{
		Service:    service,
		AccountID:  a.ID,
		StaticIpID: ip.ID,
	}
	db.Save(rel)
	return nil
}

func IpRelatedAccountNum(service string, ipName string) int {
	var ip StaticIp
	db.First(&ip, "name = ?", ipName)
	if ip.Name == "" {
		return 0
	}

	var n int64
	db.Model(&StaticIpAccountRelation{}).Where("staticIpID = ?", ip.ID).Count(&n)
	return int(n)
}
