package db

import (
	"airdrop-bot/cfg"
	"airdrop-bot/ent"
	"airdrop-bot/ent/account"
	"airdrop-bot/ent/node"
	"airdrop-bot/ent/steprun"
	"airdrop-bot/log"
	"context"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const (
	ArbitrumService = "arbitrum"
)

func GetClient() *ent.Client {
	return client
}

var client *ent.Client

func Open(dir string) {

	c, err := ent.Open("sqlite3", fmt.Sprintf("file:%v?&cache=shared&_fk=1", dir))
	if err != nil {
		log.Panicf("failed opening connection to sqlite: %v", err)
	}

	// Run the auto migration tool.
	if err := c.Schema.Create(context.Background()); err != nil {
		log.Panicf("failed creating schema resources: %v", err)
	}
	client = c
}

func StepBeenDone(accountId int, step string) bool {
	c, err := client.StepRun.Query().Where(
		steprun.HasAccountWith(account.ID(int(accountId))),
		steprun.Step(step),
	).Count(context.Background())
	if err != nil {
		return false
	}
	return c > 0
}

func FindFirstPendingTask(nodeID int) *ent.StepRun {
	s, err := client.StepRun.Query().Where(
		steprun.HasNodeWith(node.ID(nodeID)),
		steprun.StatusEQ(steprun.StatusPending),
	).First(context.TODO())
	if err != nil {
		log.Errorf("query step tun error: %v", err)
		return nil
	}
	return s
}

func AddOrUpdateNode(c *cfg.Heartbeat) *ent.Node {

	n, err := client.Node.Query().Where(node.Name(c.NodeName)).First(context.Background())
	if err != nil {
		log.Infof("node %s not found, save new one", c.NodeName)
		n1, err := client.Node.Create().SetName(c.NodeName).SetDnsName(c.DnsName).SetRegion(c.AWSRegion).Save(context.Background())
		if err != nil {
			log.Errorf("create node error: %v", err)
		}
		n = n1
	}
	client.Node.Update().Where(node.ID(n.ID)).SetUpdatedAt(time.Now()).ExecX(context.Background())
	return n
}

func FindAccount(id int) *ent.Account {
	a, err := client.Account.Get(context.Background(), id)
	if err != nil {
		log.Errorf("account %d not found: %v", id, err)
		return nil
	}

	return a
}

func FetchAllAccounts() []*ent.Account {
	return client.Account.Query().AllX(context.Background())
}

func AccountNum() int {
	n, err := client.Account.Query().Count(context.Background())
	if err != nil {
		log.Errorf("count account error: %v", err)
		return 0
	}
	return n
}

func SaveAccount(mnemonic string, address string, privateKey string) {
	client.Account.Create().SetAddress(address).SetMnemonic(mnemonic).SetPrivateKey(privateKey).SaveX(context.Background())
}

func PickANodeRandom() *ent.Node {
	n, err := client.Node.Query().Order(ent.Desc(node.FieldUpdatedAt)).First(context.Background())
	if err != nil {
		log.Errorf("query node error: %v", err)
		return nil
	}
	return n
}
