package cfg

import (
	"airdrop-bot/log"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

const (
	ChromeModXvfb = "xvfb"
)

type Config struct {
	Server ServerConfig `mapstructure:"server"`
	Node   NodeConfig   `mapstructure:"node"`
}

type ServerConfig struct {
	AccountsPerIp    int        `mapstructure:"accountsPerIp"`
	AccountsToGen    int        `mapstructure:"accountsToGen"`
	Owlracle         Owlracle   `mapstructure:"owlracle"`
	GasFeeAcceptable int        `mapstructure:"gasFeeAcceptable"`
	Cloudflare       Cloudflare `mapstructure:"cloudflare"`
	AWS              AWS        `mapstructure:"aws"`
	Binance          Binance    `mapstructure:"binance"`
	Token            string     `mapstructure:"token"`
	Dir              string     `mapstructure:"dir"`
	DbFile           string     `mapstructure:"dbFile"`
}

type NodeConfig struct {
	WalletPassword string `mapstructure:"walletPassword"`
	Dir            string `mapstructure:"dir"`
	ChromeMode     string `mapstructure:"chromeMode"`
	DnsName        string `mapstructure:"dnsName"`
	NodeName       string `mapstructure:"nodeName"`
	Region         string `mapstructure:"region"`
	Token          string `mapstructure:"token"`
	ServerUrl      string `mapstructure:"serverUrl"`
}

func (c *NodeConfig) XvfbMod() bool {
	return c.ChromeMode == ChromeModXvfb
}

type AWS struct {
	AccessKeyID     string `mapstructure:"accessKeyID"`
	SecretAccessKey string `mapstructure:"secretAccessKey"`
}

type Cloudflare struct {
	ApiKey    string `mapstructure:"apiKey"`
	ZoneId    string `mapstructure:"zoneId"`
	KvId      string `mapstructure:"kvId"`
	AccountId string `mapstructure:"accountId"`
}

type Owlracle struct {
	ApiKey    string `mapstructure:"apiKey"`
	ApiSecret string `mapstructure:"apiSecret"`
	Disable   bool   `mapstructure:"disable"`
}

type Binance struct {
	ApiKey    string `mapstructure:"apiKey"`
	SecretKey string `mapstructure:"secretKey"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")

	var cc Config
	// optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Info("create config file")
			viper.SafeWriteConfig()
		} else {
			// Config file was found but another error was produced
		}

		return nil, errors.Wrap(err, "load config")
	}

	if err := viper.Unmarshal(&cc); err != nil {
		return nil, errors.Wrap(err, "unmarshal file")
	}
	return &cc, err
}
