package huobi

import (
	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
)

func main() {

	// Get the list of accounts owned by this API user and print the detail on console
	client2 := new(client.AccountClient).Init("config.AccessKey", "config.SecretKey", "config.Host")
	resp, err := client2.GetAccountInfo()
	if err != nil {
		applogger.Error("Get account error: %s", err)
	} else {
		applogger.Info("Get account, count=%d", len(resp))
		for _, result := range resp {
			applogger.Info("account: %+v", result)
		}
	}

}
